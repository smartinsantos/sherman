package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	_ "sherman/src/app/testing"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("it should return default values if not .env file is found", func(t *testing.T) {
		cd, err := os.Getwd()
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		if err := os.Chdir("./src/app/config"); err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		config := Get()
		assert.Equal(t, &DefaultConfig, config)

		if err := os.Chdir(cd); err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
	})

	t.Run("it should return proper configs based on environments", func(t *testing.T) {
		ce := os.Getenv("ENV")
		testConfig := Get()
		if err := os.Setenv("ENV", "prod"); err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
		config := Get()
		assert.NotEqual(t, testConfig, config)

		if err := os.Setenv("ENV", ce); err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}
	})
}
