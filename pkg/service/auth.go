package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
	"github.com/vitalik-ez/Chat-Golang/pkg/repository"
)

const solt = "asdasdqwddsfaasdasdqw"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}
