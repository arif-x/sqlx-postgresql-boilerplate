package main

import (
	"github.com/arif-x/sqlx-gofiber-boilerplate/config"
	"github.com/arif-x/sqlx-gofiber-boilerplate/pkg/server"
)

func main() {
	config.LoadAllConfigs(".env")
	server.Serve()
}
