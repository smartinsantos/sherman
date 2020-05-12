package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"root/src/delivery/handler/presenter"
	"root/src/delivery/handler/validator"
	"root/src/domain"
)

// Handler for /user/[routes]
type UserHandler struct {
	UserUseCase domain.UserUseCase
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

	if err := h.UserUseCase.CreateUser(&user); err != nil {
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

	userRecord, err := h.UserUseCase.Login(&user)
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
