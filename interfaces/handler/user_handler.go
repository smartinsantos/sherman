package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smartinsantos/go-auth-api/domain/entity"
	"github.com/smartinsantos/go-auth-api/infrastructure/datastore"
	"net/http"
)

// UserHandler struct defines the dependencies that will be used
type UserHandler struct {
	ds *datastore.UserDataStore
}

// Registers the user
func (uc * UserHandler) Register (context *gin.Context) {
	var user entity.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"message": "Invalid json",
			"error": err,
		})
		return
	}

	// TODO: validate inputs
	newUser, err := uc.ds.CreateUser(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"message": "Unable to create user",
			"error": err,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"message": "User created",
		"data": newUser.Presenter(),
	})
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


