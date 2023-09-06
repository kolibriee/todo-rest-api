package tryrest

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description " db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type TodoListUpdate struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func ValidateTodoListUpdate(list TodoListUpdate) error {
	if list.Title == nil && list.Description == nil {
		return errors.New("update must have title or description")
	}
	return nil
}
