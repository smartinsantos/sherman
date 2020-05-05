package router

import (
	"github.com/gin-gonic/gin"
	"github.com/smartinsantos/go-auth-api/config"
	"github.com/smartinsantos/go-auth-api/controller"
	"github.com/smartinsantos/go-auth-api/router/middleware"
)

// App router constructor
func New(appController *controller.AppController) *gin.Engine {
	env := config.Get()
	if !env.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	{
		r.GET("/", func(context *gin.Context) {
			context.String(200, "Hello from /")
		})
	}

	v1g := r.Group("/api/v1")
	{
		// users
		v1g.GET("/user/auth", appController.User.VerifyAuth)
		v1g.POST("/user/register", appController.User.Register)
		v1g.POST("/user/login", appController.User.Login)
		v1g.POST("/user/refresh-token", appController.User.RefreshToken)
	}

	return r
}