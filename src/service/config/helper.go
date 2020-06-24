package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"strconv"
)

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

func loadFromEnv() GlobalConfig {
	envMap, err := godotenv.Read(".env")

	if err != nil {
		log.Error().Msg("config error: couldn't read contents of .env file, using defaults")
	}

	return GlobalConfig{
		App: AppConfig{
			Env:   getKey(envMap, "APP_ENV", DefaultConfig.App.Env),
			Debug: getKeyAsBool(envMap, "APP_DEBUG", DefaultConfig.App.Debug),
			Port:  getKeyAsInt(envMap, "APP_PORT", DefaultConfig.App.Port),
			Addr:  getKey(envMap, "APP_ADDR", DefaultConfig.App.Addr),
		},
		DB: DBConfig{
			Driver:      getKey(envMap, "DB_DRIVER", DefaultConfig.DB.Driver),
			Name:        getKey(envMap, "DB_NAME", DefaultConfig.DB.Name),
			User:        getKey(envMap, "DB_USER", DefaultConfig.DB.User),
			Pass:        getKey(envMap, "DB_PASS", DefaultConfig.DB.Pass),
			Host:        getKey(envMap, "DB_HOST", DefaultConfig.DB.Host),
			Port:        getKey(envMap, "DB_PORT", DefaultConfig.DB.Port),
			ExposedPort: getKey(envMap, "DB_EXPOSED_PORT", DefaultConfig.DB.ExposedPort),
		},
		Jwt: JwtConfig{Secret: getKey(envMap, "JWT_SECRET", DefaultConfig.Jwt.Secret)},
	}
}
