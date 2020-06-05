package registry

import (
	"database/sql"
	"errors"
	"github.com/sarulabs/di"
	"sherman/src/app/database"
	"sherman/src/delivery/handler"
	"sherman/src/domain/auth"
	"sherman/src/repository/mysqlds"
	"sherman/src/usecase"
	"sync"
)

var (
	// definitions of the application services.
	registry = []di.Def{
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
				db := ctn.Get("mysql-db").(*sql.DB)
				return mysqlds.NewUserRepository(db), nil
			},
		},
		{
			Name:  "mysql-security-token-repository",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				db := ctn.Get("mysql-db").(*sql.DB)
				return mysqlds.NewSecurityTokenRepository(db), nil
			},
		},
		{
			Name:  "security-token-usecase",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				securityTokenRepo := ctn.Get("mysql-security-token-repository").(auth.SecurityTokenRepository)
				return usecase.NewSecurityTokenUseCase(securityTokenRepo), nil
			},
		},
		{
			Name:  "user-usecase",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				userRepo := ctn.Get("mysql-user-repository").(auth.UserRepository)
				return usecase.NewUserUseCase(userRepo), nil
			},
		},
		{
			Name:  "user-handler",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				userUseCase := ctn.Get("user-usecase").(auth.UserUseCase)
				securityTokenUseCase := ctn.Get("security-token-usecase").(auth.SecurityTokenUseCase)
				return handler.NewUserHandler(userUseCase, securityTokenUseCase), nil
			},
		},
	}
	container di.Container
	once      sync.Once
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
