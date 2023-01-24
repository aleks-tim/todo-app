package service

import (
	"github.com/aleks-tim/todo-app"
	"github.com/aleks-tim/todo-app/pkg/repository"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParceToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Autorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(*repo),
	}
}
