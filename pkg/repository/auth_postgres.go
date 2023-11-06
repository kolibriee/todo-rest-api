package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	domain "github.com/kostylevdev/todo-rest-api"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(signinuser domain.SignInUser) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id, name, username FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	if err := r.db.Get(&user, query, signinuser.Username, signinuser.Password); err != nil {
		return domain.User{}, err
	}
	return user, nil
}
