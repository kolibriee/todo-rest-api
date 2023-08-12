package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	tryrest "github.com/kolibri7557/try-rest-api"
	"github.com/kolibri7557/try-rest-api/pkg/repository"
)

const (
	tokenTTL   = 12 * time.Hour
	signingKey = "dfdsfe3f3gdgdfdh&#jkl#sq"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user tryrest.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(signinuser tryrest.SignInUser) (int, string, error) {
	signinuser.Password = s.generatePasswordHash(signinuser.Password)
	user, err := s.repo.GetUser(signinuser)
	if err != nil {
		return 0, "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	tokenGenerated, err := token.SignedString([]byte(signingKey))
	return user.Id, tokenGenerated, err
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_HASH_SALT"))))
}
