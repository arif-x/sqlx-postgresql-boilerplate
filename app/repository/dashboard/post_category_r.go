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

type PostCategoryRepository interface {
	Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.PostCategory, int, error)
	Show(ID string) (model.PostCategoryShow, error)
	Store(model *model.StorePostCategory) (model.PostCategory, error)
	Update(ID string, request *model.UpdatePostCategory) (model.PostCategory, error)
	Destroy(ID string) (model.PostCategory, error)
}

type PostCategoryRepo struct {
	db *database.DB
}

func (repo *PostCategoryRepo) Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.PostCategory, int, error) {
	_select := "id, name, created_at, updated_at, deleted_at"
	_conditions := database.Search([]string{"name"}, search)
	_order := database.OrderBy(sort_by, sort)
	_limit := database.Limit(limit, offset)

	count_query := fmt.Sprintf(`SELECT count(*) FROM post_categories %s`, _conditions)
	var count int
	_ = repo.db.QueryRow(count_query).Scan(&count)

	query := fmt.Sprintf(`SELECT %s FROM post_categories %s %s %s`, _select, _conditions, _order, _limit)

	rows, err := repo.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()
	var items []model.PostCategory
	for rows.Next() {
		var i model.PostCategory
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

func (repo *PostCategoryRepo) Show(ID string) (model.PostCategoryShow, error) {
	var postCategory model.PostCategoryShow
	query := "SELECT id, name, created_at, updated_at, deleted_at FROM post_categories WHERE id = $1 LIMIT 1"
	err := repo.db.QueryRowContext(context.Background(), query, ID).Scan(
		&postCategory.ID,
		&postCategory.Name,
		&postCategory.CreatedAt,
		&postCategory.UpdatedAt,
		&postCategory.DeletedAt,
	)
	if err != nil {
		return model.PostCategoryShow{}, err
	}
	return postCategory, err
}

func (repo *PostCategoryRepo) Store(request *model.StorePostCategory) (model.PostCategory, error) {
	query := `INSERT INTO "post_categories" (id, name, created_at) VALUES($1, $2, $3) 
	RETURNING id, name, created_at`
	var postCategory model.PostCategory
	err := repo.db.QueryRowContext(context.Background(), query, uuid.New(), request.Name, time.Now()).Scan(
		&postCategory.ID,
		&postCategory.Name,
		&postCategory.CreatedAt,
	)
	if err != nil {
		return model.PostCategory{}, err
	}
	return postCategory, err
}

func (repo *PostCategoryRepo) Update(ID string, request *model.UpdatePostCategory) (model.PostCategory, error) {
	query := `UPDATE "post_categories" SET name = $2, updated_at = $3 WHERE id = $1 
	RETURNING id, name, created_at, updated_at`
	var postCategory model.PostCategory
	err := repo.db.QueryRowContext(context.Background(), query, ID, request.Name, time.Now()).Scan(
		&postCategory.ID,
		&postCategory.Name,
		&postCategory.CreatedAt,
		&postCategory.UpdatedAt,
	)
	if err != nil {
		return model.PostCategory{}, err
	}
	return postCategory, err
}

func (repo *PostCategoryRepo) Destroy(ID string) (model.PostCategory, error) {
	query := `UPDATE "post_categories" SET updated_at = $2, deleted_at = $3 WHERE id = $1 
	RETURNING id, name, created_at, updated_at, deleted_at`
	var postCategory model.PostCategory
	err := repo.db.QueryRowContext(context.Background(), query, ID, time.Now(), time.Now()).Scan(
		&postCategory.ID,
		&postCategory.Name,
		&postCategory.CreatedAt,
		&postCategory.UpdatedAt,
		&postCategory.DeletedAt,
	)
	if err != nil {
		return model.PostCategory{}, err
	}
	return postCategory, err
}

func NewPostCategoryRepo(db *database.DB) PostCategoryRepository {
	return &PostCategoryRepo{db}
}
