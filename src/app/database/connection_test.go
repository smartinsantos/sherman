package database

import (
	"github.com/stretchr/testify/assert"
	"sherman/src/app/config"
	_ "sherman/src/app/testing"
	"testing"
)

func TestNewConnection(t *testing.T) {
	cfg := config.Get()
	t.Run("it should succeed", func(t *testing.T) {
		_, err := NewConnection(&ConnectionConfig{
			Driver: cfg.DB.Driver,
			User:   cfg.DB.User,
			Pass:   cfg.DB.Pass,
			Host:   "localhost",
			Port:   cfg.DB.ExposedPort,
			Name:   cfg.DB.Name,
		})

		assert.NoError(t, err)
	})

	t.Run("it should return an error", func(t *testing.T) {
		_, err := NewConnection(&ConnectionConfig{
			Driver: "some_unsupported_driver",
			User:   cfg.DB.User,
			Pass:   cfg.DB.Pass,
			Host:   "localhost",
			Port:   cfg.DB.ExposedPort,
			Name:   cfg.DB.Name,
		})

		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "DB_DRIVER: some_unsupported_driver, not supported")
		}
	})

	t.Run("it should return an error", func(t *testing.T) {
		_, err := NewConnection(&ConnectionConfig{})

		assert.Error(t, err)
	})
}
