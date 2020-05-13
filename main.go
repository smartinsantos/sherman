package main

import (
	_ "root/config"
	"root/src/app/router"
)

func main() {
	router.Serve()
}
