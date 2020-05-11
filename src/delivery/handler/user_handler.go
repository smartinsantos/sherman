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
	var user domain.User

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
		return
	}

	errors := validator.ValidateUserParams(&user, "register")
	if len(errors) > 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"errors": errors,
		})
		return
	}

	if err := uh.userUseCase.CreateUser(&user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error": err,
		})

	}

	context.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": presenter.PresentUser(&user),
	})
}

// Logs the user
func (uh *UserHandler) Login (context *gin.Context) {
	var user domain.User

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
		return
	}

	if errors := validator.ValidateUserParams(&user, "login"); len(errors) > 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"errors": errors,
		})
		return
	}

	userRecord, err := uh.userUseCase.Login(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data": gin.H{
			"token": "",
			"user": presenter.PresentUser(&userRecord),
		},
	})
}
