package middleware

import (
	"github.com/labstack/echo/v4"
	cmc "sherman/src/service/middleware/config"
	"strconv"
	"strings"
	"time"
)

// ZeroLog returns a middleware that logs HTTP requests.
func (s *service) ZeroLog() echo.MiddlewareFunc {
	return s.ZeroLogWithConfig(&cmc.DefaultZeroLogConfig)
}

// ZeroLogWithConfig returns a middleware that logs HTTP requests.
func (s *service) ZeroLogWithConfig(cfg *cmc.ZeroLogConfig) echo.MiddlewareFunc {
	// Defaults
	if cfg.Skipper == nil {
		cfg.Skipper = cmc.DefaultZeroLogConfig.Skipper
	}

	if len(cfg.FieldMap) == 0 {
		cfg.FieldMap = cmc.DefaultZeroLogConfig.FieldMap
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			if cfg.Skipper(ctx) {
				return next(ctx)
			}

			req := ctx.Request()
			res := ctx.Response()
			start := time.Now()

			if err = next(ctx); err != nil {
				ctx.Error(err)
			}

			stop := time.Now()
			entry := cfg.Logger.Info()

			for k, v := range cfg.FieldMap {
				if v == "" {
					continue
				}

				switch v {
				case "@id":
					id := req.Header.Get(echo.HeaderXRequestID)

					if id == "" {
						id = res.Header().Get(echo.HeaderXRequestID)
					}

					entry = entry.Str(k, id)
				case "@remote_ip":
					entry = entry.Str(k, ctx.RealIP())
				case "@uri":
					entry = entry.Str(k, req.RequestURI)
				case "@host":
					entry = entry.Str(k, req.Host)
				case "@method":
					entry = entry.Str(k, req.Method)
				case "@path":
					p := req.URL.Path

					if p == "" {
						p = "/"
					}

					entry = entry.Str(k, p)
				case "@protocol":
					entry = entry.Str(k, req.Proto)
				case "@referer":
					entry = entry.Str(k, req.Referer())
				case "@user_agent":
					entry = entry.Str(k, req.UserAgent())
				case "@status":
					entry = entry.Int(k, res.Status)
				case "@error":
					if err != nil {
						entry = entry.Err(err)
					}
				case "@latency":
					l := stop.Sub(start)
					entry = entry.Str(k, strconv.FormatInt(int64(l), 10))
				case "@latency_human":
					entry = entry.Str(k, stop.Sub(start).String())
				case "@bytes_in":
					cl := req.Header.Get(echo.HeaderContentLength)

					if cl == "" {
						cl = "0"
					}

					entry = entry.Str(k, cl)
				case "@bytes_out":
					entry = entry.Str(k, strconv.FormatInt(res.Size, 10))
				default:
					switch {
					case strings.HasPrefix(v, "@header:"):
						entry = entry.Str(k, ctx.Request().Header.Get(v[8:]))
					case strings.HasPrefix(v, "@query:"):
						entry = entry.Str(k, ctx.QueryParam(v[7:]))
					case strings.HasPrefix(v, "@form:"):
						entry = entry.Str(k, ctx.FormValue(v[6:]))
					case strings.HasPrefix(v, "@cookie:"):
						cookie, err := ctx.Cookie(v[8:])
						if err == nil {
							entry = entry.Str(k, cookie.Value)
						}
					}
				}
			}

			entry.Msg("handle request")

			return
		}
	}
}
