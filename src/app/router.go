package app

import (
	"log"

	"github.com/gin-gonic/gin"

	"root/config"
	"root/src/app/middleware"
)

func Mount() {
	cfg := config.Get()
	if cfg.App.Debug {
		log.Println("Server Running on DEBUG mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// main router
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// init handlers
	handler, err := newHandlers()
	if err != nil {
		log.Fatalln(err)
	}

	{
		r.GET("/", func(context *gin.Context) {
			context.String(200, "Hello from /")
		})
	}

	// v1 group
	v1g := r.Group("/api/v1")
	{
		v1g.POST("/user/register", handler.userHandler.Register)
		v1g.POST("/user/login", handler.userHandler.Login)
	}

	log.Println("Server Running on PORT", cfg.App.Addr)
	log.Fatal(r.Run(cfg.App.Addr))
}
