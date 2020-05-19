package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"root/config"
	"root/src/app/registry"
	"root/src/delivery/handler"
	"root/src/utils/middleware"
)

// Serve mounts the base application router
func Serve() {
	cfg := config.Get()
	if cfg.App.Debug {
		log.Println("Server Running on DEBUG mode")
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	diContainer, err := registry.GetAppContainer()
	if err != nil {
		log.Fatalln(err.Error())
	}
	// root router
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/", func(context *gin.Context) {
		context.String(200, "Hello from /")
	})

	// handlers
	userHandler := diContainer.Get("user-handler").(*handler.UserHandler)

	// v1 /api/v1/ group (auth-less)
	v1 := r.Group("/api/v1")
	v1.POST("/user/register", userHandler.Register)
	v1.POST("/user/login", userHandler.Login)

	// v1a api/v1/ group (auth)
	v1.Use(middleware.AuthMiddleware())
	v1.GET("/protected", func(context *gin.Context) {
		context.String(200, "Protected resource")
	})

	// run the server
	log.Println("Server Running on PORT", cfg.App.Addr)
	log.Fatalln(r.Run(cfg.App.Addr))
}
