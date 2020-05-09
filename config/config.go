package config

import (
	"github.com/gobuffalo/packr"
	"github.com/joho/godotenv"
	"log"
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

var instance = &Config {
	AppConfig: AppConfig {
		Env: "local",
		Debug: true,
		Addr: ":8080",
	},
	DBConfig: DBConfig {
		Driver: "mysql",
		Name: "db_name",
		User: "db_user",
		Pass: "db_password",
		Host: "db_host",
		Port: "db_port",
	},
}
var once sync.Once

// Get returns Config instance
func Get() *Config {
	once.Do(func () {
		rootBox := packr.NewBox("../")
		envStr, err := rootBox.FindString(".env")
		if err != nil {
			log.Println("Error: No .env file found")
		}

		envMap, readErr := godotenv.Unmarshal(envStr)
		if readErr != nil {
			log.Println("Error: Cannot read contents of .env file")
		}

		instance = &Config{
			AppConfig: AppConfig{
				Env: getKey(envMap,"APP_ENV", instance.AppConfig.Env),
				Debug: getKeyAsBool(envMap,"APP_DEBUG", instance.AppConfig.Debug),
				Addr: getKey(envMap,"APP_ADDR", instance.AppConfig.Addr),
			},
			DBConfig: DBConfig{
				Driver: getKey(envMap,"DB_DRIVER", instance.DBConfig.Driver),
				Name: getKey(envMap,"DB_NAME", instance.DBConfig.Name),
				User:	getKey(envMap,"DB_USER", instance.DBConfig.User),
				Pass: getKey(envMap,"DB_PASS", instance.DBConfig.Pass),
				Host:	getKey(envMap,"DB_HOST", instance.DBConfig.Host),
				Port: getKey(envMap,"DB_PORT", instance.DBConfig.Port),
			},
		}
	})
	return instance
}

func getKey(env map[string]string, key string, defaultValue string) string {
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