package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"sync"
)

type (
	// AppConfig type definition
	AppConfig struct {
		Debug bool
		Port  int
		Addr  string
	}
	// DBConfig type definition
	DBConfig struct {
		Driver      string
		Name        string
		User        string
		Pass        string
		Host        string
		Port        string
		ExposedPort string
		Path        string
	}
	// JwtConfig type definition
	JwtConfig struct {
		Secret string
	}
	// GlobalConfig type definition
	GlobalConfig struct {
		App AppConfig
		DB  DBConfig
		Jwt JwtConfig
	}
)

var (
	once   sync.Once
	config *GlobalConfig
	// DefaultConfig contains default values of global config
	DefaultConfig = GlobalConfig{
		App: AppConfig{
			Debug: false,
			Port:  5000,
			Addr:  ":5000",
		},
		DB: DBConfig{
			Driver:      "mysql",
			Name:        "sherman",
			User:        "db_user",
			Pass:        "db_password",
			Host:        "app-mysql",
			Port:        "3306",
			ExposedPort: "5001",
		},
		Jwt: JwtConfig{
			Secret: "jwt_secret",
		},
	}
)

// Get returns singleton instance of GlobalConfig
func Get() *GlobalConfig {
	if env := os.Getenv("ENV"); env == "test" {
		envMap, err := godotenv.Read(".env.test")
		if err != nil {
			log.Error().Msg("config error: couldn't read env.test file, using defaults")
		}
		return generateConfig(envMap)
	}

	once.Do(func() {
		envMap, err := godotenv.Read(".env")
		if err != nil {
			log.Error().Msg("config error: couldn't read .env file, using defaults")
		}
		config = generateConfig(envMap)
	})

	return config
}

func generateConfig(envMap map[string]string) *GlobalConfig {
	return &GlobalConfig{
		App: AppConfig{
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
			Path:        getKey(envMap, "DB_PATH", ""),
		},
		Jwt: JwtConfig{Secret: getKey(envMap, "JWT_SECRET", DefaultConfig.Jwt.Secret)},
	}
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
