package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smartinsantos/go-auth-api/delivery/handler/presenter"
	"github.com/smartinsantos/go-auth-api/delivery/handler/validator"
	"github.com/smartinsantos/go-auth-api/domain"
	"net/http"
)


type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(userUseCase domain.UserUseCase) UserHandler {
	return UserHandler {
		userUseCase: userUseCase,
	}
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

	errors := validator.ValidateUserParams(&user, "register")
	if len(errors) > 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"message": "Validation error",
			"errors": errors,
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
		"data": presenter.PresentUser(newUser),
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
