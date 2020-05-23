package router

import (
	"github.com/labstack/echo"
	echoMiddleWare "github.com/labstack/echo/middleware"
	"github.com/rs/zerolog/log"
	"root/src/app/config"
	"root/src/app/registry"
	"root/src/delivery/handler"
	"root/src/utils/middleware"
)

// Serve mounts the base application router
func Serve() {
	cfg := config.Get()
	if cfg.App.Debug {
		log.Info().Msg("Server Running on DEBUG mode")
	}

	diContainer, err := registry.GetAppContainer()
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	// root router : /
	router := echo.New()
	router.Use(middleware.CORSMiddleware)
	if cfg.App.Debug {
		router.Use(echoMiddleWare.Logger())
	}

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
	log.Fatal().Err(router.Start(cfg.App.Addr))
}
