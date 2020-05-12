package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"root/config"
)

// Creates a db connection
func NewConnection() (*sql.DB, error) {
	cfg := config.Get()

	var connectionUrl string

	switch cfg.Db.Driver {
		case "mysql": {
			connectionUrl = fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				cfg.Db.User,
				cfg.Db.Pass,
				cfg.Db.Host,
				cfg.Db.Port,
				cfg.Db.Name,
			)
		}
		default: {
			errorMessage := fmt.Sprintf("DB_DRIVER: %s, not supported", cfg.Db.Driver)
			return nil, errors.New(errorMessage)
		}
	}

	db, err := sql.Open(cfg.Db.Driver, connectionUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}