package repository

type Autorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository interface {
	Autorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
