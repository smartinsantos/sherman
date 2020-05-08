package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"root/app"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error: No .env file found")
		panic(err)
	}
}

func main() {
	app.Mount()
}