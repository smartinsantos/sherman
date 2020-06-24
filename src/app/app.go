package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sarulabs/di"
	"net/http"
	"os"
	"sherman/src/delivery/handler"
	"sherman/src/service/config"
	cmw "sherman/src/service/middleware"
	cmc "sherman/src/service/middleware/config"
)

// Run mounts the base application router
func Run(ctn di.Container) {
	cfg := ctn.Get("config").(config.Config).Get()
	if cfg.App.Debug {
		// pretty logger
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Info().Msg("Server Running on DEBUG mode")
	}

	router := echo.New()
	router.Use(emw.Recover())
	router.Use(emw.CORSWithConfig(cmc.CustomCorsConfig))

	cmws := ctn.Get("middleware-service").(cmw.Middleware)
	router.Use(cmws.ZeroLog())

	// root routes : /
	router.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello from /")
	})
	// routes: /api/v1
	v1Router := router.Group("/api/v1")
	// routes: /api/v1/users
	userRouter := v1Router.Group("/users")
	{
		userHandler := ctn.Get("user-handler").(handler.UserHandler)

		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
		userRouter.PATCH("/refresh-token", userHandler.RefreshAccessToken)
		userRouter.GET("/:id", userHandler.GetUser, cmws.JWT())
		userRouter.DELETE("/logout", userHandler.Logout, cmws.JWT())
	}

	// run the server
	log.Info().Msg(fmt.Sprintf("Server Running on PORT%s", cfg.App.Addr))
	log.Fatal().Err(router.Start(cfg.App.Addr))
}