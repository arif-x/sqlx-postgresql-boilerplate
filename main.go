package main

import (
	"github.com/arif-x/sqlx-postgresql-boilerplate/cmd"
	_ "github.com/arif-x/sqlx-postgresql-boilerplate/docs"
)

// Swagger Config
// @title SQLX PostgreSQL Boilerplate API
// @version 1.0
// @description SQLX PostgreSQL Boilerplate API Swag.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
