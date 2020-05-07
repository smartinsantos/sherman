package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/smartinsantos/go-auth-api/config"
	"github.com/smartinsantos/go-auth-api/delivery/handler"
	"github.com/smartinsantos/go-auth-api/repository/datastore"
	"github.com/smartinsantos/go-auth-api/usecase"
	"github.com/smartinsantos/go-auth-api/utils/middleware"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error: No .env file found")
	}
}

func main() {
	// TODO: move all of this code to a di container
	env := config.Get()

	// init db
	connectionUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBConfig.User,
		env.DBConfig.Pass,
		env.DBConfig.Host,
		env.DBConfig.Port,
		env.DBConfig.Name,
	)

	db, err := gorm.Open(env.DBConfig.Driver, connectionUrl)
	if err != nil {
		log.Fatal(err)
	}

	// init repositories
	dsUserRepository := datastore.NewDsUserRepository(db)

	// init use cases
	userUseCase := usecase.NewUserUseCase(dsUserRepository)

	// init handlers
	userHandler := handler.NewUserHandler(userUseCase)

	// init router
	if !env.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	{
		r.GET("/", func(context *gin.Context) {
			context.String(200, "Hello from /")
		})
	}

	v1g := r.Group("/api/v1")
	{
		// users
		v1g.POST("/user/register", userHandler.Register)
	}

	log.Fatal(r.Run(env.AppConfig.Addr))
}