package router

import (
	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/sarulabs/di"
	"net/http"
	"sherman/src/delivery/handler"
	cmw "sherman/src/service/middleware"
	cmc "sherman/src/service/middleware/config"
)

// New creates an instance of application router
func New(ctn di.Container) *echo.Echo {
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

	return router
}
