package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"root/src/delivery/handler/presenter"
	"root/src/delivery/handler/validator"
	"root/src/domain/auth"
	"root/src/utils/exception"
	"root/src/utils/response"
)

// UserHandler handler for /user/[routes]
type UserHandler struct {
	UserUseCase auth.UserUseCase
}

// Register registers the user
func (h *UserHandler) Register(ctx *gin.Context) {
	var user auth.User
	res := response.NewResponse()

	if err := ctx.BindJSON(&user); err != nil {
		res.SetError(http.StatusUnprocessableEntity, err.Error())
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	errors := validator.ValidateUserParams(&user, "register")
	if len(errors) > 0 {
		res.SetErrors(http.StatusUnprocessableEntity, errors)
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	if err := h.UserUseCase.Register(&user); err != nil {
		switch err.(type) {
		case *exception.DuplicateEntryError:
			res.SetError(http.StatusForbidden, err.Error())
		default:
			res.SetInternalServerError()
		}
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	res.SetData(http.StatusCreated, gin.H {
		"user": presenter.PresentUser(&user),
	})
	ctx.JSON(res.GetStatus(), res.GetBody())
}

// Login logs the user in
func (h *UserHandler) Login(ctx *gin.Context) {
	var user auth.User
	res := response.NewResponse()

	if err := ctx.BindJSON(&user); err != nil {
		res.SetError(http.StatusUnprocessableEntity, err.Error())
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	if errors := validator.ValidateUserParams(&user, "login"); len(errors) > 0 {
		res.SetErrors(http.StatusUnprocessableEntity, errors)
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	userRecord, refreshToken, accessToken, err := h.UserUseCase.Login(&user)
	if err != nil {
		switch err.(type) {
		case *exception.NotFoundError:
			res.SetError(http.StatusNotFound, err.Error())
		case *exception.UnAuthorizedError:
			res.SetError(http.StatusUnauthorized, err.Error())
		default:
			res.SetInternalServerError()
		}
		ctx.JSON(res.GetStatus(), res.GetBody())
		return
	}

	res.SetData(http.StatusOK, gin.H {
		"access_token": accessToken,
		"user": presenter.PresentUser(&userRecord),
	})

	//@TODO: add secure to cookie when tls is ready
	ctx.SetCookie("REFRESH_TOKEN", refreshToken, 3600, "/", ctx.Request.Host, false, true)
	ctx.JSON(res.GetStatus(), res.GetBody())
}

// RefreshAccessToken refreshes user access token
func (h *UserHandler) RefreshAccessToken(ctx *gin.Context) {


	ctx.String(200, "RefreshAccessToken")
}

// Logouts the user
func (h *UserHandler) Logout(ctx *gin.Context) {
	ctx.String(200, "Logout")
}