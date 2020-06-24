package database

import (
	"database/sql"
	"errors"
	"fmt"
	"sherman/src/service/config"
)

// NewConnection creates a db connection
func NewConnection(dbCfg config.DBConfig) (*sql.DB, error) {
	var connectionURL string

	switch dbCfg.Driver {
	case "mysql":
		connectionURL = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbCfg.User,
			dbCfg.Pass,
			dbCfg.Host,
			dbCfg.Port,
			dbCfg.Name,
		)
	default:
		errorMessage := fmt.Sprintf("DB_DRIVER: %s, not supported", dbCfg.Driver)
		return nil, errors.New(errorMessage)
	}

	db, err := sql.Open(dbCfg.Driver, connectionURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
