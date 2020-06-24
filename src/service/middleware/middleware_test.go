package middleware

import (
	"bytes"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sherman/mocks"
	_ "sherman/src/app/testing"
	"sherman/src/domain/auth"
	"sherman/src/service/config"
	cmc "sherman/src/service/middleware/config"
	"strings"
	"testing"
)

type middlewareMockDeps struct {
	config          config.GlobalConfig
	securityService *mocks.Security
}

func genMockMiddleware() (Middleware, middlewareMockDeps) {
	mDeps := middlewareMockDeps{
		config:          config.TestConfig,
		securityService: new(mocks.Security),
	}
	m := New(mDeps.config, mDeps.securityService)
	return m, mDeps
}

func TestJWT(t *testing.T) {
	t.Run("request should go thru", func(t *testing.T) {
		m, mDeps := genMockMiddleware()
		mDeps.securityService.
			On("GetAndValidateAccessToken", mock.Anything).
			Return(auth.TokenMetadata{}, nil)

		e := echo.New()
		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}
		h := m.JWT()(handler)
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		if assert.NoError(t, h(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "test", rec.Body.String())
		}
	})

	t.Run("request should not go thru", func(t *testing.T) {
		m, mDeps := genMockMiddleware()
		mDeps.securityService.
			On("GetAndValidateAccessToken", mock.Anything).
			Return(auth.TokenMetadata{}, errors.New("some error"))

		e := echo.New()
		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}
		h := m.JWT()(handler)
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		if assert.NoError(t, h(ctx)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
			assert.Equal(t, "{\"data\":null,\"error\":\"invalid token\"}\n", rec.Body.String())
		}
	})
}

func TestZeroLog(t *testing.T) {
	t.Run("ZeroLog with default config", func(t *testing.T) {
		m, _ := genMockMiddleware()
		e := echo.New()
		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}
		h := m.ZeroLog()(handler)
		req := httptest.NewRequest(echo.GET, "/some", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		if assert.NoError(t, h(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("ZeroLogWithConfig custom config", func(t *testing.T) {
		m, _ := genMockMiddleware()
		e := echo.New()

		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}
		b := new(bytes.Buffer)
		logger := log.Output(zerolog.ConsoleWriter{Out: b, NoColor: true})
		fields := cmc.DefaultZeroLogConfig.FieldMap
		fields["empty"] = ""
		fields["id"] = "@id"
		fields["path"] = "@path"
		fields["protocol"] = "@protocol"
		fields["referer"] = "@referer"
		fields["user_agent"] = "@user_agent"
		fields["store"] = "@header:store"
		fields["filter_name"] = "@query:name"
		fields["username"] = "@form:username"
		fields["session"] = "@cookie:session"
		fields["latency_human"] = "@latency_human"
		fields["bytes_in"] = "@bytes_in"
		fields["bytes_out"] = "@bytes_out"
		fields["referer"] = "@referer"
		fields["user"] = "@header:user"
		config := cmc.ZeroLogConfig{
			Logger:   logger,
			FieldMap: fields,
		}
		h := m.ZeroLogWithConfig(&config)(handler)

		form := url.Values{}
		form.Add("username", "doejohn")

		req := httptest.NewRequest(echo.POST, "http://some?name=john", strings.NewReader(form.Encode()))
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationForm)
		req.Header.Add("Referer", "http://foo.bar")
		req.Header.Add("User-Agent", "cli-agent")
		req.Header.Add(echo.HeaderXForwardedFor, "http://foo.bar")
		req.Header.Add("user", "admin")
		req.AddCookie(&http.Cookie{
			Name:  "session",
			Value: "A1B2C3",
		})

		rec := httptest.NewRecorder()
		rec.Header().Add(echo.HeaderXRequestID, "123")

		ctx := e.NewContext(req, rec)

		if assert.NoError(t, h(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			res := b.String()

			tests := []struct {
				str string
				err string
			}{
				{"handle request", "invalid log: handle request info not found"},
				{"id=123", "invalid log: request id not found"},
				{`remote_ip=http://foo.bar`, "invalid log: remote ip not found"},
				{`uri=http://some?name=john`, "invalid log: uri not found"},
				{"host=some", "invalid log: host not found"},
				{"method=POST", "invalid log: method not found"},
				{"status=200", "invalid log: status not found"},
				{"latency=", "invalid log: latency not found"},
				{"latency_human=", "invalid log: latency_human not found"},
				{"bytes_in=0", "invalid log: bytes_in not found"},
				{"bytes_out=4", "invalid log: bytes_out not found"},
				{"path=/", "invalid log: path not found"},
				{"protocol=HTTP/1.1", "invalid log: protocol not found"},
				{`referer=http://foo.bar`, "invalid log: referer not found"},
				{"user_agent=cli-agent", "invalid log: user_agent not found"},
				{"user=admin", "invalid log: header user not found"},
				{"filter_name=john", "invalid log: query filter_name not found"},
				{"username=doejohn", "invalid log: form field username not found"},
				{"session=A1B2C3", "invalid log: cookie session not found"},
			}

			for _, test := range tests {
				if !strings.Contains(res, test.str) {
					t.Error(test.err)
				}
			}
		}
	})

	t.Run("ZeroLogWithConfig empty config", func(t *testing.T) {
		m, _ := genMockMiddleware()
		e := echo.New()
		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}
		h := m.ZeroLogWithConfig(&cmc.ZeroLogConfig{})(handler)
		req := httptest.NewRequest(echo.GET, "/some", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		if assert.NoError(t, h(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("ZeroLogWithConfig with skipper", func(t *testing.T) {
		m, _ := genMockMiddleware()
		e := echo.New()
		config := cmc.DefaultZeroLogConfig
		config.Skipper = func(c echo.Context) bool {
			return true
		}
		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}
		h := m.ZeroLogWithConfig(&config)(handler)
		req := httptest.NewRequest(echo.GET, "/some", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		if assert.NoError(t, h(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("ZeroLogWithConfig retrieves an error", func(t *testing.T) {
		m, _ := genMockMiddleware()
		e := echo.New()
		b := new(bytes.Buffer)
		logger := log.Output(zerolog.ConsoleWriter{Out: b, NoColor: true})
		config := cmc.ZeroLogConfig{
			Logger: logger,
		}
		handler := func(c echo.Context) error {
			return errors.New("error")
		}
		h := m.ZeroLogWithConfig(&config)(handler)
		req := httptest.NewRequest(echo.GET, "/some", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.Error(t, h(ctx)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			res := b.String()

			if !strings.Contains(res, "status=500") {
				t.Errorf("invalid log: wrong status code")
			}

			if !strings.Contains(res, `error=error`) {
				t.Errorf("invalid log: error not found")
			}
		}
	})
}
