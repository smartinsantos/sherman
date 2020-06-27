package testing

import (
	"fmt"
	// makes mysql driver available on test environment
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {
	rd := os.Getenv("ROOT_DIR")
	if rd == "" {
		log.Fatal().Msg("ROOT_DIR env variable is required to run tests, use ROOT_DIR=[project_root] go test [test_directory]")
		return
	}
	if err := os.Chdir(rd); err != nil {
		log.Fatal().Msg(fmt.Sprintf("error: can not change directory to %s", rd))
	}
}
