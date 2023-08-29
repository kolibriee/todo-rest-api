package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	tryrest "github.com/kolibri7557/try-rest-api"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) CreateItem(listId int, item tryrest.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var ItemId int
	queryCreateToDoItem := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(queryCreateToDoItem, item.Title, item.Description)
	if err := row.Scan(&ItemId); err != nil {
		tx.Rollback()
		return 0, err
	}
	queryCreateListsItems := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(queryCreateListsItems, listId, ItemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}
	return ItemId, nil
}
