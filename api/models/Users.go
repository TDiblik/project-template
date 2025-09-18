package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UsersModelDB struct {
	Id            uuid.UUID     `db:"id"`
	Email         string        `db:"email"`
	EmailVerified bool          `db:"email_verified"`
	PasswordHash  SQLNullString `db:"password_hash"`

	FirstName SQLNullString `db:"first_name"`
	LastName  SQLNullString `db:"last_name"`

	Handle    SQLNullString `db:"handle"`
	AvatarUrl SQLNullString `db:"avatar_url"`
	Active    bool          `db:"active"`

	GithubId     SQLNullString `db:"github_id"`
	GithubHandle SQLNullString `db:"github_handle"`
	GithubUrl    SQLNullString `db:"github_url"`

	GoogleId SQLNullString `db:"google_id"`

	FacebookId  SQLNullString `db:"facebook_id"`
	FacebookUrl SQLNullString `db:"facebook_url"`

	SpotifyId  SQLNullString `db:"spotify_id"`
	SpotifyUrl SQLNullString `db:"spotify_url"`

	LastLoginAt sql.NullTime `db:"last_login_at"`

	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}
