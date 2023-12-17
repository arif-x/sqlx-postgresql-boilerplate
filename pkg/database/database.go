package database

import (
	"fmt"

	"github.com/arif-x/sqlx-gofiber-boilerplate/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct{ *sqlx.DB }

var defaultDB = &DB{}

func (db *DB) connect(cfg *config.DB) (err error) {
	dbURI := fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s",
		cfg.Host,
		cfg.Port,
		cfg.SslMode,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db.DB, err = sqlx.Open("postgres", dbURI)
	if err != nil {
		return err
	}

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return fmt.Errorf("can't sent ping to database, %w", err)
	}

	return nil
}

// GetDB returns db instance
func GetDB() *DB {
	return defaultDB
}

// ConnectDB sets the db client of database using default configuration
func ConnectDB() error {
	return defaultDB.connect(config.DBCfg())
}
