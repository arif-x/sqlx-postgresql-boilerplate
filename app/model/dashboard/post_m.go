package dashboard

import (
	"time"

	jsonutil "github.com/arif-x/sqlx-gofiber-boilerplate/pkg/json"
	"github.com/google/uuid"
)

type Post struct {
	UUID             uuid.UUID         `db:"uuid" json:"uuid"`
	PostCategoryUUID uuid.UUID         `db:"post_category_uuid" json:"post_category_uuid"`
	UserUUID         uuid.UUID         `db:"user_uuid" json:"user_uuid"`
	Title            string            `db:"title" json:"title"`
	Content          string            `db:"content" json:"content"`
	CreatedAt        time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt        *time.Time        `db:"updated_at" json:"updated_at"`
	DeletedAt        *time.Time        `db:"deleted_at" json:"deleted_at"`
	User             *jsonutil.JSONRaw `db:"user" json:"user"`
	PostCategory     *jsonutil.JSONRaw `db:"post_category" json:"post_category"`
}

type PostShow struct {
	UUID             uuid.UUID         `db:"uuid" json:"uuid"`
	PostCategoryUUID uuid.UUID         `db:"post_category_uuid" json:"post_category_uuid"`
	UserUUID         uuid.UUID         `db:"user_uuid" json:"user_uuid"`
	Title            string            `db:"title" json:"title"`
	Content          string            `db:"content" json:"content"`
	CreatedAt        time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt        *time.Time        `db:"updated_at" json:"updated_at"`
	User             *jsonutil.JSONRaw `db:"user" json:"user"`
	PostCategory     *jsonutil.JSONRaw `db:"post_category" json:"post_category"`
}

type StorePost struct {
	PostCategoryUUID uuid.UUID `json:"post_category_uuid" form:"post_category_uuid"`
	UserUUID         uuid.UUID `json:"user_uuid" form:"user_uuid"`
	Title            string    `json:"title" form:"title"`
	Content          string    `json:"content" form:"content"`
}

type UpdatePost struct {
	PostCategoryUUID uuid.UUID `json:"post_category_uuid" form:"post_category_uuid"`
	UserUUID         uuid.UUID `json:"user_uuid" form:"user_uuid"`
	Title            string    `json:"title" form:"title"`
	Content          string    `json:"content" form:"content"`
}
