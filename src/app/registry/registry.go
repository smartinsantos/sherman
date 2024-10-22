package registry

import (
	"database/sql"
	"github.com/rs/zerolog/log"
	"github.com/sarulabs/di"
	"sherman/src/app/config"
	"sherman/src/app/database"
	"sherman/src/delivery/handler"
	"sherman/src/domain/auth"
	"sherman/src/repository/mysqlds"
	"sherman/src/service/middleware"
	"sherman/src/service/presenter"
	"sherman/src/service/security"
	"sherman/src/service/validator"
	"sherman/src/usecase"
	"sync"
)

var (
	container di.Container
	once      sync.Once
)

func makeRegistry(cfg *config.GlobalConfig) []di.Def {
	return []di.Def{
		{
			Name:  "mysql-db",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				db, err := database.NewConnection(cfg)
				if err != nil {
					log.Error().Msg(err.Error())
				}
				return db, err
			},
			Close: func(db interface{}) error {
				return db.(*sql.DB).Close()
			},
		},
		{
			Name:  "middleware-service",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				securityService := ctn.Get("security-service").(security.Security)
				return middleware.New(cfg, securityService), nil
			},
		},
		{
			Name:  "presenter-service",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return presenter.New(), nil
			},
		},
		{
			Name:  "security-service",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return security.New(cfg), nil
			},
		},
		{
			Name:  "validator-service",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return validator.New(), nil
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
			Name:  "mysql-user-repository",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				db := ctn.Get("mysql-db").(*sql.DB)
				return mysqlds.NewUserRepository(db), nil
			},
		},
		{
			Name:  "security-token-usecase",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				securityTokenRepo := ctn.Get("mysql-security-token-repository").(auth.SecurityTokenRepository)
				securityService := ctn.Get("security-service").(security.Security)
				return usecase.NewSecurityTokenUseCase(securityTokenRepo, securityService), nil
			},
		},
		{
			Name:  "user-usecase",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				userRepo := ctn.Get("mysql-user-repository").(auth.UserRepository)
				securityService := ctn.Get("security-service").(security.Security)
				return usecase.NewUserUseCase(userRepo, securityService), nil
			},
		},
		{
			Name:  "user-handler",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				userUseCase := ctn.Get("user-usecase").(auth.UserUseCase)
				securityTokenUseCase := ctn.Get("security-token-usecase").(auth.SecurityTokenUseCase)
				validatorService := ctn.Get("validator-service").(validator.Validator)
				securityService := ctn.Get("security-service").(security.Security)
				presenterService := ctn.Get("presenter-service").(presenter.Presenter)
				return handler.NewUserHandler(
					userUseCase,
					securityTokenUseCase,
					validatorService,
					securityService,
					presenterService,
				), nil
			},
		},
	}
}

// Get retrieves an instance of app container with dependency injected service
func Get() (di.Container, error) {
	once.Do(func() {
		builder, err := di.NewBuilder()
		if err != nil {
			container = nil
			return
		}

		if err := builder.Add(makeRegistry(config.Get())...); err != nil {
			container = nil
			return
		}
		container = builder.Build()
	})

	return container, nil
}
