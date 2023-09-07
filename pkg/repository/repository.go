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
	GetAllLists(userId int) ([]tryrest.TodoList, error)
	GetListById(userId int, id int) (tryrest.TodoList, error)
	DeleteList(userId int, id int) error
	UpdateList(userId int, id int, list tryrest.TodoListUpdate) error
}

type TodoItem interface {
	CreateItem(listId int, item tryrest.TodoItem) (int, error)
	GetAllItems(userId int, listId int) ([]tryrest.TodoItem, error)
	GetItemById(userId int, itemId int) (tryrest.TodoItem, error)
	DeleteItem(userId int, itemId int) error
	UpdateItem(userId int, itemId int, item tryrest.TodoItemUpdate) error
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
