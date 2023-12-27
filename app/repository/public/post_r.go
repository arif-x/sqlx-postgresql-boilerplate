package public

import (
	"context"
	"fmt"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/public"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
)

type PostRepository interface {
	Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.Post, int, error)
	PostCategoryPost(uuid string, limit int, offset uint, search string, sort_by string, sort string) (model.PostCategoryWithPost, int, error)
	UserPost(uuid string, limit int, offset uint, search string, sort_by string, sort string) (model.UserWithPost, int, error)
	Show(UUID string) (model.Post, error)
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
			return nil, 0, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, 0, err
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return items, count, nil
}

func (repo *PostRepo) PostCategoryPost(uuid string, limit int, offset uint, search string, sort_by string, sort string) (model.PostCategoryWithPost, int, error) {
	_limit := database.Limit(limit, offset)
	_select := fmt.Sprintf(`
	post_categories.uuid,
    post_categories.name,
    COALESCE(
        (
            SELECT json_agg(
                json_build_object(
                    'uuid', posts.uuid,
                    'post_category_uuid', posts.post_category_uuid,
                    'user_uuid', posts.user_uuid,
                    'title', posts.title,
                    'content', posts.content,
                    'user', CASE 
                        WHEN user_uuid_u IS NULL THEN null
                            ELSE json_build_object(
                                    'uuid', user_uuid_u,
                                    'name', user_name,
                                    'username', user_username,
                                    'email', user_email,
                                    'created_at', user_created_at,
                                    'updated_at', user_updated_at
                            )
                        END,
                    'post_category', CASE 
                        WHEN post_category_uuid_u IS NULL THEN null
                            ELSE json_build_object(
                                    'uuid', post_category_uuid_u,
                                    'name', post_category_name,
                                    'created_at', post_category_created_at,
                                    'updated_at', post_category_updated_at
                            )
                        END
                )
            ) 
            FROM (SELECT 
            	posts.*, 
            	users.uuid as user_uuid_u, users.name as user_name, users.username as user_username, users.email as user_email, users.created_at as user_created_at, users.updated_at as user_updated_at, 
            	post_categories.uuid as post_category_uuid_u, post_categories.name as post_category_name, post_categories.created_at as post_category_created_at, post_categories.updated_at as post_category_updated_at 
            	FROM posts
	            LEFT JOIN users ON users.uuid = posts.user_uuid
	            LEFT JOIN post_categories ON post_categories.uuid = posts.post_category_uuid
	            WHERE posts.post_category_uuid = $1
	            %s
            ) posts
        ), '[]'
	) AS post,
    post_categories.created_at,
    post_categories.updated_at
	`, _limit)

	count_query := `SELECT count(*) FROM posts WHERE post_category_uuid = $1`
	var count int
	_ = repo.db.QueryRow(count_query, uuid).Scan(&count)

	query := fmt.Sprintf(`SELECT %s FROM post_categories WHERE post_categories.uuid = $1`, _select)

	var items model.PostCategoryWithPost

	err := repo.db.QueryRowContext(context.Background(), query, uuid).Scan(
		&items.UUID,
		&items.Name,
		&items.Post,
		&items.CreatedAt,
		&items.UpdatedAt,
	)

	if err != nil {
		return model.PostCategoryWithPost{}, 0, err
	}

	return items, count, nil
}

func (repo *PostRepo) UserPost(uuid string, limit int, offset uint, search string, sort_by string, sort string) (model.UserWithPost, int, error) {
	_limit := database.Limit(limit, offset)
	_select := fmt.Sprintf(`
	users.uuid,
    users.name,
    users.username,
    users.email,
    COALESCE(
        (
            SELECT json_agg(
                json_build_object(
                    'uuid', posts.uuid,
                    'post_category_uuid', posts.post_category_uuid,
                    'user_uuid', posts.user_uuid,
                    'title', posts.title,
                    'content', posts.content,
                    'user', CASE 
                        WHEN user_uuid_u IS NULL THEN null
                            ELSE json_build_object(
                                    'uuid', user_uuid_u,
                                    'name', user_name,
                                    'username', user_username,
                                    'email', user_email,
                                    'created_at', user_created_at,
                                    'updated_at', user_updated_at
                            )
                        END,
                    'post_category', CASE 
                        WHEN post_category_uuid_u IS NULL THEN null
                            ELSE json_build_object(
                                    'uuid', post_category_uuid_u,
                                    'name', post_category_name,
                                    'created_at', post_category_created_at,
                                    'updated_at', post_category_updated_at
                            )
                        END
                )
            ) 
            FROM (SELECT 
            	posts.*, 
            	users.uuid as user_uuid_u, users.name as user_name, users.username as user_username, users.email as user_email, users.created_at as user_created_at, users.updated_at as user_updated_at, 
            	post_categories.uuid as post_category_uuid_u, post_categories.name as post_category_name, post_categories.created_at as post_category_created_at, post_categories.updated_at as post_category_updated_at 
            	FROM posts
	            LEFT JOIN users ON users.uuid = posts.user_uuid
	            LEFT JOIN post_categories ON post_categories.uuid = posts.post_category_uuid 
	            WHERE posts.user_uuid = $1
	            %s
            ) posts
        ), '[]'
	) AS post,
    users.created_at,
    users.updated_at
	`, _limit)

	count_query := `SELECT count(*) FROM posts WHERE user_uuid = $1`
	var count int
	_ = repo.db.QueryRow(count_query, uuid).Scan(&count)

	query := fmt.Sprintf(`SELECT %s FROM users WHERE users.uuid = $1`, _select)

	var items model.UserWithPost

	err := repo.db.QueryRowContext(context.Background(), query, uuid).Scan(
		&items.UUID,
		&items.Name,
		&items.Username,
		&items.Email,
		&items.Post,
		&items.CreatedAt,
		&items.UpdatedAt,
	)

	if err != nil {
		return model.UserWithPost{}, 0, err
	}

	return items, count, nil
}

func (repo *PostRepo) Show(UUID string) (model.Post, error) {
	var post model.Post
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
    	WHEN users.uuid IS NULL THEN null
    	ELSE json_build_object(
        	'uuid', users.uuid,
        	'name', users.name,
        	'username', users.username,
        	'email', users.email,
        	'created_at', users.created_at,
        	'updated_at', users.updated_at
    	)
	END AS user,
	CASE 
    	WHEN post_categories.uuid IS NULL THEN null
		ELSE json_build_object(
        	'uuid', post_categories.uuid,
        	'name', post_categories.name,
        	'created_at', post_categories.created_at,
        	'updated_at', post_categories.updated_at
		)
	END AS post_category
	FROM posts LEFT JOIN users ON users.uuid = posts.user_uuid LEFT JOIN post_categories ON post_categories.uuid = posts.post_category_uuid
	WHERE posts.uuid = $1 LIMIT 1
	`

	err := repo.db.QueryRowContext(context.Background(), query, UUID).Scan(
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

func NewPostRepo(db *database.DB) PostRepository {
	return &PostRepo{db}
}
