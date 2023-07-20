package service

import "github.com/kolibri7557/try-rest-api/pkg/repository"

type Autorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service interface {
	Autorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
