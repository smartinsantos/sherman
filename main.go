package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "sherman/src/app/config"
	"sherman/src/app/router"
)

func main() {
	router.Serve()
}
