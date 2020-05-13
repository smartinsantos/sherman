package registry

import (
	"database/sql"
	"github.com/sarulabs/di"
	"root/src/app/database"
	"root/src/delivery/handler"
	"root/src/domain"
	"root/src/repository/mysqlds"
	"root/src/usecase"
)

// Registry - definitions of the application services.
var Registry = []di.Def {
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
			var userRepository domain.UserRepository
			userRepository = &mysqlds.UserRepository {
				DB: ctn.Get("mysql-db").(*sql.DB),
			}
			return userRepository, nil
		},
	},
	{
		Name:  "user-usecase",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			var userUseCase domain.UserUseCase
			userUseCase = &usecase.UserUseCase {
				UserRepo: ctn.Get("mysql-user-repository").(*mysqlds.UserRepository),
			}
			return userUseCase, nil
		},
	},
	{
		Name:  "user-handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &handler.UserHandler {
				UserUseCase: ctn.Get("user-usecase").(*usecase.UserUseCase),
			}, nil
		},
	},
}
