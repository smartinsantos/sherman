package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserController struct defines the dependencies that will be used
type UserController struct {}

// UserController constructor
func NewUserController() *UserController {
	uc := UserController{}
	return &uc
}

func (uc * UserController) Register (context *gin.Context) {
	context.String(http.StatusOK, "Register")
}

func (uc * UserController) Login (context *gin.Context) {
	context.String(http.StatusOK, "Login")
}

func (uc * UserController) RefreshToken (context *gin.Context) {
	context.String(http.StatusOK, "RefreshToken")
}

func (uc * UserController) VerifyAuth (context *gin.Context) {
	context.String(http.StatusOK, "VerifyAuth")
}


