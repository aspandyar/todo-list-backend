package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/aspandyar/todo-list/internal/models"
	"github.com/aspandyar/todo-list/pkg/repository"
)

const salt = "ajhdajsdh1lk23m1j3n1"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
