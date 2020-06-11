package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sherman/src/app/utils/exception"
	"sherman/src/app/utils/response"
	"sherman/src/domain/auth"
	"sherman/src/service/presenter"
	"sherman/src/service/security"
	"sherman/src/service/validator"
)

type (
	// UserHandler handler for /user/[routes]
	UserHandler interface {
		Register(ctx echo.Context) error
		Login(ctx echo.Context) error
		RefreshAccessToken(ctx echo.Context) error
		GetUser(ctx echo.Context) error
		Logout(ctx echo.Context) error
	}

	userHandler struct {
		userUseCase          auth.UserUseCase
		securityTokenUseCase auth.SecurityTokenUseCase
		validator            validator.Validator
		security             security.Security
		presenter            presenter.Presenter
	}
)

// NewUserHandler constructor
func NewUserHandler(
	uuc auth.UserUseCase,
	stuc auth.SecurityTokenUseCase,
	vs validator.Validator,
	ss security.Security,
	ps presenter.Presenter,
) UserHandler {
	return &userHandler{
		userUseCase:          uuc,
		securityTokenUseCase: stuc,
		validator:            vs,
		security:             ss,
		presenter:            ps,
	}
}

// Register registers the user
func (h *userHandler) Register(ctx echo.Context) error {
	var user auth.User
	res := response.NewResponse()

	if err := ctx.Bind(&user); err != nil {
		res.SetInternalServerError()
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	errors := h.validator.ValidateUserParams(&user, "register")
	if len(errors) > 0 {
		res.SetErrors(http.StatusUnprocessableEntity, errors)
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	if err := h.userUseCase.Register(&user); err != nil {
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
func (h *userHandler) Login(ctx echo.Context) error {
	var user auth.User
	res := response.NewResponse()

	if err := ctx.Bind(&user); err != nil {
		res.SetInternalServerError()
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	if errors := h.validator.ValidateUserParams(&user, "login"); len(errors) > 0 {
		res.SetErrors(http.StatusUnprocessableEntity, errors)
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	verifiedUser, err := h.userUseCase.VerifyCredentials(&user)

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

	accessToken, err := h.securityTokenUseCase.GenAccessToken(verifiedUser.ID)
	if err != nil {
		res.SetInternalServerError()
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	refreshToken, err := h.securityTokenUseCase.GenRefreshToken(verifiedUser.ID)
	if err != nil {
		res.SetInternalServerError()
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	res.SetData(http.StatusOK, response.D{"access_token": accessToken.Token})
	// TODO: add secure to cookie when tls is ready
	ctx.SetCookie(&http.Cookie{
		Name:     "REFRESH_TOKEN",
		Value:    refreshToken.Token,
		MaxAge:   3600,
		Path:     "/",
		Domain:   ctx.Request().Host,
		Secure:   false,
		HttpOnly: true,
	})
	return ctx.JSON(res.GetStatus(), res.GetBody())
}

// RefreshAccessToken refreshes user access token
func (h *userHandler) RefreshAccessToken(ctx echo.Context) error {
	res := response.NewResponse()

	refreshTokenMetadata, err := h.security.GetAndValidateRefreshToken(ctx)
	if err != nil {
		res.SetError(http.StatusUnauthorized, "invalid refresh token")
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	if !h.securityTokenUseCase.IsRefreshTokenStored(&refreshTokenMetadata) {
		res.SetError(http.StatusUnauthorized, "invalid refresh token")
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	accessToken, err := h.securityTokenUseCase.GenAccessToken(refreshTokenMetadata.UserID)
	if err != nil {
		res.SetError(http.StatusUnauthorized, err.Error())
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	res.SetData(http.StatusOK, response.D{"access_token": accessToken.Token})
	return ctx.JSON(res.GetStatus(), res.GetBody())
}

// GetUser gets the user from access token
func (h *userHandler) GetUser(ctx echo.Context) error {
	res := response.NewResponse()
	userID := ctx.Param("id")

	user, err := h.userUseCase.GetUserByID(userID)
	if err != nil {
		switch err.(type) {
		case *exception.NotFoundError:
			res.SetError(http.StatusNotFound, err.Error())
		default:
			res.SetInternalServerError()
		}
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	res.SetData(http.StatusOK, response.D{"user": h.presenter.PresentUser(&user)})
	return ctx.JSON(res.GetStatus(), res.GetBody())
}

// Logout logs out the user
func (h *userHandler) Logout(ctx echo.Context) error {
	res := response.NewResponse()

	refreshTokenMetadata, err := h.security.GetAndValidateRefreshToken(ctx)
	if err != nil {
		res.SetError(http.StatusUnauthorized, err.Error())
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	if err := h.securityTokenUseCase.RemoveRefreshToken(&refreshTokenMetadata); err != nil {
		res.SetInternalServerError()
		return ctx.JSON(res.GetStatus(), res.GetBody())
	}

	// TODO: add secure to cookie when tls is ready
	ctx.SetCookie(&http.Cookie{
		Name:     "REFRESH_TOKEN",
		Value:    "",
		MaxAge:   0,
		Path:     "/",
		Domain:   ctx.Request().Host,
		Secure:   false,
		HttpOnly: true,
	})
	res.SetData(http.StatusOK, nil)
	return ctx.JSON(res.GetStatus(), res.GetBody())
}
