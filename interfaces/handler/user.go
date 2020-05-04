package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smartinsantos/go-auth-api/infrastructure/datastore"
	"net/http"
)

// userController struct defines the dependencies that will be used
type UserHandler struct {
	ds *datastore.UserDataStore
}

func (uc * UserHandler) Register (context *gin.Context) {
	context.String(http.StatusOK, "Register")
}

func (uc * UserHandler) Login (context *gin.Context) {
	context.String(http.StatusOK, "Login")
}

func (uc * UserHandler) RefreshToken (context *gin.Context) {
	context.String(http.StatusOK, "RefreshToken")
}

func (uc * UserHandler) VerifyAuth (context *gin.Context) {
	context.String(http.StatusOK, "VerifyAuth")
}


