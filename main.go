package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/smartinsantos/go-auth-api/config"
	"github.com/smartinsantos/go-auth-api/interfaces/controller"
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

	// users
	userController := controller.NewUserController()
	router.GET("/user/auth", userController.VerifyAuth)
	router.POST("/user/register", userController.Register)
	router.POST("/user/login", userController.Login)
	router.POST("/user/refresh-token", userController.Login)

	log.Fatal(router.Run(env.AppConfig.Addr))
}