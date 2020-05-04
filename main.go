package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/smartinsantos/go-auth-api/config"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error: No .env file found")
	}
}

func main() {
	env := config.Get()

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(200, "Hello from /")
	})

	log.Fatal(router.Run(env.AppConfig.Addr))
}