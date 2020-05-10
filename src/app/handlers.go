package app

import (
	_ "github.com/go-sql-driver/mysql"
	"root/src/app/database"

	"root/src/delivery/handler"
	"root/src/repository/mysqlds"
	"root/src/usecase"
)



type handlers struct {
	userHandler *handler.UserHandler
}

// Creates route handlers with all dependencies injected
func newHandlers() (*handlers, error) {
	var err error

	db, err := database.NewConnection()
	if err != nil {
		return nil, err
	}

	// repositories
	// user
	dsUserRepository, err := mysqlds.NewDsUserRepository(db)
	if err != nil {
		return nil, err
	}

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