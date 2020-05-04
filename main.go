package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/smartinsantos/go-auth-api/config"
	"github.com/smartinsantos/go-auth-api/infrastructure/datastore"
	"github.com/smartinsantos/go-auth-api/interfaces/handler"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error: No .env file found")
	}
}

func main() {
	env := config.Get()

	ads := datastore.New()
	ah := handler.New(ads)

	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(200, "Hello from /")
	})

	// users
	r.GET("/api/v1/user/auth", ah.User.VerifyAuth)
	r.POST("/api/v1/user/register", ah.User.Register)
	r.POST("/api/v1/user/login", ah.User.Login)
	r.POST("/api/v1/user/refresh-token", ah.User.RefreshToken)

	log.Fatal(r.Run(env.AppConfig.Addr))
}