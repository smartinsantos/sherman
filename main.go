package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"sherman/src/app/config"
	"sherman/src/app/registry"
	"sherman/src/app/router"
	"time"
)

func main() {
	cfg := config.Get()
	if cfg.App.Debug {
		// pretty logger
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Info().Msg("Server Running on DEBUG mode")
	}

	diContainer, err := registry.Get()
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}
	r := router.New(diContainer)


	// Start server
	go func() {
		log.Info().Msg(fmt.Sprintf("Server Running on PORT%s", cfg.App.Addr))
		if err := r.Start(cfg.App.Addr); err != nil {
			log.Info().Msg("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with  a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.Shutdown(ctx); err != nil {
		log.Fatal().Err(err)
	}
}
