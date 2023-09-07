package service

import (
	tryrest "github.com/kolibri7557/try-rest-api"
	"github.com/kolibri7557/try-rest-api/pkg/repository"
)

type Autorization interface {
	CreateUser(user tryrest.User) (int, error)
	GenerateToken(signinuser tryrest.SignInUser) (int, string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateList(userID int, list tryrest.TodoList) (int, error)
	GetAllLists(userID int) ([]tryrest.TodoList, error)
	GetListById(userID int, id int) (tryrest.TodoList, error)
	DeleteList(userID int, id int) error
	UpdateList(userID int, id int, list tryrest.TodoListUpdate) error
}

type TodoItem interface {
	CreateItem(userID int, listID int, item tryrest.TodoItem) (int, error)
	GetAllItems(userID int, listID int) ([]tryrest.TodoItem, error)
	GetItemById(userID int, itemId int) (tryrest.TodoItem, error)
}

type Service struct {
	Autorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		TodoList:     NewTodoListService(repos.TodoList),
		TodoItem:     NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
