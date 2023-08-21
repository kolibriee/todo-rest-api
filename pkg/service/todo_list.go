package service

import (
	tryrest "github.com/kolibri7557/try-rest-api"
	"github.com/kolibri7557/try-rest-api/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateList(userId int, list tryrest.TodoList) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *TodoListService) GetAllLists(userId int) ([]tryrest.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *TodoListService) GetByIdList(userId int, ListId int) (tryrest.TodoList, error) {
	return s.repo.GetByIdList(userId, ListId)
}
