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

	container, err := registry.NewAppContainer()
	if err != nil {
		log.Fatalln(err.Error())
	}
	// root router
	r := gin.Default()
	// root router middleware
	r.Use(middleware.CORSMiddleware())
	// root routes
	{
		r.GET("/", func(context *gin.Context) {
			context.String(200, "Hello from /")
		})
	}

	// v1 group
	v1g := r.Group("/api/v1")
	{
		userHandler := container.Get("user-handler").(*handler.UserHandler)
		v1g.POST("/user/register", userHandler.Register)
		v1g.POST("/user/login", userHandler.Login)
	}
	// run the server
	log.Println("Server Running on PORT", cfg.App.Addr)
	log.Fatalln(r.Run(cfg.App.Addr))
}
