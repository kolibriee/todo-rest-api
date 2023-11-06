package service

import (
	"github.com/kostylevdev/todo-rest-api/internal/domain"
	"github.com/kostylevdev/todo-rest-api/internal/repository"
)

type Autorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(signinuser domain.SignInUser) (int, string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateList(userID int, list domain.TodoList) (int, error)
	GetAllLists(userID int) ([]domain.TodoList, error)
	GetListById(userID int, id int) (domain.TodoList, error)
	DeleteList(userID int, id int) error
	UpdateList(userID int, id int, list domain.TodoListUpdate) error
}

type TodoItem interface {
	CreateItem(userID int, listID int, item domain.TodoItem) (int, error)
	GetAllItems(userID int, listID int) ([]domain.TodoItem, error)
	GetItemById(userID int, itemId int) (domain.TodoItem, error)
	DeleteItem(userID int, itemId int) error
	UpdateItem(userID int, itemId int, item domain.TodoItemUpdate) error
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
