package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kostylevdev/todo-rest-api/internal/config"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	todoItemsTable  = "todo_items"
	usersListsTable = "users_lists"
	listsItemsTable = "lists_items"
)

func NewPostgresDB(cfg *config.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
