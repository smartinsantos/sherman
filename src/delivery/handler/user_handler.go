package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"root/src/app/exception"
	"root/src/delivery/handler/presenter"
	"root/src/delivery/handler/response"
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
	var res response.Response

	if err := ctx.BindJSON(&user); err != nil {
		res = response.Response {
			Status: http.StatusUnprocessableEntity,
			Error: err.Error(),
		}
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	errors := validator.ValidateUserParams(&user, "register")
	if len(errors) > 0 {
		res = response.Response {
			Status: http.StatusUnprocessableEntity,
			Errors: errors,
		}
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	if err := h.UserUseCase.CreateUser(&user); err != nil {
		switch err.(type) {
		case *exception.DuplicateEntryError:
			res = response.Response {
				Status: http.StatusForbidden,
				Error: err.Error(),
			}
		default:
			res = response.Response {
				Status: http.StatusForbidden,
				Error: "internal server error",
			}
		}
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	res = response.Response {
		Status: http.StatusCreated,
		Data: gin.H {
			"user": presenter.PresentUser(&user),
		},
	}
	ctx.JSON(res.GetStatus(), res.GetBody())
}

// Login logs the user in
func (h *UserHandler) Login (ctx *gin.Context) {
	var user domain.User
	var res response.Response

	if err := ctx.BindJSON(&user); err != nil {
		res = response.Response {
			Status: http.StatusUnprocessableEntity,
			Error: err.Error(),
		}
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	if errors := validator.ValidateUserParams(&user, "login"); len(errors) > 0 {
		res = response.Response {
			Status: http.StatusUnprocessableEntity,
			Errors: errors,
		}
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	userRecord, err := h.UserUseCase.Login(&user)
	if err != nil {
		switch err.(type) {
		case *exception.NotFoundError:
			res = response.Response {
				Status: http.StatusNotFound,
				Error: err.Error(),
			}
		case *exception.UnAuthorizedError:
			res = response.Response {
				Status: http.StatusUnauthorized,
				Error: err.Error(),
			}
		default:
			res = response.Response {
				Status: http.StatusInternalServerError,
				Error: "internal server error",
			}
		}
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	res = response.Response {
		Status: http.StatusOK,
		Data: gin.H {
			"token": "",
			"user": presenter.PresentUser(&userRecord),
		},
	}
	ctx.JSON(res.GetStatus(), res.GetBody())
}
