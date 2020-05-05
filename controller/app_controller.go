package controller

import (
	"github.com/smartinsantos/go-auth-api/infrastructure/datastore"
)

// AppController wraps all applications handlers
type AppController struct {
	User UserController
}

// AppController constructor
func New(ads *datastore.AppDataStore) *AppController {
	ac := AppController{
		User: UserController{ ds: &ads.User },
	}
	return &ac
}