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

type UserRepository interface {
	Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.User, int, error)
	Show(ID string) (model.UserShow, error)
	Store(model *model.StoreUser) (model.User, error)
	Update(ID string, request *model.UpdateUser) (model.User, error)
	Destroy(ID string) (model.User, error)
}

type UserRepo struct {
	db *database.DB
}

func (repo *UserRepo) Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.User, int, error) {
	_select := "id, name, email, username, created_at, updated_at, deleted_at"
	_conditions := database.Search([]string{"name", "email", "username"}, search)
	_order := database.OrderBy(sort_by, sort)
	_limit := database.Limit(limit, offset)

	count_query := fmt.Sprintf(`SELECT count(*) FROM users %s`, _conditions)
	var count int
	_ = repo.db.QueryRow(count_query).Scan(&count)

	query := fmt.Sprintf(`SELECT %s FROM users %s %s %s`, _select, _conditions, _order, _limit)

	rows, err := repo.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()
	var items []model.User
	for rows.Next() {
		var i model.User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Username,
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

func (repo *UserRepo) Show(ID string) (model.UserShow, error) {
	var user model.UserShow
	query := "SELECT id, name, email, username, created_at, updated_at, deleted_at FROM users WHERE id = $1 LIMIT 1"
	err := repo.db.QueryRowContext(context.Background(), query, ID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return model.UserShow{}, err
	}
	return user, err
}

func (repo *UserRepo) Store(request *model.StoreUser) (model.User, error) {
	query := `INSERT INTO "users" (id, name, username, email, password, created_at) VALUES($1, $2, $3, $4, $5, $6) 
	RETURNING id, name, username, email, created_at`
	var user model.User
	err := repo.db.QueryRowContext(context.Background(), query, uuid.New(), request.Name, request.Username, request.Email, request.Password, time.Now()).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (repo *UserRepo) Update(ID string, request *model.UpdateUser) (model.User, error) {
	if request.Password == "" {
		query := `UPDATE "users" SET name = $2, username = $3, email = $4, updated_at = $5 WHERE id = $1 
		RETURNING id, name, username, email, created_at, updated_at`
		var user model.User
		err := repo.db.QueryRowContext(context.Background(), query, ID, request.Name, request.Username, request.Email, time.Now()).Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return model.User{}, err
		}
		return user, err
	} else {
		query := `UPDATE "users" SET name = $2, username = $3, email = $4, password = $5, updated_at = $6 WHERE id = $1 
		RETURNING id, name, username, email, created_at, updated_at`
		var user model.User
		err := repo.db.QueryRowContext(context.Background(), query, ID, request.Name, request.Username, request.Email, request.Password, time.Now()).Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return model.User{}, err
		}
		return user, err
	}
}

func (repo *UserRepo) Destroy(ID string) (model.User, error) {
	query := `UPDATE "users" SET updated_at = $2, deleted_at = $3 WHERE id = $1 
	RETURNING id, name, username, email, created_at, updated_at, deleted_at`
	var user model.User
	err := repo.db.QueryRowContext(context.Background(), query, ID, time.Now(), time.Now()).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func NewUserRepo(db *database.DB) UserRepository {
	return &UserRepo{db}
}
