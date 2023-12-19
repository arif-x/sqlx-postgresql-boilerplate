package main

import (
	"log"

	seeds "github.com/arif-x/sqlx-gofiber-boilerplate/database/seeder"
	seeders "github.com/danvergara/seeder"
	"github.com/jmoiron/sqlx"

	// postgres driver.
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Open("postgres", "postgresql://postgres:password@localhost:5432/go_boiler?sslmode=disable")
	if err != nil {
		log.Fatalf("error opening a connection with the database %s\n", err)
	}

	s := seeds.NewSeed(db)

	if err := seeders.Execute(s); err != nil {
		log.Fatalf("error seeding the db %s\n", err)
	}
}
