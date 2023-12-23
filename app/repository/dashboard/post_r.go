package dashboard

import (
	"context"
	"fmt"
	"log"
	"time"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/dashboard"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	"github.com/google/uuid"
)

type PostRepository interface {
	Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.Post, int, error)
	Show(ID string) (model.PostShow, error)
	Store(model *model.StorePost) (model.Post, error)
	Update(ID string, request *model.UpdatePost) (model.Post, error)
	Destroy(ID string) (model.Post, error)
}

type PostRepo struct {
	db *database.DB
}

func (repo *PostRepo) Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.Post, int, error) {
	_select := `
	posts.uuid,
	user_uuid,
	post_category_uuid,
    title,
    content,
    posts.created_at,
    posts.updated_at,
	CASE 
    	WHEN users.id IS NULL THEN null
    	ELSE json_build_object(
        	'id', users.id,
        	'name', users.name,
        	'username', users.username,
        	'email', users.email,
        	'created_at', users.created_at,
        	'updated_at', users.updated_at
    	)
	END AS user,
	CASE 
    	WHEN post_categories.id IS NULL THEN null
		ELSE json_build_object(
        	'id', post_categories.id,
        	'name', post_categories.name,
        	'created_at', post_categories.created_at,
        	'updated_at', post_categories.updated_at
		)
	END AS post_category
	`
	_conditions := database.Search([]string{"title", "content", "users.name", "post_categories.name"}, search)
	_order := database.OrderBy("posts.id", sort)
	_limit := database.Limit(limit, offset)

	count_query := fmt.Sprintf(`SELECT count(*) FROM posts LEFT JOIN users ON users.uuid = posts.user_uuid LEFT JOIN post_categories ON post_categories.uuid = posts.post_category_uuid %s`, _conditions)
	var count int
	_ = repo.db.QueryRow(count_query).Scan(&count)

	query := fmt.Sprintf(`SELECT %s FROM posts LEFT JOIN users ON users.uuid = posts.user_uuid LEFT JOIN post_categories ON post_categories.uuid = posts.post_category_uuid %s %s %s`, _select, _conditions, _order, _limit)

	rows, err := repo.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()
	var items []model.Post
	for rows.Next() {
		var i model.Post
		// var UserJSON *string
		// var PostCategoryJSON *string
		err := rows.Scan(
			&i.UUID,
			&i.PostCategoryUUID,
			&i.UserUUID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.User,
			&i.PostCategory,
		)
		if err != nil {
			log.Fatal(err)
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
	posts.uuid,
	user_uuid,
	post_category_uuid,
	title,
	content,
	posts.created_at,
	posts.updated_at,
	CASE 
    	WHEN users.id IS NULL THEN null
    	ELSE json_build_object(
        	'id', users.id,
        	'name', users.name,
        	'username', users.username,
        	'email', users.email,
        	'created_at', users.created_at,
        	'updated_at', users.updated_at
    	)
	END AS user,
	CASE 
    	WHEN post_categories.id IS NULL THEN null
		ELSE json_build_object(
        	'id', post_categories.id,
        	'name', post_categories.name,
        	'created_at', post_categories.created_at,
        	'updated_at', post_categories.updated_at
		)
	END AS post_category
	FROM posts LEFT JOIN users ON users.uuid = posts.user_uuid LEFT JOIN post_categories ON post_categories.uuid = posts.post_category_uuid
	WHERE posts.id = $1 LIMIT 1
	`

	err := repo.db.QueryRowContext(context.Background(), query, ID).Scan(
		&post.UUID,
		&post.PostCategoryUUID,
		&post.UserUUID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.User,
		&post.PostCategory,
	)

	return post, err
}

func (repo *PostRepo) Store(request *model.StorePost) (model.Post, error) {
	query := `INSERT INTO "posts" (uuid, post_category_uuid, user_uuid, title, content, created_at) VALUES($1, $2, $3, $4, $5, $6) 
	RETURNING uuid, post_category_uuid, user_uuid, title, content, created_at`
	var post model.Post
	err := repo.db.QueryRowContext(context.Background(), query, uuid.New(), request.PostCategoryUUID, request.UserUUID, request.Title, request.Content, time.Now()).Scan(
		&post.UUID,
		&post.PostCategoryUUID,
		&post.UserUUID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
	)
	if err != nil {
		return model.Post{}, err
	}
	return post, err
}

func (repo *PostRepo) Update(ID string, request *model.UpdatePost) (model.Post, error) {
	query := `UPDATE "posts" SET post_category_uuid = $2, user_uuid = $3, title = $4, content = $5, updated_at = $6 WHERE uuid = $1 
		RETURNING uuid, post_category_uuid, user_uuid, title, content, created_at, updated_at`
	var post model.Post
	err := repo.db.QueryRowContext(context.Background(), query, ID, request.PostCategoryUUID, request.UserUUID, request.Title, request.Content, time.Now()).Scan(
		&post.UUID,
		&post.PostCategoryUUID,
		&post.UserUUID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return model.Post{}, err
	}
	return post, err
}

func (repo *PostRepo) Destroy(ID string) (model.Post, error) {
	query := `UPDATE "posts" SET updated_at = $2, deleted_at = $3 WHERE uuid = $1 
	RETURNING uuid, post_category_uuid, user_uuid, title, content, created_at, updated_at, deleted_at`
	var post model.Post
	err := repo.db.QueryRowContext(context.Background(), query, ID, time.Now(), time.Now()).Scan(
		&post.UUID,
		&post.PostCategoryUUID,
		&post.UserUUID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.DeletedAt,
	)
	if err != nil {
		return model.Post{}, err
	}
	return post, err
}

func NewPostRepo(db *database.DB) PostRepository {
	return &PostRepo{db}
}
