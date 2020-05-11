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
func (h *UserHandler) Register (ctx *gin.Context) {
	var user domain.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
		return
	}

	errors := validator.ValidateUserParams(&user, "register")
	if len(errors) > 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"errors": errors,
		})
		return
	}

	if err := h.userUseCase.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error": err,
		})

	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": presenter.PresentUser(&user),
	})
}

// Logs the user
func (h *UserHandler) Login (ctx *gin.Context) {
	var user domain.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
		return
	}

	if errors := validator.ValidateUserParams(&user, "login"); len(errors) > 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"errors": errors,
		})
		return
	}

	userRecord, err := h.userUseCase.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data": gin.H{
			"token": "",
			"user": presenter.PresentUser(&userRecord),
		},
	})
}
