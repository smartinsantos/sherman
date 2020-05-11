package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"root/src/delivery/handler/presenter"
	"root/src/delivery/handler/validator"
	"root/src/domain"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(userUseCase domain.UserUseCase) *UserHandler {
	return &UserHandler {
		userUseCase: userUseCase,
	}
}

// Registers the user
func (uh *UserHandler) Register (context *gin.Context) {
	var userParams domain.User

	if err := context.ShouldBindJSON(&userParams); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"message": "Invalid json",
			"error": err,
		})
		return
	}

	errors := validator.ValidateUserParams(&userParams, "register")
	if len(errors) > 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"message": "Validation error",
			"errors": errors,
		})
		return
	}

	user, err := uh.userUseCase.CreateUser(&userParams)

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
		"data": presenter.PresentUser(user),
	})
}

// Logs the user
func (uh *UserHandler) Login (context *gin.Context) {
	var userParams domain.User

	if err := context.ShouldBindJSON(&userParams); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"message": "Invalid json",
			"error": err,
		})
		return
	}

	errors := validator.ValidateUserParams(&userParams, "login")
	if len(errors) > 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"message": "Validation error",
			"errors": errors,
		})
		return
	}

	user, err := uh.userUseCase.Login(&userParams)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"message": "Unable to login user",
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"message": "User logged in",
		"data": gin.H{
			"token": "todo",
			"user": presenter.PresentUser(user),
		},
	})
}

// Refreshes user token
func (uh *UserHandler) RefreshToken (context *gin.Context) {
	context.String(http.StatusOK, "RefreshToken")
}
