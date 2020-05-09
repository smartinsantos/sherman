package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"root/config"
	"root/delivery/handler"
	"root/repository/datastore"
	"root/usecase"
)

type handlers struct {
	userHandler *handler.UserHandler
}

// Creates route handlers with all dependencies injected
func NewHandlers() (*handlers, error) {
	db, err := newDbConnection()
	if err != nil {
		return nil, err
	}

	// repositories
	dsUserRepository := datastore.NewDsUserRepository(db)

	// use cases
	userUseCase := usecase.NewUserUseCase(dsUserRepository)

	// handlers
	userHandler := handler.NewUserHandler(userUseCase)

	// app handler
	appHandler := handlers{
		userHandler: userHandler,
	}
	return &appHandler, nil
}

func newDbConnection() (*gorm.DB, error) {
	cfg := config.Get()
	connectionUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Db.User,
		cfg.Db.Pass,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Name,
	)
	db, err := gorm.Open(cfg.Db.Driver, connectionUrl)
	if err != nil {
		return nil, err
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}