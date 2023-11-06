package service

import (
	"github.com/kostylevdev/todo-rest-api/internal/domain"
	"github.com/kostylevdev/todo-rest-api/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	repoList repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, repoList repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		repoList: repoList}
}

func (s *TodoItemService) CreateItem(userId int, listId int, item domain.TodoItem) (int, error) {
	_, err := s.repoList.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.CreateItem(listId, item)
}

func (s *TodoItemService) GetAllItems(userId int, listId int) ([]domain.TodoItem, error) {
	return s.repo.GetAllItems(userId, listId)
}

func (s *TodoItemService) GetItemById(userId int, itemId int) (domain.TodoItem, error) {
	return s.repo.GetItemById(userId, itemId)
}

func (s *TodoItemService) DeleteItem(userId int, itemId int) error {
	return s.repo.DeleteItem(userId, itemId)
}

func (s *TodoItemService) UpdateItem(userId int, itemId int, item domain.TodoItemUpdate) error {
	if err := domain.ValidateTodoItemUpdate(item); err != nil {
		return err
	}
	return s.repo.UpdateItem(userId, itemId, item)
}
