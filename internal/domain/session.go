package domain

import "time"

type Session struct {
	Id           int       `db:"id"`
	UserId       int       `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	ExpiresAt    time.Time `db:"expires_at"`
	ClientIP     string    `db:"ip"`
}
