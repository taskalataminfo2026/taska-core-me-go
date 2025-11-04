package utils

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"taska-core-me-go/cmd/api/utils/detached"
)

type contextKey string

var contextKeyJWT = contextKey("jwtToken")

func CreateRequestContext(c echo.Context) context.Context {
	if c == nil {
		return context.Background()
	}

	ctx := c.Request().Context()
	newCtx := NewContextFlow(ctx)

	jwtToken := GetJWTFromHeader(c.Request().Header)
	if jwtToken != "" {
		newCtx = context.WithValue(newCtx, contextKeyJWT, jwtToken)
	}

	return newCtx
}

func GetJWTFromHeader(headers http.Header) string {
	authHeader := headers.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return ""
}

func NewContextFlow(ctx context.Context) context.Context {
	newCtx := detached.Detach(ctx)
	defer newCtx.Done()

	return newCtx
}

// Test.

func CreateRequest(req *http.Request) context.Context {
	if req == nil {
		return context.Background()
	}
	ctx := NewContextFlow(req.Context())
	return ctx
}

func GetTestRequestWithHeaders() *http.Request {
	req, _ := http.NewRequest("test", "test", nil)
	req.Header.Set("x-request-id", "request-id")
	return req
}

// MergeHeaders combina los headers por defecto con otros personalizados.
func MergeHeaders(defaults, custom map[string]string) map[string]string {
	for k, v := range custom {
		defaults[k] = v
	}
	return defaults
}

func DefaultHeaders() map[string]string {
	return map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
}
