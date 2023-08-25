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

func (s *TodoListService) GetListById(userId int, ListId int) (tryrest.TodoList, error) {
	return s.repo.GetListById(userId, ListId)
}

func (s *TodoListService) DeleteList(userId int, ListId int) error {
	return s.repo.DeleteList(userId, ListId)
}

func (s *TodoListService) UpdateList(userId int, ListId int, list tryrest.TodoListUpdate) error {
	if err := tryrest.ValidateTodoListUpdate(list); err != nil {
		return err
	}
	return s.repo.UpdateList(userId, ListId, list)
}
