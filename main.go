package main

import (
	_ "root/config"
	"root/src/app"
)

func main() {
	app.Serve()
}