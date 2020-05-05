package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smartinsantos/go-auth-api/domain/entity"
	"github.com/smartinsantos/go-auth-api/infrastructure/datastore"
	"net/http"
)

// UserHandler struct defines the dependencies that will be used
type UserHandler struct {
	ds *datastore.UserDataStore
}

// Registers the user
func (uc * UserHandler) Register (context *gin.Context) {
	// TODO: make automated response object generators
	var user entity.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"message": "Invalid json",
		})
		return
	}

	// TODO: validate inputs
	newUser, err := uc.ds.CreateUser(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"message": "Unable to create user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"message": "User created",
		"data": gin.H{
			"id": newUser.ID,
			"email_address": newUser.EmailAddress,
			"first_name": newUser.FirstName,
			"last_name": newUser.LastName,
			"active": newUser.Active,
			"created_at": newUser.CreatedAt,
			"updated_at": newUser.UpdatedAt,
		},
	})
}

// Logs the user
func (uc * UserHandler) Login (context *gin.Context) {
	context.String(http.StatusOK, "Login")
}

// Refreshes user token
func (uc * UserHandler) RefreshToken (context *gin.Context) {
	context.String(http.StatusOK, "RefreshToken")
}

// Verify that the user token is still valid
func (uc * UserHandler) VerifyAuth (context *gin.Context) {
	context.String(http.StatusOK, "VerifyAuth")
}


