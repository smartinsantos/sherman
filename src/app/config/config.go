package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"strconv"
	"sync"
)

type (
	appConfig struct {
		Env   string
		Debug bool
		Addr  string
	}
	dbConfig struct {
		Driver string
		Name   string
		User   string
		Pass   string
		Host   string
		Port   string
	}
	jwtConfig struct {
		Secret string
	}

	// Config Wrapper for all configurations
	Config struct {
		App appConfig
		DB  dbConfig
		Jwt jwtConfig
	}
)

var (
	defaultConfig = &Config{
		App: appConfig{
			Env:   "local",
			Debug: true,
			Addr:  ":8080",
		},
		DB: dbConfig{
			Driver: "mysql",
			Name:   "db_name",
			User:   "db_user",
			Pass:   "db_password",
			Host:   "app-mysql",
			Port:   "3306",
		},
		Jwt: jwtConfig{
			Secret: "jwt_secret",
		},
	}
	config *Config
	once   sync.Once
)

// Get returns an instance of Config
func Get() *Config {
	once.Do(func() {
		envMap, err := godotenv.Read(".env")

		if err != nil {
			log.Error().Msg("config error: couldn't read contents of .env file, using defaults")
		}

		config = &Config{
			App: appConfig{
				Env:   getKey(envMap, "APP_ENV", defaultConfig.App.Env),
				Debug: getKeyAsBool(envMap, "APP_DEBUG", defaultConfig.App.Debug),
				Addr:  getKey(envMap, "APP_ADDR", defaultConfig.App.Addr),
			},
			DB: dbConfig{
				Driver: getKey(envMap, "DB_DRIVER", defaultConfig.DB.Driver),
				Name:   getKey(envMap, "DB_NAME", defaultConfig.DB.Name),
				User:   getKey(envMap, "DB_USER", defaultConfig.DB.User),
				Pass:   getKey(envMap, "DB_PASS", defaultConfig.DB.Pass),
				Host:   getKey(envMap, "DB_HOST", defaultConfig.DB.Host),
				Port:   getKey(envMap, "DB_PORT", defaultConfig.DB.Port),
			},
			Jwt: jwtConfig{Secret: getKey(envMap, "JWT_SECRET", defaultConfig.Jwt.Secret)},
		}
	})
	return config
}

func getKey(env map[string]string, key, defaultValue string) string {
	if value, exist := env[key]; exist {
		return value
	}
	return defaultValue
}

func getKeyAsBool(env map[string]string, key string, defaultValue bool) bool {
	valueStr := getKey(env, key, "")
	if val, err := strconv.ParseBool(valueStr); err == nil {
		return val
	}
	return defaultValue
}

func getKeyAsInt(env map[string]string, key string, defaultValue int) int {
	valueStr := getKey(env, key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
