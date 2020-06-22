package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"sherman/src/app/config"
	"sherman/src/app/registry"
	"sherman/src/delivery/handler"
	cmw "sherman/src/service/middleware"
	cmc "sherman/src/service/middleware/config"
)

// Run mounts the base application router
func Run() {
	cfg := config.Get()
	if cfg.App.Debug {
		// pretty logger
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Info().Msg("Server Running on DEBUG mode")
	}

	diContainer, err := registry.GetAppContainer()
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	router := echo.New()
	router.Use(emw.Recover())
	router.Use(emw.CORSWithConfig(cmc.CustomCorsConfig))

	cmws := diContainer.Get("middleware-service").(cmw.Middleware)
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
		userHandler := diContainer.Get("user-handler").(handler.UserHandler)

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
