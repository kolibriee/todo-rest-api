package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	tryrest "github.com/kolibri7557/try-rest-api"
)

type todoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *todoListPostgres {
	return &todoListPostgres{db: db}
}

func (r *todoListPostgres) CreateList(userId int, list tryrest.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var idList int
	queryCreateToDoList := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(queryCreateToDoList, list.Title, list.Description)
	if err := row.Scan(&idList); err != nil {
		tx.Rollback()
		return 0, err
	}
	queryCreateUserLists := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(queryCreateUserLists, userId, idList)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}
	return idList, nil
}

func (r *todoListPostgres) GetAllLists(userId int) ([]tryrest.TodoList, error) {
	var lists []tryrest.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1", todoListsTable, usersListsTable)
	if err := r.db.Select(&lists, query, userId); err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *todoListPostgres) GetListById(userId int, ListId int) (tryrest.TodoList, error) {
	var list tryrest.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListsTable)
	if err := r.db.Get(&list, query, userId, ListId); err != nil {
		return tryrest.TodoList{}, err
	}
	return list, nil
}

func (r *todoListPostgres) DeleteList(userId int, ListId int) error {
	query := fmt.Sprintf("DELETE FROM %s t1 USING %s u1 WHERE t1.id = u1.list_id AND u1.user_id = $1 AND u1.list_id = $2", todoListsTable, usersListsTable)
	exec, err := r.db.Exec(query, userId, ListId)
	if err != nil {
		return err
	}
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("list not found")
	}
	return nil
}
