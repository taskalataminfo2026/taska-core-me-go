package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
	"net/http"
	"strings"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/services"
)

func JWTMiddleware(jwtService services.IJWTServices) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			authHeader := c.Request().Header.Get(constants.HeaderAuthorization)

			if authHeader == "" {
				return response_capture.RespondError(c,
					response_capture.NewErrorME(ctx, http.StatusUnauthorized, nil, constants.ErrTokenMissing))
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return response_capture.RespondError(c,
					response_capture.NewErrorME(ctx, http.StatusUnauthorized, nil, constants.ErrTokenBadFormat))
			}

			token := strings.TrimSpace(parts[1])
			if token == "" {
				return response_capture.RespondError(c,
					response_capture.NewErrorME(ctx, http.StatusUnauthorized, nil, constants.ErrTokenEmpty))
			}

			claims, err := jwtService.ValidateToken(ctx, token)
			if err != nil {
				status := http.StatusUnauthorized
				message := constants.ErrTokenInvalid

				if strings.Contains(err.Error(), "expired") {
					message = constants.ErrTokenExpired
				} else if strings.Contains(err.Error(), "signature") {
					message = constants.ErrTokenSignatureInvalid
				}

				return response_capture.RespondError(c,
					response_capture.NewErrorME(ctx, status, err, message))
			}

			ctx = context.WithValue(ctx, logger.ContextKeyUserID, claims.UserID)
			ctx = context.WithValue(ctx, logger.ContextKeyUserRole, claims.Role)
			ctx = context.WithValue(ctx, logger.ContextKeyRequestID, c.Response().Header().Get(echo.HeaderXRequestID))

			c.Set(constants.ContextKeyClaims, claims)
			c.Set(constants.ContextKeyToken, token)

			return next(c)
		}
	}
}
