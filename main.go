package main

import (
	_ "github.com/go-sql-driver/mysql"
	"root/src/app"
)

func main() {
	app.Mount()
}