package config

import (
	"github.com/joho/godotenv"
	"log"
	"strconv"
	"sync"
)

type (
	appConfig struct {
		Env		string
		Debug	bool
		Addr	string
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
		Db  dbConfig
		Jwt jwtConfig
	}
)

var (
	defaultConfig = &Config {
		App: appConfig {
			Env:   	"local",
			Debug: 	true,
			Addr:  	":8080",
		},
		Db: dbConfig {
			Driver: "mysql",
			Name:   "db_name",
			User:   "db_user",
			Pass:   "db_password",
			Host:   "db_host",
			Port:   "db_port",
		},
		Jwt: jwtConfig {
			Secret: "jwt_secret",
		},
	}
	config *Config
	once sync.Once
)

// Get returns Config instance
func Get() *Config {
	once.Do(func() {
		envMap, err := godotenv.Read(".env")

		if err != nil {
			log.Fatalln("couldn't read contents of .env file")
		}

		config = &Config {
			App: appConfig {
				Env:   getKey(envMap, "APP_ENV", defaultConfig.App.Env),
				Debug: getKeyAsBool(envMap, "APP_DEBUG", defaultConfig.App.Debug),
				Addr:  getKey(envMap, "APP_ADDR", defaultConfig.App.Addr),
			},
			Db: dbConfig {
				Driver: getKey(envMap, "DB_DRIVER", defaultConfig.Db.Driver),
				Name:   getKey(envMap, "DB_NAME", defaultConfig.Db.Name),
				User:   getKey(envMap, "DB_USER", defaultConfig.Db.User),
				Pass:   getKey(envMap, "DB_PASS", defaultConfig.Db.Pass),
				Host:   getKey(envMap, "DB_HOST", defaultConfig.Db.Host),
				Port:   getKey(envMap, "DB_PORT", defaultConfig.Db.Port),
			},
			Jwt: jwtConfig { Secret: getKey(envMap, "JWT_SECRET", defaultConfig.Jwt.Secret) },
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
