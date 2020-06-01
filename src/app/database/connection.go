package database

import (
	"database/sql"
	"errors"
	"fmt"
	"sherman/src/app/config"
)

// NewConnection creates a db connection
func NewConnection() (*sql.DB, error) {
	cfg := config.Get()

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
	default:
		errorMessage := fmt.Sprintf("DB_DRIVER: %s, not supported", cfg.DB.Driver)
		return nil, errors.New(errorMessage)
	}

	db, err := sql.Open(cfg.DB.Driver, connectionURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
