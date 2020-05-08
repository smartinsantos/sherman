package app

import (
	"github.com/gin-gonic/gin"
	"root/config"
	"root/utils/middleware"
	"log"
)

func Mount() {
	cfg := config.Get()
	if cfg.AppConfig.Debug {
		log.Println("Server Running on DEBUG mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// main router
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// init handlers
	handler, err := NewHandlers()
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
	}

	log.Println("Server Running on port", cfg.AppConfig.Addr)
	log.Fatal(r.Run(cfg.AppConfig.Addr))
}
