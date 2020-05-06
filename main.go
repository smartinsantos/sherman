package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/smartinsantos/go-auth-api/config"
	"github.com/smartinsantos/go-auth-api/utils/middleware"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error: No .env file found")
	}
}

func main() {
	env := config.Get()

	// init db
	DBURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBConfig.User,
		env.DBConfig.Pass,
		env.DBConfig.Host,
		env.DBConfig.Port,
		env.DBConfig.Name,
	)

	db, err := gorm.Open(env.DBConfig.Driver, DBURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

		if !env.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// init router
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	{
		r.GET("/", func(context *gin.Context) {
			context.String(200, "Hello from /")
		})
	}

	//v1g := r.Group("/api/v1")
	//{
	//	// users
	//	v1g.GET("/user/auth", appController.User.VerifyAuth)
	//	v1g.POST("/user/register", appController.User.Register)
	//	v1g.POST("/user/login", appController.User.Login)
	//	v1g.POST("/user/refresh-token", appController.User.RefreshToken)
	//}

	log.Fatal(r.Run(env.AppConfig.Addr))
}