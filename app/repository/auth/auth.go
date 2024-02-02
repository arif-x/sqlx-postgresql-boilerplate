package auth

import (
	"context"
	"time"

	model "github.com/arif-x/sqlx-postgresql-boilerplate/app/model/auth"
	"github.com/arif-x/sqlx-postgresql-boilerplate/pkg/database"
	"github.com/google/uuid"
)

type AuthRepository interface {
	Login(Username string) (model.User, string, []string, error)
	Register(*model.Register) (model.User, string, []string, error)
	Verify(username string) (model.User, string, []string, error)
	ForgotPassword(*model.ForgotPassword) (model.User, error)
	ChangeForgotPassword(username string, password string) (model.User, string, []string, error)
}

type AuthRepo struct {
	db *database.DB
}

func (repo *AuthRepo) Login(Username string) (model.User, string, []string, error) {
	var user model.User
	query := `SELECT uuid, name, email, username, password, role_uuid, email_verified_at, is_active, created_at, updated_at, deleted_at FROM users 
	WHERE (username = $1 OR email = $1) AND deleted_at IS NULL LIMIT 1`
	err := repo.db.QueryRowContext(context.Background(), query, Username).Scan(
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.RoleUUID,
		&user.EmailVerifiedAt,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	var permissions []string
	var role_name string

	if err != nil {
		return model.User{}, role_name, permissions, err
	}

	get_role_name_query := `SELECT name FROM roles WHERE uuid = $1`
	role_err := repo.db.QueryRowContext(context.Background(), get_role_name_query, user.RoleUUID).Scan(
		&role_name,
	)
	if role_err != nil {
		return user, role_name, []string{}, err
	}

	get_role_has_permission_query := `SELECT permissions.name as permission FROM "role_has_permissions" 
	JOIN roles ON roles.uuid = role_has_permissions.role_uuid 
	JOIN permissions ON permissions.uuid = role_has_permissions.permission_uuid 
	WHERE roles.uuid = $1`

	permissionsRows, err := repo.db.QueryContext(context.Background(), get_role_has_permission_query, user.RoleUUID)
	if err != nil {
		return user, role_name, []string{}, err
	}
	defer permissionsRows.Close()

	for permissionsRows.Next() {
		var permission string
		err := permissionsRows.Scan(&permission)
		if err != nil {
			return user, role_name, []string{}, err
		}
		permissions = append(permissions, permission)
	}

	if err := permissionsRows.Err(); err != nil {
		return user, role_name, []string{}, err
	}

	return user, role_name, permissions, err
}

func (repo *AuthRepo) Register(request *model.Register) (model.User, string, []string, error) {
	inactive_q := "SELECT uuid FROM roles WHERE name = 'Inactive' LIMIT 1"
	var inactive_role_uuid uuid.UUID
	_ = repo.db.QueryRow(inactive_q).Scan(&inactive_role_uuid)

	query := `INSERT INTO "users" (uuid, name, username, email, password, role_uuid, created_at) VALUES($1, $2, $3, $4, $5, $6, $7) 
	RETURNING uuid, name, email, username, password, role_uuid, email_verified_at, is_active, created_at, updated_at, deleted_at`
	var user model.User
	err := repo.db.QueryRowContext(context.Background(), query, uuid.New(), request.Name, request.Username, request.Email, request.Password, inactive_role_uuid, time.Now()).Scan(
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.RoleUUID,
		&user.EmailVerifiedAt,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	var permissions []string
	var role_name string

	if err != nil {
		return model.User{}, role_name, permissions, nil
	}

	get_role_name_query := `SELECT name FROM roles WHERE uuid = $1`
	role_err := repo.db.QueryRowContext(context.Background(), get_role_name_query, user.RoleUUID).Scan(
		&role_name,
	)
	if role_err != nil {
		return user, role_name, []string{}, err
	}

	get_role_has_permission_query := `SELECT permissions.name as permission FROM "role_has_permissions" 
	JOIN roles ON roles.uuid = role_has_permissions.role_uuid 
	JOIN permissions ON permissions.uuid = role_has_permissions.permission_uuid 
	WHERE roles.uuid = $1`

	permissionsRows, err := repo.db.QueryContext(context.Background(), get_role_has_permission_query, user.RoleUUID)
	if err != nil {
		return user, role_name, []string{}, err
	}
	defer permissionsRows.Close()

	for permissionsRows.Next() {
		var permission string
		err := permissionsRows.Scan(&permission)
		if err != nil {
			return user, role_name, []string{}, err
		}
		permissions = append(permissions, permission)
	}

	if err := permissionsRows.Err(); err != nil {
		return user, role_name, []string{}, err
	}

	return user, role_name, permissions, err
}

func (repo *AuthRepo) Verify(username string) (model.User, string, []string, error) {
	var new_role_uuid string
	get_verified_role_uuid_query := `SELECT uuid FROM roles WHERE lower(name) = 'verified'`
	verified_role_err := repo.db.QueryRowContext(context.Background(), get_verified_role_uuid_query).Scan(
		&new_role_uuid,
	)
	if verified_role_err != nil {
		return model.User{}, "", []string{}, verified_role_err
	}

	query := `UPDATE "users" SET email_verified_at = $2, is_active = $3, updated_at = $4, role_uuid = $5 WHERE username = $1 
	RETURNING uuid, name, email, username, password, role_uuid, email_verified_at, is_active, created_at, updated_at, deleted_at`
	var user model.User
	err := repo.db.QueryRowContext(context.Background(), query, username, time.Now(), true, time.Now(), new_role_uuid).Scan(
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.RoleUUID,
		&user.EmailVerifiedAt,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	var permissions []string
	var role_name string

	if err != nil {
		return model.User{}, role_name, permissions, err
	}

	get_role_name_query := `SELECT name FROM roles WHERE uuid = $1`
	role_err := repo.db.QueryRowContext(context.Background(), get_role_name_query, user.RoleUUID).Scan(
		&role_name,
	)
	if role_err != nil {
		return user, role_name, []string{}, err
	}

	get_role_has_permission_query := `SELECT permissions.name as permission FROM "role_has_permissions" 
	JOIN roles ON roles.uuid = role_has_permissions.role_uuid 
	JOIN permissions ON permissions.uuid = role_has_permissions.permission_uuid 
	WHERE roles.uuid = $1`

	permissionsRows, err := repo.db.QueryContext(context.Background(), get_role_has_permission_query, user.RoleUUID)
	if err != nil {
		return user, role_name, []string{}, err
	}
	defer permissionsRows.Close()

	for permissionsRows.Next() {
		var permission string
		err := permissionsRows.Scan(&permission)
		if err != nil {
			return user, role_name, []string{}, err
		}
		permissions = append(permissions, permission)
	}

	if err := permissionsRows.Err(); err != nil {
		return user, role_name, []string{}, err
	}

	return user, role_name, permissions, err
}

func (repo *AuthRepo) ForgotPassword(request *model.ForgotPassword) (model.User, error) {
	var user model.User
	query := `SELECT uuid, name, email, username, password, role_uuid, email_verified_at, is_active, created_at, updated_at, deleted_at
	FROM users WHERE username = $1 OR email = $1 AND deleted_at = NULL LIMIT 1`
	err := repo.db.QueryRowContext(context.Background(), query, request.Username).Scan(
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.RoleUUID,
		&user.EmailVerifiedAt,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *AuthRepo) ChangeForgotPassword(username string, password string) (model.User, string, []string, error) {
	query := `UPDATE "users" SET password = $2, updated_at = $3 WHERE username = $1 
	RETURNING uuid, name, email, username, password, role_uuid, email_verified_at, is_active, created_at, updated_at, deleted_at`
	var user model.User
	err := repo.db.QueryRowContext(context.Background(), query, username, password, time.Now()).Scan(
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.RoleUUID,
		&user.EmailVerifiedAt,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	var permissions []string
	var role_name string

	if err != nil {
		return model.User{}, role_name, permissions, err
	}

	get_role_name_query := `SELECT name FROM roles WHERE uuid = $1`
	role_err := repo.db.QueryRowContext(context.Background(), get_role_name_query, user.RoleUUID).Scan(
		&role_name,
	)
	if role_err != nil {
		return user, role_name, []string{}, err
	}

	get_role_has_permission_query := `SELECT permissions.name as permission FROM "role_has_permissions" 
	JOIN roles ON roles.uuid = role_has_permissions.role_uuid 
	JOIN permissions ON permissions.uuid = role_has_permissions.permission_uuid 
	WHERE roles.uuid = $1`

	permissionsRows, err := repo.db.QueryContext(context.Background(), get_role_has_permission_query, user.RoleUUID)
	if err != nil {
		return user, role_name, []string{}, err
	}
	defer permissionsRows.Close()

	for permissionsRows.Next() {
		var permission string
		err := permissionsRows.Scan(&permission)
		if err != nil {
			return user, role_name, []string{}, err
		}
		permissions = append(permissions, permission)
	}

	if err := permissionsRows.Err(); err != nil {
		return user, role_name, []string{}, err
	}

	return user, role_name, permissions, err
}

func NewAuthRepo(db *database.DB) AuthRepository {
	return &AuthRepo{db}
}
