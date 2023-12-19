package dashboard

import (
	"time"

	"github.com/arif-x/sqlx-gofiber-boilerplate/app/trait"
	"github.com/google/uuid"
)

type Post struct {
	ID             uuid.UUID     `db:"id" json:"id"`
	PostCategoryID string        `db:"post_category_id" json:"post_category_id"`
	UserID         string        `db:"user_id" json:"user_id"`
	Title          string        `db:"title" json:"title"`
	Content        string        `db:"content" json:"content"`
	CreatedAt      time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt      *time.Time    `db:"updated_at" json:"updated_at"`
	User           trait.JSONRaw `db:"user" json:"user"`
	PostCategory   trait.JSONRaw `db:"post_category" json:"post_category"`
}

type PostShow struct {
	ID             uuid.UUID    `db:"id" json:"id"`
	PostCategoryID string       `db:"post_category_id" json:"post_category_id"`
	UserID         string       `db:"user_id" json:"user_id"`
	Title          string       `db:"title" json:"title"`
	Content        string       `db:"content" json:"content"`
	CreatedAt      time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt      *time.Time   `db:"updated_at" json:"updated_at"`
	User           User         `db:"user" json:"user"`
	PostCategory   PostCategory `db:"post_category" json:"post_category"`
}

type StorePost struct {
	ID             string `json:"id"`
	PostCategoryID string `json:"post_category_id"`
	UserID         string `json:"user_id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}

type UpdatePost struct {
	ID             string `json:"id"`
	PostCategoryID string `json:"post_category_id"`
	UserID         string `json:"user_id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}
