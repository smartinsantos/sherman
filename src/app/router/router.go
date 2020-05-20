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
	// root router : /
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello from /")
	})

	// router: /api/v1
	v1Router := router.Group("/api/v1")

	// router: /api/v1/users
	userRouter := v1Router.Group("/users")
	userHandler := diContainer.Get("user-handler").(*handler.UserHandler)
	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
		userRouter.GET("/refresh-token", userHandler.RefreshAccessToken)

		// auth middleware protected routes
		userRouter.Use(middleware.UserAuthMiddleware())
		userRouter.GET("/user/:id", userHandler.GetUser)
		userRouter.GET("/logout", userHandler.Logout)
	}

	// run the server
	log.Println("Server Running on PORT", cfg.App.Addr)
	log.Fatalln(router.Run(cfg.App.Addr))
}
