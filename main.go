package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sherman/src/app/config"
	"sherman/src/app/registry"
	"sherman/src/app/router"
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
	// run the server
	log.Info().Msg(fmt.Sprintf("Server Running on PORT%s", cfg.App.Addr))
	log.Fatal().Err(r.Start(cfg.App.Addr))
}
