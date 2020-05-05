package main

import (
	"github.com/joho/godotenv"
	"github.com/smartinsantos/go-auth-api/config"
	"github.com/smartinsantos/go-auth-api/controller"
	"github.com/smartinsantos/go-auth-api/infrastructure/datastore"
	"github.com/smartinsantos/go-auth-api/router"
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
	appController := controller.New(appDataStore)
	// Application Router
	appRouter := router.New(appController)
	log.Fatal(appRouter.Run(env.AppConfig.Addr))
}