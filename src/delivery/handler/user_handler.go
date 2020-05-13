package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"root/src/delivery/handler/presenter"
	"root/src/delivery/handler/validator"
	"root/src/domain"
)

// UserHandler handler for /user/[routes]
type UserHandler struct {
	UserUseCase domain.UserUseCase
}

// Register registers the user
func (h *UserHandler) Register (ctx *gin.Context) {
	var user domain.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {
			"error": err.Error(),
			"data": nil,
		})
		return
	}

	errors := validator.ValidateUserParams(&user, "register")
	if len(errors) > 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {
			"errors": errors,
			"data": nil,
		})
		return
	}

	if err := h.UserUseCase.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {
			"error": err,
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H {
		"data": presenter.PresentUser(&user),
	})
}

// Login logs the user in
func (h *UserHandler) Login (ctx *gin.Context) {
	var user domain.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {
			"error": err.Error(),
			"data": nil,
		})
		return
	}

	if errors := validator.ValidateUserParams(&user, "login"); len(errors) > 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H {
			"errors": errors,
			"data": nil,
		})
		return
	}

	userRecord, err := h.UserUseCase.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H {
		"data": gin.H {
			"token": "",
			"user": presenter.PresentUser(&userRecord),
		},
	})
}
