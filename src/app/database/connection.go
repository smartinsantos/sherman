package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// ConnectionConfig is the definition of the database cfg param
type ConnectionConfig struct {
	Driver string
	User   string
	Pass   string
	Host   string
	Port   string
	Name   string
}

// NewConnection creates a db connection
func NewConnection(cfg *ConnectionConfig) (*sql.DB, error) {
	var connectionURL string

	switch cfg.Driver {
	case "mysql":
		connectionURL = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.User,
			cfg.Pass,
			cfg.Host,
			cfg.Port,
			cfg.Name,
		)
	default:
		errorMessage := fmt.Sprintf("DB_DRIVER: %s, not supported", cfg.Driver)
		return nil, errors.New(errorMessage)
	}

	db, err := sql.Open(cfg.Driver, connectionURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
