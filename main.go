package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "root/src/app/config"
	"root/src/app/router"
)

func main() {
	router.Serve()
}
