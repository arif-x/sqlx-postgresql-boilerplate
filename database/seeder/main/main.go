package main

import (
	"fmt"
	"log"

	"github.com/arif-x/sqlx-gofiber-boilerplate/config"
	seeds "github.com/arif-x/sqlx-gofiber-boilerplate/database/seeder"
	seeders "github.com/danvergara/seeder"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	// postgres driver.
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("can't load .env file. error: %v", err)
	}
	config.LoadDBCfg()
	url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", config.DBCfg().User, config.DBCfg().Password, config.DBCfg().Host, config.DBCfg().Port, config.DBCfg().Name)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		log.Fatalf("error opening a connection with the database %s\n", err)
	}

	s := seeds.NewSeed(db)

	if err := seeders.Execute(s); err != nil {
		log.Fatalf("error seeding the db %s\n", err)
	}
}
