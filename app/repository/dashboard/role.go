package dashboard

import (
	"context"
	"fmt"
	"time"

	model "github.com/arif-x/sqlx-postgresql-boilerplate/app/model/dashboard"
	"github.com/arif-x/sqlx-postgresql-boilerplate/pkg/database"
	"github.com/google/uuid"
)

type RoleRepository interface {
	Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.Role, int, error)
	Show(UUID string) (model.ShowRole, error)
	Store(model *model.StoreRole) (model.Role, error)
	Update(UUID string, request *model.UpdateRole) (model.Role, error)
	Destroy(UUID string) (model.Role, error)
}

type RoleRepo struct {
	db *database.DB
}

func (repo *RoleRepo) Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.Role, int, error) {
	_select := "uuid, name, is_active, created_at, updated_at, deleted_at"
	_conditions := database.Search([]string{"name"}, search, "roles.deleted_at")
	_order := database.OrderBy(sort_by, sort)
	_limit := database.Limit(limit, offset)

	count_query := fmt.Sprintf(`SELECT count(*) FROM roles %s`, _conditions)
	var count int
	_ = repo.db.QueryRow(count_query).Scan(&count)

	query := fmt.Sprintf(`SELECT %s FROM roles %s %s %s`, _select, _conditions, _order, _limit)

	rows, err := repo.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()
	var items []model.Role
	for rows.Next() {
		var i model.Role
		if err := rows.Scan(
			&i.UUID,
			&i.Name,
			&i.IsActive,
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

func (repo *RoleRepo) Show(UUID string) (model.ShowRole, error) {
	var role model.ShowRole
	query := "SELECT uuid, name, is_active, created_at, updated_at, deleted_at FROM roles WHERE uuid = $1 AND roles.deleted_at IS NULL LIMIT 1"
	err := repo.db.QueryRowContext(context.Background(), query, UUID).Scan(
		&role.UUID,
		&role.Name,
		&role.IsActive,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
	)
	if err != nil {
		return model.ShowRole{}, err
	}
	return role, err
}

func (repo *RoleRepo) Store(request *model.StoreRole) (model.Role, error) {
	query := `INSERT INTO "roles" (uuid, name, is_active, created_at) VALUES($1, $2, $3, $4) 
	RETURNING uuid, name, is_active, created_at`
	var role model.Role
	err := repo.db.QueryRowContext(context.Background(), query, uuid.New(), request.Name, request.IsActive, time.Now()).Scan(
		&role.UUID,
		&role.Name,
		&role.IsActive,
		&role.CreatedAt,
	)
	if err != nil {
		return model.Role{}, err
	}
	return role, err
}

func (repo *RoleRepo) Update(UUID string, request *model.UpdateRole) (model.Role, error) {
	query := `UPDATE "roles" SET name = $2, is_active = $3, updated_at = $4 WHERE uuid = $1 
	RETURNING uuid, name, is_active, created_at, updated_at`
	var role model.Role
	err := repo.db.QueryRowContext(context.Background(), query, UUID, request.Name, request.IsActive, time.Now()).Scan(
		&role.UUID,
		&role.Name,
		&role.IsActive,
		&role.CreatedAt,
		&role.UpdatedAt,
	)
	if err != nil {
		return model.Role{}, err
	}
	return role, err
}

func (repo *RoleRepo) Destroy(UUID string) (model.Role, error) {
	query := `UPDATE "roles" SET updated_at = $2, deleted_at = $3 WHERE uuid = $1 
	RETURNING uuid, name, is_active, created_at, updated_at, deleted_at`
	var role model.Role
	err := repo.db.QueryRowContext(context.Background(), query, UUID, time.Now(), time.Now()).Scan(
		&role.UUID,
		&role.Name,
		&role.IsActive,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
	)
	if err != nil {
		return model.Role{}, err
	}
	return role, err
}

func NewRoleRepo(db *database.DB) RoleRepository {
	return &RoleRepo{db}
}
