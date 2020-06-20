package testing

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {
	rd := os.Getenv("ROOT_DIR")
	if rd == "" {
		log.Fatal().Msg("ROOT_DIR env variable is required to run tests, use ROOT_DIR=[project_root] go test [test_directory]")
		return
	}

	err := os.Chdir(rd)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("error: can not change directory to %s", rd))
	}
}
