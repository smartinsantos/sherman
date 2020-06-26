package main

import (
	"github.com/rs/zerolog/log"
	"sherman/src/app"
	"sherman/src/app/registry"
)

func main() {
	diContainer, err := registry.GetAppContainer()
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	app.Run(diContainer)
}
