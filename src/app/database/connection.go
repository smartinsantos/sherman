package database

import (
	"database/sql"
	"errors"
	"fmt"
	// mysql driver import
	_ "github.com/go-sql-driver/mysql"
	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"sherman/src/app/config"
)

// NewConnection creates a db connection
func NewConnection(cfg *config.GlobalConfig) (*sql.DB, error) {
	var connectionURL string

	switch cfg.DB.Driver {
	case "mysql":
		connectionURL = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DB.User,
			cfg.DB.Pass,
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.Name,
		)
	case "sqlite3":
		connectionURL = cfg.DB.Path
	default:
		errorMessage := fmt.Sprintf("DB_DRIVER: %s, not supported", cfg.DB.Driver)
		return nil, errors.New(errorMessage)
	}

	db, _ := sql.Open(cfg.DB.Driver, connectionURL)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
