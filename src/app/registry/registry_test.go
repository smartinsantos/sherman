package registry

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"sherman/src/app/config"
	_ "sherman/src/app/testing"
	"sherman/src/delivery/handler"
	"sherman/src/domain/auth"
	"sherman/src/service/middleware"
	"sherman/src/service/presenter"
	"sherman/src/service/security"
	"sherman/src/service/validator"
	"testing"
)

func TestGetAppContainer(t *testing.T) {
	cfg := config.DefaultConfig
	cfg.DB.Host = "localhost"
	cfg.DB.Port = "5001"

	t.Run("it should return the same instance every time is called", func(t *testing.T) {
		diContainer, err := GetAppContainer(cfg)
		if assert.NoError(t, err) {
			diContainer2, _ := GetAppContainer(cfg)
			assert.Equal(t, diContainer, diContainer2)
		}
	})

	t.Run("it should have all expected definitions", func(t *testing.T) {
		diContainer, err := GetAppContainer(cfg)

		if assert.NoError(t, err) {
			_, ok := diContainer.Get("mysql-db").(*sql.DB)
			assert.True(t, ok)
			_, ok = diContainer.Get("middleware-service").(middleware.Middleware)
			assert.True(t, ok)
			_, ok = diContainer.Get("presenter-service").(presenter.Presenter)
			assert.True(t, ok)
			_, ok = diContainer.Get("security-service").(security.Security)
			assert.True(t, ok)
			_, ok = diContainer.Get("validator-service").(validator.Validator)
			assert.True(t, ok)
			_, ok = diContainer.Get("mysql-security-token-repository").(auth.SecurityTokenRepository)
			assert.True(t, ok)
			_, ok = diContainer.Get("mysql-user-repository").(auth.UserRepository)
			assert.True(t, ok)
			_, ok = diContainer.Get("security-token-usecase").(auth.SecurityTokenUseCase)
			assert.True(t, ok)
			_, ok = diContainer.Get("user-usecase").(auth.UserUseCase)
			assert.True(t, ok)
			_, ok = diContainer.Get("user-handler").(handler.UserHandler)
			assert.True(t, ok)
		}
	})
}
