package config

import (
	"os"
	"strconv"
	"sync"
)

// App configuration
type AppConfig struct {
    Env		string
    Debug	bool
    Addr	string
}

// Database configuration
type DBConfig struct {
	Driver 	string
	Name	string
	User	string
	Pass	string
	Host	string
	Port	string
}

// Config Wrapper for all configurations
type Config struct {
    AppConfig AppConfig
    DBConfig DBConfig
}

var instance *Config
var once sync.Once

// Get returns Config instance
func Get() *Config {
	once.Do(func () {
		instance = &Config{
			AppConfig: AppConfig{
				Env: getEnv("APP_ENV", "local"),
				Debug: getEnvAsBool("APP_DEBUG", true),
				Addr: getEnv("APP_ADDR", ":8080"),
			},
			DBConfig: DBConfig{
				Driver: getEnv("DB_DRIVER", "db-driver"),
				Name: getEnv("DB_NAME", "db_name"),
				User:	getEnv("DB_USER", "db_user"),
				Pass: getEnv("DB_PASS", "db_password"),
				Host:	getEnv("DB_HOST", "db_user"),
				Port: getEnv("DB_PORT", "db_port"),
			},
		}
	})
	return instance
}

func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
		return value
    }
    return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
		return value
    }
    return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
    valStr := getEnv(name, "")
    if val, err := strconv.ParseBool(valStr); err == nil {
		return val
    }
    return defaultVal
}