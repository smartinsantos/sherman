package config

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/joho/godotenv"
	"log"
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
	Driver string
	Name   string
	User   string
	Pass   string
	Host   string
	Port   string
}

// Config Wrapper for all configurations
type Config struct {
	App AppConfig
	Db  DBConfig
}

var defaultConfig = &Config {
	App: AppConfig {
		Env:   	"local",
		Debug: 	true,
		Addr:  	":8080",
	},
	Db: DBConfig {
		Driver: "mysql",
		Name:   "db_name",
		User:   "db_user",
		Pass:   "db_password",
		Host:   "db_host",
		Port:   "db_port",
	},
}

var config *Config
var once sync.Once

// Get returns Config instance
func Get() *Config {
	once.Do(func() {
		var err error

		appDir, err := os.Getwd()
		if err != nil {
			log.Fatalln("could't read root directory")
		}

		rootBox := packr.New("root", appDir)
		envStr, err := rootBox.FindString(".env")
		if err != nil {
			log.Fatalln("no .env file found")
		}

		envMap, err := godotenv.Unmarshal(envStr)
		if err != nil {
			log.Fatalln("couldn't read contents of .env file")
		}

		config = &Config{
			App: AppConfig{
				Env:   getKey(envMap, "APP_ENV", defaultConfig.App.Env),
				Debug: getKeyAsBool(envMap, "APP_DEBUG", defaultConfig.App.Debug),
				Addr:  getKey(envMap, "APP_ADDR", defaultConfig.App.Addr),
			},
			Db: DBConfig{
				Driver: getKey(envMap, "DB_DRIVER", defaultConfig.Db.Driver),
				Name:   getKey(envMap, "DB_NAME", defaultConfig.Db.Name),
				User:   getKey(envMap, "DB_USER", defaultConfig.Db.User),
				Pass:   getKey(envMap, "DB_PASS", defaultConfig.Db.Pass),
				Host:   getKey(envMap, "DB_HOST", defaultConfig.Db.Host),
				Port:   getKey(envMap, "DB_PORT", defaultConfig.Db.Port),
			},
		}
	})
	return config
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

//func getKeyAsInt(env map[string]string, key string, defaultValue int) int {
//	valueStr := getKey(env, key, "")
//	if value, err := strconv.Atoi(valueStr); err == nil {
//		return value
//	}
//	return defaultValue
//}
