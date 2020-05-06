package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smartinsantos/go-auth-api/domain"
	"net/http"
)

// UserHandler struct defines the dependencies that will be used
type UserHandler struct {
	userUseCase domain.UserUseCase
}

// Registers the user
func (uh *UserHandler) Register (context *gin.Context) {
	var user domain.User

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

	newUser, err := uh.userUseCase.CreateUser(&user)

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
func (uh *UserHandler) Login (context *gin.Context) {
	context.String(http.StatusOK, "Login")
}

// Refreshes user token
func (uh *UserHandler) RefreshToken (context *gin.Context) {
	context.String(http.StatusOK, "RefreshToken")
}
