package config

import (
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ZeroLogConfig defines the config for ZeroLog middleware.
type ZeroLogConfig struct {
	// FieldMap set a list of fields with tags
	//
	// Tags to constructed the .go fields.
	//
	// - @id (Request ID)
	// - @remote_ip
	// - @uri
	// - @host
	// - @method
	// - @path
	// - @protocol
	// - @referer
	// - @user_agent
	// - @status
	// - @error
	// - @latency (In nanoseconds)
	// - @latency_human (Human readable)
	// - @bytes_in (Bytes received)
	// - @bytes_out (Bytes sent)
	// - @header:<NAME>
	// - @query:<NAME>
	// - @form:<NAME>
	// - @cookie:<NAME>
	FieldMap map[string]string
	// Logger it is a zerolog logger
	Logger zerolog.Logger
	// Skipper defines a function to skip middleware.
	Skipper emw.Skipper
}

// DefaultZeroLogConfig is the default ZeroLog middleware config.
var DefaultZeroLogConfig = ZeroLogConfig{
	FieldMap: map[string]string{
		"remote_ip": "@remote_ip",
		"uri":       "@uri",
		"host":      "@host",
		"method":    "@method",
		"status":    "@status",
		"latency":   "@latency",
		"error":     "@error",
	},
	Logger: log.Logger,
	Skipper: emw.DefaultSkipper,
}
