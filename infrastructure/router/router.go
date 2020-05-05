package router

import (
	"github.com/gin-gonic/gin"
	"github.com/smartinsantos/go-auth-api/config"
	"github.com/smartinsantos/go-auth-api/interfaces/handler"
)

// App router constructor
func New(appHandler *handler.AppHandler) *gin.Engine {
	env := config.Get()
	if !env.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	{
		r.GET("/", func(context *gin.Context) {
			context.String(200, "Hello from /")
		})
	}

	v1g := r.Group("/api/v1")
	{
		// users
		v1g.GET("/user/auth", appHandler.User.VerifyAuth)
		v1g.POST("/user/register", appHandler.User.Register)
		v1g.POST("/user/login", appHandler.User.Login)
		v1g.POST("/user/refresh-token", appHandler.User.RefreshToken)
	}

	return r
}