package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kostylevdev/todo-rest-api/internal/domain"
	"github.com/kostylevdev/todo-rest-api/internal/repository"
)

const (
	accessTokenTTL  = 15 * time.Minute
	refreshTokenTTL = 30 * 24 * time.Hour
)

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(user domain.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) SignIn(clientIP string, signinuser domain.SignInUserInput) (string, string, error) {
	signinuser.Password = s.generatePasswordHash(signinuser.Password)
	user, err := s.repo.GetUser(signinuser)
	if err != nil {
		return "", "", err
	}
	accessToken, err := s.GenerateAccessToken(user.Id)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := s.repo.CreateSession(domain.Session{
		UserId:    user.Id,
		ExpiresAt: time.Now().Add(refreshTokenTTL),
		ClientIP:  clientIP,
	})
	if err != nil {
		return "", "", errors.New("can't create refresh token")
	}
	return accessToken, refreshToken, nil
}

func (s *AuthService) Refresh(refreshToken string, IP string) (string, string, error) {
	session, err := s.repo.GetSession(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}
	if session.ClientIP != IP {
		return "", "", errors.New("invalid ip")
	}
	if session.ExpiresAt.Before(time.Now()) {
		return "", "", errors.New("session expired")
	}
	accessToken, err := s.GenerateAccessToken(session.UserId)
	if err != nil {
		return "", "", err
	}
	newRefreshToken, err := s.repo.CreateSession(domain.Session{
		UserId:    session.UserId,
		ExpiresAt: time.Now().Add(refreshTokenTTL),
		ClientIP:  IP,
	})
	if err != nil {
		return "", "", errors.New("can't create refresh token")
	}
	if err := s.repo.DeleteSession(refreshToken); err != nil {
		return "", "", errors.New("can't delete old refresh token")
	}
	return accessToken, newRefreshToken, nil
}
func (s *AuthService) GenerateAccessToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", userId),
		ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
	)
	accessToken, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	if err != nil {
		return "", errors.New("can't generate access token")
	}
	return accessToken, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("TOKEN_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_HASH_SALT"))))
}
