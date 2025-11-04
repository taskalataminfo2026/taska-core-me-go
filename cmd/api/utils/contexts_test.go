package utils

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ---------- Tests CreateRequestContext ----------
func TestCreateRequestContext_NilContext(t *testing.T) {
	ctx := CreateRequestContext(nil)
	if ctx == nil {
		t.Errorf("expected non-nil context, got nil")
	}
}

// ---------- Tests GetJWTFromHeader ----------
func TestGetJWTFromHeader(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		expected   string
	}{
		{"No header", "", ""},
		{"Invalid header", "Token abc123", ""},
		{"Valid Bearer token", "Bearer mytoken123", "mytoken123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			result := GetJWTFromHeader(headers)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestCreateRequestContext_WithNoAuthHeader(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctx := CreateRequestContext(c)

	if val := ctx.Value(contextKeyJWT); val != nil {
		t.Errorf("expected nil jwt token, got %v", val)
	}
}

func TestCreateRequestContext_WithAuthHeader(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer secret123")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctx := CreateRequestContext(c)

	if val := ctx.Value(contextKeyJWT); val != "secret123" {
		t.Errorf("expected jwt token %q, got %v", "secret123", val)
	}
}

// ---------- Tests NewContextFlow ----------
func TestNewContextFlow(t *testing.T) {
	parent := context.Background()
	newCtx := NewContextFlow(parent)

	if newCtx == nil {
		t.Errorf("expected non-nil context")
	}

	if newCtx == parent {
		t.Errorf("expected a detached context, got same as parent")
	}
}
