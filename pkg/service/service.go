package service

import (
	"github.com/aleks-tim/todo-app"
	"github.com/aleks-tim/todo-app/pkg/repository"
	// "github.com/sirupsen/logrus"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
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
	// logrus.Println("NewService()")

	return &Service{
		Autorization: NewAuthService(*repo),
	}
}
