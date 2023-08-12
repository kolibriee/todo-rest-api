package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	tryrest "github.com/kolibri7557/try-rest-api"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user tryrest.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(signinuser tryrest.SignInUser) (tryrest.User, error) {
	var user tryrest.User
	query := fmt.Sprintf("SELECT id, name, username FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	row := r.db.QueryRow(query, signinuser.Username, signinuser.Password)
	if err := row.Scan(&user.Id, &user.Name, &user.Username); err != nil {
		return tryrest.User{}, err
	}
	return user, nil
}
