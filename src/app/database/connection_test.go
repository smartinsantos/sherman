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
		_, err := NewConnection(cfg.DB)

		assert.NoError(t, err)
	})

	t.Run("it should return an error", func(t *testing.T) {
		dbConfig := config.TestConfig.DB
		dbConfig.Driver = "some_unsupported_driver"
		_, err := NewConnection(dbConfig)

		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "DB_DRIVER: some_unsupported_driver, not supported")
		}
	})

	t.Run("it should return an error", func(t *testing.T) {
		_, err := NewConnection(config.DBConfig{})

		assert.Error(t, err)
	})
}
