package app

import (
	"database/sql"
	"github.com/sarulabs/di"
	"root/src/app/database"
	"root/src/delivery/handler"
	"root/src/repository/mysqlds"
	"root/src/usecase"
)

// DIContainer - definitions of the application services.
var DIContainer = []di.Def {
	{
		Name:  "mysql-db",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return database.NewConnection()
		},
		Close: func(obj interface{}) error {
			return obj.(*sql.DB).Close()
		},
	},
	{
		Name:  "mysql-user-repository",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &mysqlds.UserRepository{
				DB: ctn.Get("mysql-db").(*sql.DB),
			}, nil
		},
	},
	{
		Name:  "user-usecase",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &usecase.UserUseCase{
				Repo:   ctn.Get("mysql-user-repository").(*mysqlds.UserRepository),
			}, nil
		},
	},
	{
		Name:  "user-handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &handler.UserHandler{
				UserUseCase: ctn.Get("user-usecase").(*usecase.UserUseCase),
			}, nil
		},
	},
}