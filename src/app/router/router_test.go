package router

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"sherman/src/app/registry"
	_ "sherman/src/app/testing"
	"testing"
)

type Route struct {
	Method string
	Path   string
}

var expectedRoutes = []Route{
	{
		Method: "POST",
		Path:   "/api/v1/users/register",
	},
	{
		Method: "POST",
		Path:   "/api/v1/users/login",
	},
	{
		Method: "PATCH",
		Path:   "/api/v1/users/refresh-token",
	},
	{
		Method: "GET",
		Path:   "/api/v1/users/:id",
	},
	{
		Method: "DELETE",
		Path:   "/api/v1/users/logout",
	},
}

func containsRoute(routes []*echo.Route, method, path string) bool {
	for _, expectedRoute := range routes {
		if expectedRoute.Method == method && expectedRoute.Path == path {
			return true
		}
	}
	return false
}

func TestNew(t *testing.T) {
	diContainer, err := registry.Get()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	router := New(diContainer)
	assert.NotEmpty(t, router)
	
	routes := router.Routes()
	if !assert.Equal(t, len(expectedRoutes), len(routes)) {
		t.Fatal("the number of routes present defer from expected")
	}

	for _, eRoute := range expectedRoutes {
		assert.True(t, containsRoute(routes, eRoute.Method, eRoute.Path))
	}
}
