package mysqlds

import (
	"github.com/gchaincl/dotsql"
	"github.com/gobuffalo/packr/v2"
)

func loadSqlDot(fileName string) (*dotsql.DotSql, error) {
	var err error

	dsBox := packr.New("datastore", "./")

	repositorySqlStr, err := dsBox.FindString(fileName)
	if err != nil {
		return nil, err
	}

	dot, err := dotsql.LoadFromString(repositorySqlStr)
	if err != nil {
		return nil, err
	}

	return dot, nil
}