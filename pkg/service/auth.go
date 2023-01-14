package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/aleks-tim/todo-app"
	"github.com/aleks-tim/todo-app/pkg/repository"
)

const salt = "7-he8m_rYjKD90"

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
