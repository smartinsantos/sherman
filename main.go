package main

import (
	_ "github.com/go-sql-driver/mysql"
	"root/app"
)

func main() {
	app.Mount()
}