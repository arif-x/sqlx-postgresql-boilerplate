package dashboard

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	UserID    string     `db:"user_id" json:"user_id"`
	Title     string     `db:"title" json:"title"`
	Content   string     `db:"content" json:"content"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type CreatePost struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdatePost struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}
