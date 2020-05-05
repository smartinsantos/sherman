package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smartinsantos/go-auth-api/domain/entitity"
	"github.com/smartinsantos/go-auth-api/infrastructure/datastore"
	"net/http"
)

// UserHandler struct defines the dependencies that will be used
type UserHandler struct {
	ds *datastore.UserDataStore
}

// Registers the user
func (uc * UserHandler) Register (context *gin.Context) {
	mockUser := entitity.User{
		EmailAddress: "mock1@mock.com",
		FirstName: "mock",
		LastName: "mock",
		Password: "mockPassword",
	}

	uc.ds.CreateUser(&mockUser)

	context.String(http.StatusOK, "Register")
}

// Logs the user
func (uc * UserHandler) Login (context *gin.Context) {
	context.String(http.StatusOK, "Login")
}

// Refreshes user token
func (uc * UserHandler) RefreshToken (context *gin.Context) {
	context.String(http.StatusOK, "RefreshToken")
}

// Verify that the user token is still valid
func (uc * UserHandler) VerifyAuth (context *gin.Context) {
	context.String(http.StatusOK, "VerifyAuth")
}


