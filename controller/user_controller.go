package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/smartinsantos/go-auth-api/infrastructure/datastore"
	"github.com/smartinsantos/go-auth-api/model/entity"
	"net/http"
)

// UserController struct defines the dependencies that will be used
type UserController struct {
	ds *datastore.UserDataStore
}

// Registers the user
func (uc * UserController) Register (context *gin.Context) {
	var user entity.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"message": "Invalid json",
			"error": err,
		})
		return
	}

	validateErrors := user.Validate("register")
	if len(validateErrors) > 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"message": "Validation error",
			"errors": validateErrors,
		})
		return
	}

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
func (uc * UserController) Login (context *gin.Context) {
	context.String(http.StatusOK, "Login")
}

// Refreshes user token
func (uc * UserController) RefreshToken (context *gin.Context) {
	context.String(http.StatusOK, "RefreshToken")
}

// Verify that the user token is still valid
func (uc * UserController) VerifyAuth (context *gin.Context) {
	context.String(http.StatusOK, "VerifyAuth")
}


