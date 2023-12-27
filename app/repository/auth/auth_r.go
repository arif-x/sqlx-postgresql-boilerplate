package auth

import (
	"context"
	"time"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/auth"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
	"github.com/google/uuid"
)

type AuthRepository interface {
	Login(Username string) (model.User, error)
	Register(*model.Register) (model.User, error)
}

type AuthRepo struct {
	db *database.DB
}

func (repo *AuthRepo) Login(Username string) (model.User, error) {
	var user model.User
	query := `SELECT uuid, name, email, username, password, created_at, updated_at, deleted_at FROM users 
	WHERE username = $1 OR email = $1 AND deleted_at IS NULL LIMIT 1`
	err := repo.db.QueryRowContext(context.Background(), query, Username).Scan(
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (repo *AuthRepo) Register(request *model.Register) (model.User, error) {
	query := `INSERT INTO "users" (uuid, name, username, email, password, created_at) VALUES($1, $2, $3, $4, $5, $6) 
	RETURNING uuid, name, email, username, password, created_at, updated_at, deleted_at`
	var user model.User
	err := repo.db.QueryRowContext(context.Background(), query, uuid.New(), request.Name, request.Username, request.Email, request.Password, time.Now()).Scan(
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func NewAuthRepo(db *database.DB) AuthRepository {
	return &AuthRepo{db}
}
