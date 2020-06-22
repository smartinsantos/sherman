package main

import (
	_ "github.com/go-sql-driver/mysql"
	"sherman/src/app"
	"sherman/src/app/config"
	_ "sherman/src/app/config"
)

func main() {
	config.LoadFromEnv()
	app.Run()
}
