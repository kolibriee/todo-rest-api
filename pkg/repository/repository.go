package repository

import (
	"github.com/jmoiron/sqlx"
	tryrest "github.com/kolibri7557/try-rest-api"
)

type Autorization interface {
	CreateUser(user tryrest.User) (int, error)
	GetUser(tryrest.SignInUser) (tryrest.User, error)
}

type TodoList interface {
	CreateList(userId int, list tryrest.TodoList) (int, error)
}

type TodoItem interface {
}

type Repository struct {
	Autorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
		TodoList:     NewTodoListPostgres(db),
	}
}
