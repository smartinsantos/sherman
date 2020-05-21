package router

import (
	"github.com/labstack/echo"
	"log"
	"root/config"
	"root/src/app/registry"
	"root/src/delivery/handler"

	//"root/src/app/registry"
	//"root/src/delivery/handler"
	"root/src/utils/middleware"
)

// Serve mounts the base application router
func Serve() {
	cfg := config.Get()
	if cfg.App.Debug {
		log.Println("Server Running on DEBUG mode")
	}

	diContainer, err := registry.GetAppContainer()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// root router : /
	router := echo.New()
	router.Use(middleware.CORSMiddleware)

	router.GET("/", func(ctx echo.Context) error {
		return ctx.String(200, "Hello from /")
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
		userRouter.GET("/:id", userHandler.GetUser, middleware.UserAuthMiddleware)
		userRouter.GET("/logout", userHandler.Logout, middleware.UserAuthMiddleware)
	}

	// run the server
	log.Fatalln(router.Start(cfg.App.Addr))
}
