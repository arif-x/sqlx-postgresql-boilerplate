package dashboard

import (
	"context"
	"fmt"
	"time"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/dashboard"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	"github.com/google/uuid"
)

type PermissionRepository interface {
	Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.Permission, int, error)
	Show(UUID string) (model.ShowPermission, error)
	Store(model *model.StorePermission) (model.Permission, error)
	Update(UUID string, request *model.UpdatePermission) (model.Permission, error)
	Destroy(UUID string) (model.Permission, error)
}

type PermissionRepo struct {
	db *database.DB
}

func (repo *PermissionRepo) Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.Permission, int, error) {
	_select := "uuid, name, created_at, updated_at, deleted_at"
	_conditions := database.Search([]string{"name"}, search, "permissions.deleted_at")
	_order := database.OrderBy(sort_by, sort)
	_limit := database.Limit(limit, offset)

	count_query := fmt.Sprintf(`SELECT count(*) FROM permissions %s`, _conditions)
	var count int
	_ = repo.db.QueryRow(count_query).Scan(&count)

	query := fmt.Sprintf(`SELECT %s FROM permissions %s %s %s`, _select, _conditions, _order, _limit)

	rows, err := repo.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()
	var items []model.Permission
	for rows.Next() {
		var i model.Permission
		if err := rows.Scan(
			&i.UUID,
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
		return nil, 0, err
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return items, count, nil
}

func (repo *PermissionRepo) Show(UUID string) (model.ShowPermission, error) {
	var postCategory model.ShowPermission
	query := "SELECT uuid, name, created_at, updated_at, deleted_at FROM permissions WHERE uuid = $1 AND permissions.deleted_at IS NULL LIMIT 1"
	err := repo.db.QueryRowContext(context.Background(), query, UUID).Scan(
		&postCategory.UUID,
		&postCategory.Name,
		&postCategory.CreatedAt,
		&postCategory.UpdatedAt,
		&postCategory.DeletedAt,
	)
	if err != nil {
		return model.ShowPermission{}, err
	}
	return postCategory, err
}

func (repo *PermissionRepo) Store(request *model.StorePermission) (model.Permission, error) {
	query := `INSERT INTO "permissions" (uuid, name, created_at) VALUES($1, $2, $3) 
	RETURNING uuid, name, created_at`
	var postCategory model.Permission
	err := repo.db.QueryRowContext(context.Background(), query, uuid.New(), request.Name, time.Now()).Scan(
		&postCategory.UUID,
		&postCategory.Name,
		&postCategory.CreatedAt,
	)
	if err != nil {
		return model.Permission{}, err
	}
	return postCategory, err
}

func (repo *PermissionRepo) Update(UUID string, request *model.UpdatePermission) (model.Permission, error) {
	query := `UPDATE "permissions" SET name = $2, updated_at = $3 WHERE uuid = $1 
	RETURNING uuid, name, created_at, updated_at`
	var postCategory model.Permission
	err := repo.db.QueryRowContext(context.Background(), query, UUID, request.Name, time.Now()).Scan(
		&postCategory.UUID,
		&postCategory.Name,
		&postCategory.CreatedAt,
		&postCategory.UpdatedAt,
	)
	if err != nil {
		return model.Permission{}, err
	}
	return postCategory, err
}

func (repo *PermissionRepo) Destroy(UUID string) (model.Permission, error) {
	query := `UPDATE "permissions" SET updated_at = $2, deleted_at = $3 WHERE uuid = $1 
	RETURNING uuid, name, created_at, updated_at, deleted_at`
	var postCategory model.Permission
	err := repo.db.QueryRowContext(context.Background(), query, UUID, time.Now(), time.Now()).Scan(
		&postCategory.UUID,
		&postCategory.Name,
		&postCategory.CreatedAt,
		&postCategory.UpdatedAt,
		&postCategory.DeletedAt,
	)
	if err != nil {
		return model.Permission{}, err
	}
	return postCategory, err
}

func NewPermissionRepo(db *database.DB) PermissionRepository {
	return &PermissionRepo{db}
}
