package main

import (
	_ "github.com/go-sql-driver/mysql"
	"sherman/src/app"
	_ "sherman/src/app/config"
)

func main() {
	app.Run()
}
