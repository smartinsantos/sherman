package main

import (
	"github.com/joho/godotenv"
	"github.com/smartinsantos/go-auth-api/config"
	"github.com/smartinsantos/go-auth-api/infrastructure/datastore"
	"github.com/smartinsantos/go-auth-api/infrastructure/router"
	"github.com/smartinsantos/go-auth-api/interfaces/handler"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error: No .env file found")
	}
}

func main() {
	env := config.Get()
	// Application DataStore
	appDataStore, err := datastore.New()
	if err != nil {
		panic(err)
	}
	defer appDataStore.Close()
	// Application RequestHandlers
	appHandler := handler.New(appDataStore)
	// Application Router
	appRouter := router.New(appHandler)
	log.Fatal(appRouter.Run(env.AppConfig.Addr))
}