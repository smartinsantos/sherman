package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di"
	"github.com/sarulabs/di-example/config/logging"
	"log"
	"root/config"
	"root/src/app/middleware"
	"root/src/delivery/handler"
)

// Mounts the base application router
func Mount() {
	cfg := config.Get()
	if cfg.App.Debug {
		log.Println("Server Running on DEBUG mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = builder.Add(DIContainer...)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	ctn := builder.Build()

	// main router
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	{
		r.GET("/", func(context *gin.Context) {
			context.String(200, "Hello from /")
		})
	}

	// v1 group
	v1g := r.Group("/api/v1")
	{
		userHandler := ctn.Get("user-handler").(*handler.UserHandler)
		v1g.POST("/user/register", userHandler.Register)
		v1g.POST("/user/login", userHandler.Login)
	}

	log.Println("Server Running on PORT", cfg.App.Addr)
	log.Fatal(r.Run(cfg.App.Addr))
}
