package service

import (
	tryrest "github.com/kolibri7557/try-rest-api"
	"github.com/kolibri7557/try-rest-api/pkg/repository"
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

func (s *TodoItemService) CreateItem(userId int, listId int, item tryrest.TodoItem) (int, error) {
	_, err := s.repoList.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.CreateItem(listId, item)
}
