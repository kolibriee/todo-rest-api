package service

import (
	"crypto/sha1"
	"fmt"
	"os"

	tryrest "github.com/kolibri7557/try-rest-api"
	"github.com/kolibri7557/try-rest-api/pkg/repository"
)

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

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_HASH_SALT"))))
}
