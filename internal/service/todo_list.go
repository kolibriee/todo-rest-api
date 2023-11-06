package service

import (
	"github.com/kostylevdev/todo-rest-api/internal/domain"
	"github.com/kostylevdev/todo-rest-api/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateList(userId int, list domain.TodoList) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *TodoListService) GetAllLists(userId int) ([]domain.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *TodoListService) GetListById(userId int, ListId int) (domain.TodoList, error) {
	return s.repo.GetListById(userId, ListId)
}

func (s *TodoListService) DeleteList(userId int, ListId int) error {
	return s.repo.DeleteList(userId, ListId)
}

func (s *TodoListService) UpdateList(userId int, ListId int, list domain.TodoListUpdate) error {
	if err := domain.ValidateTodoListUpdate(list); err != nil {
		return err
	}
	return s.repo.UpdateList(userId, ListId, list)
}
