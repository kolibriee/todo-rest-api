package service

import (
	"errors"
	"time"

	"github.com/kostylevdev/todo-rest-api/internal/domain"
	"github.com/kostylevdev/todo-rest-api/internal/repository"
	"github.com/kostylevdev/todo-rest-api/pkg/auth"
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
	user.Password = auth.GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) SignIn(clientIP string, signinuser domain.SignInUserInput) (string, string, error) {
	signinuser.Password = auth.GeneratePasswordHash(signinuser.Password)
	user, err := s.repo.GetUser(signinuser)
	if err != nil {
		return "", "", err
	}
	accessToken, err := auth.GenerateAccessToken(accessTokenTTL, user.Id)
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
		return "", "", errors.New("invalid refresh token" + err.Error())
	}
	if session.ClientIP != IP {
		return "", "", errors.New("invalid ip")
	}
	if session.ExpiresAt.Before(time.Now()) {
		return "", "", errors.New("session expired")
	}
	accessToken, err := auth.GenerateAccessToken(accessTokenTTL, session.UserId)
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
