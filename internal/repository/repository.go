package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kostylevdev/todo-rest-api/internal/domain"
)

type Autorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(domain.SignInUser) (domain.User, error)
}

type TodoList interface {
	CreateList(userId int, list domain.TodoList) (int, error)
	GetAllLists(userId int) ([]domain.TodoList, error)
	GetListById(userId int, id int) (domain.TodoList, error)
	DeleteList(userId int, id int) error
	UpdateList(userId int, id int, list domain.TodoListUpdate) error
}

type TodoItem interface {
	CreateItem(listId int, item domain.TodoItem) (int, error)
	GetAllItems(userId int, listId int) ([]domain.TodoItem, error)
	GetItemById(userId int, itemId int) (domain.TodoItem, error)
	DeleteItem(userId int, itemId int) error
	UpdateItem(userId int, itemId int, item domain.TodoItemUpdate) error
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
		TodoItem:     NewTodoItemPostgres(db),
	}
}
