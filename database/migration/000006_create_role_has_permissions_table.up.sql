CREATE TABLE IF NOT EXISTS role_has_permissions (
    role_uuid UUID NOT NULL,
    permission_uuid UUID NOT NULL
);