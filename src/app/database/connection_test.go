package database

import (
	"github.com/stretchr/testify/assert"
	"sherman/src/app/config"
	_ "sherman/src/app/testing"
	"testing"
)

func TestNewConnection(t *testing.T) {
	t.Run("it should succeed", func(t *testing.T) {
		_, err := NewConnection(&config.TestConfig)

		assert.NoError(t, err)
	})

	t.Run("it should return an error", func(t *testing.T) {
		cfg := config.DefaultConfig
		cfg.DB.Driver = "some_unsupported_driver"
		_, err := NewConnection(&cfg)

		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "DB_DRIVER: some_unsupported_driver, not supported")
		}
	})

	t.Run("it should return an error", func(t *testing.T) {
		_, err := NewConnection(&config.GlobalConfig{})

		assert.Error(t, err)
	})
}
