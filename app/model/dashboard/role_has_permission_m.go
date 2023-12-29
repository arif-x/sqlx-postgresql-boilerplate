package dashboard

import (
	jsonutil "github.com/arif-x/sqlx-gofiber-boilerplate/pkg/json"
	"github.com/google/uuid"
)

type ShowRoleHasPermission struct {
	UUID       uuid.UUID         `db:"uuid" json:"uuid"`
	Name       string            `db:"name" json:"name"`
	Permission *jsonutil.JSONRaw `db:"permission" json:"permission"`
}

type StoreRoleHasPermission struct {
	RoleUUID       string   `db:"role_uuid" json:"role_uuid"`
	PermissionUUID []string `db:"permission_uuid" json:"permission_uuid"`
}
