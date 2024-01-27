package main

import (
	"github.com/arif-x/sqlx-gofiber-boilerplate/config"
	_ "github.com/arif-x/sqlx-gofiber-boilerplate/docs"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/server"
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
	config.LoadAllConfigs(".env")
	server.Serve()
}
