package registry

import (
	"database/sql"
	"errors"
	"github.com/sarulabs/di"
	"root/src/app/database"
	"root/src/delivery/handler"
	"root/src/domain/auth"
	"root/src/repository/mysqlds"
	"root/src/usecase"
	"sync"
)

var (
	// definitions of the application services.
	registry = []di.Def {
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
				var userRepository auth.UserRepository = &mysqlds.UserRepository {
					DB: ctn.Get("mysql-db").(*sql.DB),
				}
				return userRepository, nil
			},
		},
		{
			Name:  "mysql-security-token-repository",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				var securityTokenRepository auth.SecurityTokenRepository = &mysqlds.SecurityTokenRepository {
					DB: ctn.Get("mysql-db").(*sql.DB),
				}
				return securityTokenRepository, nil
			},
		},
		{
			Name:  "security-token-usecase",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				var securityTokenUseCase auth.SecurityTokenUseCase = &usecase.SecurityTokenUseCase {
					SecurityTokenRepo: ctn.Get("mysql-security-token-repository").(*mysqlds.SecurityTokenRepository),
				}
				return securityTokenUseCase, nil
			},
		},
		{
			Name:  "user-usecase",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				var userUseCase auth.UserUseCase = &usecase.UserUseCase {
					UserRepo: ctn.Get("mysql-user-repository").(*mysqlds.UserRepository),
					SecurityTokenUseCase: ctn.Get("security-token-usecase").(*usecase.SecurityTokenUseCase),
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
					SecurityTokenUseCase: ctn.Get("security-token-usecase").(*usecase.SecurityTokenUseCase),
				}, nil
			},
		},
	}
	container di.Container
	once sync.Once
)

// GetAppContainer retrieves an instance of app container with dependency injected services
func GetAppContainer() (di.Container, error) {
	once.Do(func() {
		builder, err := di.NewBuilder()
		if err != nil {
			container = nil
			return
		}
		err = builder.Add(registry...)
		if err != nil {
			container = nil
			return
		}
		container = builder.Build()
	})

	if container == nil {
		return nil, errors.New("could not create container")
	}

	return container, nil
}
