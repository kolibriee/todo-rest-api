package repository

import "github.com/jmoiron/sqlx"

type Autorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository interface {
	Autorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
