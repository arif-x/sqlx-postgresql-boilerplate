package dashboard

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/dashboard"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
)

type PostRepository interface {
	Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.Post, int, error)
	Show(ID string) (model.PostShow, error)
}

type PostRepo struct {
	db *database.DB
}

func (repo *PostRepo) Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.Post, int, error) {
	_select := `
	posts.id,
	user_id,
	post_category_id,
    title,
    content,
    posts.created_at,
    posts.updated_at,
	json_build_object(
        'id', users.id,
        'name', users.name,
        'username', users.username,
        'email', users.email,
        'created_at', users.created_at,
        'updated_at', users.updated_at
    )::jsonb as user,
	json_build_object(
        'id', post_categories.id,
        'name', post_categories.name,
        'created_at', post_categories.created_at,
        'updated_at', post_categories.updated_at
    )::jsonb as post_category
	`
	_conditions := database.Search([]string{"title", "content", "users.name", "post_categories.name"}, search)
	_order := database.OrderBy("posts.created_at", sort)
	_limit := database.Limit(limit, offset)

	count_query := fmt.Sprintf(`SELECT count(*) FROM posts LEFT JOIN users ON users.id = posts.user_id LEFT JOIN post_categories ON post_categories.id = posts.post_category_id %s`, _conditions)
	var count int
	_ = repo.db.QueryRow(count_query).Scan(&count)

	query := fmt.Sprintf(`SELECT %s FROM posts LEFT JOIN users ON users.id = posts.user_id LEFT JOIN post_categories ON post_categories.id = posts.post_category_id %s %s %s`, _select, _conditions, _order, _limit)

	rows, err := repo.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()
	var items []model.Post
	for rows.Next() {
		var i model.Post
		if err := rows.Scan(
			&i.ID,
			&i.PostCategoryID,
			&i.UserID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.User,
			&i.PostCategory,
		); err != nil {
			return nil, 0, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		log.Fatal(err)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return items, count, nil
}

func (repo *PostRepo) Show(ID string) (model.PostShow, error) {
	var post model.PostShow
	query := `
	SELECT 
	posts.id,
	user_id,
	post_category_id,
	title,
	content,
	posts.created_at,
	posts.updated_at,
	json_build_object(
		'id', users.id,
		'name', users.name,
		'username', users.username,
		'email', users.email,
		'created_at', users.created_at,
		'updated_at', users.updated_at
	)::jsonb as user,
	json_build_object(
		'id', post_categories.id,
		'name', post_categories.name,
		'created_at', post_categories.created_at,
		'updated_at', post_categories.updated_at
	)::jsonb as post_category
	FROM posts LEFT JOIN users ON users.id = posts.user_id LEFT JOIN post_categories ON post_categories.id = posts.post_category_id
	WHERE posts.id = $1 LIMIT 1
	`

	var UserJSON string
	var PostCategoryJSON string

	err := repo.db.QueryRowContext(context.Background(), query, ID).Scan(
		&post.ID,
		&post.PostCategoryID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
		&UserJSON,
		&PostCategoryJSON,
	)

	if err != nil {
		return model.PostShow{}, err
	}

	var user model.User
	if err := json.Unmarshal([]byte(UserJSON), &user); err != nil {
		return model.PostShow{}, err
	}
	post.User = user

	var PostCategory model.PostCategory
	if err := json.Unmarshal([]byte(PostCategoryJSON), &PostCategory); err != nil {
		return model.PostShow{}, err
	}
	post.PostCategory = PostCategory

	return post, err
}

func NewPostRepo(db *database.DB) PostRepository {
	return &PostRepo{db}
}
