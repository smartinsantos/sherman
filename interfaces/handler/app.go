package handler

import (
	"github.com/smartinsantos/go-auth-api/infrastructure/datastore"
)

type AppHandler struct {
	User UserHandler
}

// AppHandler constructor
func New(ads *datastore.AppDataStore) *AppHandler {
	ac := AppHandler{
		User: UserHandler{ ds: &ads.User },
	}
	return &ac
}