package dashboard

import (
	"context"
	"fmt"
	"log"
	"time"

	model "github.com/arif-x/sqlx-gofiber-boilerplate/app/model/dashboard"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/database"
)

type UserRepository interface {
	Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.User, int, error)
	Show(ID int) (*model.User, error)
	Store(b *model.StoreUser) error
	Update(ID int, user *model.UpdateUser) error
	Delete(ID int) error
}

type UserRepo struct {
	db *database.DB
}

func (repo *UserRepo) Index(limit int, offset uint, search string, sort_by string, sort string) ([]model.User, int, error) {
	_select := "id, name, email, username, created_at, updated_at"
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

func (repo *UserRepo) Show(ID int) (*model.User, error) {
	query := "SELECT id, name, email, username, created_at, updated_at FROM users WHERE id = $1 LIMIT 1"
	row := repo.db.QueryRowContext(context.Background(), query, ID)
	var user *model.User
	err := row.Scan(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) Store(request *model.StoreUser) error {
	query := `INSERT INTO "users" (id, name, username, email, password, created_at) VALUES($1, $2, $3, $4, $5, $6)`
	_, err := repo.db.Exec(query, request.ID, request.Name, request.Email, request.Password, time.Now())
	return err
}

func (repo *UserRepo) Update(ID int, user *model.UpdateUser) error {
	query := `UPDATE "users" SET name = $2, username = $3, email = $4, password = $5, updated_at = $6 WHERE id = $1`
	_, err := repo.db.Exec(query, ID, user.Name, user.Username, user.Email, user.Password, time.Now())
	return err
}

func (repo *UserRepo) Delete(ID int) error {
	query := `UPDATE "users" SET updated_at = $2, deleted_at = $3 WHERE id = $1`
	_, err := repo.db.Exec(query, ID, time.Now(), time.Now())
	return err
}

func NewUserRepo(db *database.DB) UserRepository {
	return &UserRepo{db}
}
