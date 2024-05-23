package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kostylevdev/todo-rest-api/internal/domain"
)

type Autorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(domain.SignInUserInput) (domain.User, error)
	CreateSession(session domain.Session) (string, error)
	GetSession(refreshToken string) (domain.Session, error)
	DeleteSession(refreshToken string) error
}

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (int, error) {
	var id int
	checkId := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", usersTable)
	if err := r.db.Get(&id, checkId, user.Username); err != sql.ErrNoRows {
		if err != nil {
			return 0, err
		}
		return 0, errors.New("username already exists")
	}
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(signinuser domain.SignInUserInput) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id, name, username FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	if err := r.db.Get(&user, query, signinuser.Username, signinuser.Password); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *AuthPostgres) CreateSession(session domain.Session) (string, error) {
	var refreshToken string
	query := fmt.Sprintf("INSERT INTO %s (user_id, expires_at, ip) VALUES ($1, $2, $3) RETURNING refresh_token", sessionsTable)
	row := r.db.QueryRow(query, session.UserId, session.ExpiresAt, session.ClientIP)
	if err := row.Scan(&refreshToken); err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (r *AuthPostgres) GetSession(refreshToken string) (domain.Session, error) {
	var session domain.Session
	query := fmt.Sprintf("SELECT id, user_id, expires_at, ip FROM %s WHERE refresh_token=$1", sessionsTable)
	if err := r.db.Get(&session, query, refreshToken); err != nil {
		return domain.Session{}, err
	}
	return session, nil
}

func (r *AuthPostgres) DeleteSession(refreshToken string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE refresh_token=$1", sessionsTable)
	if _, err := r.db.Exec(query, refreshToken); err != nil {
		return err
	}
	return nil
}
