package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sherman/src/delivery/handler/presenter"
	"sherman/src/delivery/handler/validator"
	"sherman/src/domain/auth"
	"sherman/src/utils/exception"
	"sherman/src/utils/response"
	"sherman/src/utils/security"
)

// UserHandler handler for /user/[routes]
type UserHandler struct {
	UserUseCase auth.UserUseCase
	SecurityTokenUseCase auth.SecurityTokenUseCase
}

// Register registers the user
func (h *UserHandler) Register(ctx echo.Context) error {
	var user auth.User
	res := response.NewResponse()

	if err := ctx.Bind(&user); err != nil {
		res.SetError(http.StatusUnprocessableEntity, err.Error())
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	errors := validator.ValidateUserParams(&user, "register")
	if len(errors) > 0 {
		res.SetErrors(http.StatusUnprocessableEntity, errors)
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	if err := h.UserUseCase.Register(&user); err != nil {
		switch err.(type) {
		case *exception.DuplicateEntryError:
			res.SetError(http.StatusForbidden, err.Error())
		default:
			res.SetInternalServerError()
		}
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	res.SetData(http.StatusCreated, nil)
	return ctx.JSON(res.GetStatus(), res.GetBody())
}

// Login logs the user in
func (h *UserHandler) Login(ctx echo.Context) error {
	var user auth.User
	res := response.NewResponse()

	if err := ctx.Bind(&user); err != nil {
		res.SetError(http.StatusUnprocessableEntity, err.Error())
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	if errors := validator.ValidateUserParams(&user, "login"); len(errors) > 0 {
		res.SetErrors(http.StatusUnprocessableEntity, errors)
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	verifiedUser, err := h.UserUseCase.VerifyCredentials(&user)

	if err != nil {
		switch err.(type) {
		case *exception.NotFoundError:
			res.SetError(http.StatusNotFound, err.Error())
		case *exception.UnAuthorizedError:
			res.SetError(http.StatusUnauthorized, err.Error())
		default:
			res.SetInternalServerError()
		}
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	accessToken, err := h.SecurityTokenUseCase.GenAccessToken(verifiedUser.ID)
	if err != nil {
		res.SetInternalServerError()
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	refreshToken, err := h.SecurityTokenUseCase.GenRefreshToken(verifiedUser.ID)
	if err != nil {
		res.SetInternalServerError()
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	res.SetData(http.StatusOK, response.D{ "access_token": accessToken.Token })
	//@TODO: add secure to cookie when tls is ready
	ctx.SetCookie(&http.Cookie{
		Name: "REFRESH_TOKEN",
		Value: refreshToken.Token,
		MaxAge: 3600,
		Path: "/",
		Domain: ctx.Request().Host,
		Secure: false,
		HttpOnly: true,
	})
	return ctx.JSON(res.GetStatus(), res.GetBody())
}

// RefreshAccessToken refreshes user access token
func (h *UserHandler) RefreshAccessToken(ctx echo.Context) error {
	res := response.NewResponse()

	refreshTokenMetadata, err := security.GetAndValidateRefreshToken(ctx)
	if err != nil {
		res.SetError(http.StatusUnauthorized, "invalid refresh token")
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	if !h.SecurityTokenUseCase.IsRefreshTokenStored(&refreshTokenMetadata) {
		res.SetError(http.StatusUnauthorized, "invalid refresh token")
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	accessToken, err := h.SecurityTokenUseCase.GenAccessToken(refreshTokenMetadata.UserID)
	if err != nil {
		res.SetError(http.StatusUnauthorized, err.Error())
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	res.SetData(http.StatusOK, response.D{ "access_token": accessToken.Token })
	return ctx.JSON(res.GetStatus(), res.GetBody())
}

// GetUser gets the user from access token
func (h *UserHandler) GetUser(ctx echo.Context) error {
	res := response.NewResponse()
	userID := ctx.Param("id")

	user, err := h.UserUseCase.GetUserByID(userID)

	if err != nil {
		switch err.(type) {
		case *exception.NotFoundError:
			res.SetError(http.StatusNotFound, err.Error())
		default:
			res.SetInternalServerError()
		}
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	res.SetData(http.StatusOK, response.D{ "user": presenter.PresentUser(&user) })
	return ctx.JSON(res.GetStatus(), res.GetBody())
}

// Logout logs out the user
func (h *UserHandler) Logout(ctx echo.Context) error {
	res := response.NewResponse()

	refreshTokenMetadata, err := security.GetAndValidateRefreshToken(ctx)
	if err != nil {
		res.SetError(http.StatusUnauthorized, err.Error())
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	err = h.SecurityTokenUseCase.RemoveRefreshToken(&refreshTokenMetadata)
	if err != nil {
		res.SetInternalServerError()
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	//@TODO: add secure to cookie when tls is ready
	ctx.SetCookie(&http.Cookie{
		Name: "REFRESH_TOKEN",
		Value: "",
		MaxAge: 0,
		Path: "/",
		Domain: ctx.Request().Host,
		Secure: false,
		HttpOnly: true,
	})
	res.SetData(http.StatusOK, nil)
	return ctx.JSON(res.GetStatus(), res.GetBody())
}