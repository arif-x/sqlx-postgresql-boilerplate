package main

import (
	"github.com/arif-x/sqlx-gofiber-boilerplate/cmd"
	_ "github.com/arif-x/sqlx-gofiber-boilerplate/docs"
)

// Swagger Config
// @title SQLX GoFiber Boilerplate API
// @version 1.0
// @description SQLX GoFiber Boilerplate API Swag.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
