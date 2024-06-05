package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kostylevdev/todo-rest-api/internal/domain"
	"github.com/sirupsen/logrus"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) CreateItem(listId int, item domain.TodoItemCreate) (int, error) {
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

func (r *TodoItemPostgres) GetAllItems(userId int, listId int) ([]domain.TodoItem, error) {
	var items []domain.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON li.list_id = ul.list_id WHERE ul.list_id = $1 AND ul.user_id = $2", todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TodoItemPostgres) GetItemById(userId int, itemId int) (domain.TodoItem, error) {
	var item domain.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id INNER JOIN %s ul ON li.list_id = ul.list_id WHERE li.item_id = $1 AND ul.user_id = $2", todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return domain.TodoItem{}, err
	}
	return item, nil
}

func (r *TodoItemPostgres) DeleteItem(userId int, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND li.item_id = $2", todoItemsTable, listsItemsTable, usersListsTable)
	exec, err := r.db.Exec(query, userId, itemId)
	if err != nil {
		return err
	}
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("item not found")
	}
	return nil
}

func (r *TodoItemPostgres) UpdateItem(userId int, itemId int, item domain.TodoItemUpdate) error {
	var setValues []string
	var args []interface{}
	argsId := 1
	if item.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, item.Title)
		argsId++
	}
	if item.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argsId))
		args = append(args, item.Description)
		argsId++
	}
	if item.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argsId))
		args = append(args, item.Done)
		argsId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND li.item_id = $%d",
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argsId, argsId+1)
	args = append(args, userId, itemId)
	logrus.Debugf("query: %s", query)
	logrus.Debugf("args: %v", args)
	exec, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("item not found")
	}
	return nil
}
