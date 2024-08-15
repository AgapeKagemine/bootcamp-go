package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() (db *sql.DB, err error) {
	config := &DBConfig{
		Driver:   "pgx",
		Username: "training",
		Password: 1234,
		Host:     "localhost",
		Port:     5432,
		Database: "orchestrator-config",
	}

	connString := fmt.Sprintf("postgres://%s:%d@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err = sql.Open(config.Driver, connString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
