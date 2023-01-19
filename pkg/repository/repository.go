package repository

import (
	"github.com/aleks-tim/todo-app"
	"github.com/jmoiron/sqlx"
	// "github.com/sirupsen/logrus"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Autorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	// logrus.Println("NewRepository()")

	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}
