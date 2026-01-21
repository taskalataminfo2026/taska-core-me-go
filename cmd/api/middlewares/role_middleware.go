package middlewares

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
)

var (
	roleHierarchyCache map[string]int
	cacheMutex         sync.RWMutex
	Conn               *gorm.DB
)

func InitRoleMiddleware(conn *gorm.DB) {
	Conn = conn
}

func RequireRole(requiredRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			if Conn == nil {
				return response_capture.RespondError(c,
					response_capture.NewErrorME(ctx, http.StatusInternalServerError, nil, constants.ErrDatabaseNotInitialized))
			}

			claims, ok := c.Get(constants.ContextKeyClaims).(*models.CustomClaims)
			if !ok || claims == nil {
				return response_capture.RespondError(c,
					response_capture.NewErrorME(ctx, http.StatusUnauthorized, nil, constants.ErrTokenInvalid))
			}

			token, ok := c.Get(constants.ContextKeyToken).(string)
			if !ok || token == "" {
				return response_capture.RespondError(c,
					response_capture.NewErrorME(ctx, http.StatusUnauthorized, nil, constants.ErrTokenMissing))
			}

			blackRepo := repositories.BlacklistedTokenRepository{Conn: Conn}
			blacklisted, err := blackRepo.FirstByTokenNil(ctx, token)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return response_capture.RespondError(c,
					response_capture.NewErrorME(ctx, http.StatusInternalServerError, err, constants.ErrTokenBlacklistValidation))
			}
			if blacklisted.ID > 0 {
				return response_capture.RespondError(c,
					response_capture.NewErrorME(ctx, http.StatusUnauthorized, nil, constants.ErrTokenRevoked))
			}

			roleHierarchy, err := GetRoleHierarchy(ctx)
			if err != nil {
				return response_capture.RespondError(c,
					response_capture.NewErrorME(ctx, http.StatusInternalServerError, err, constants.ErrRoleHierarchyLoad))
			}

			userRole := claims.Role
			for _, requiredRole := range requiredRoles {
				if hasRequiredRole(userRole, requiredRole, roleHierarchy) {
					return next(c)
				}
			}

			return response_capture.RespondError(c,
				response_capture.NewErrorME(ctx, http.StatusForbidden, nil, constants.ErrPermissionDenied))
		}
	}
}

func GetRoleHierarchy(ctx context.Context) (map[string]int, error) {
	cacheMutex.RLock()
	if roleHierarchyCache != nil {
		defer cacheMutex.RUnlock()
		return roleHierarchyCache, nil
	}
	cacheMutex.RUnlock()

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if roleHierarchyCache != nil {
		return roleHierarchyCache, nil
	}

	roleRepo := repositories.RolesRepository{Conn: Conn}
	roles, err := roleRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	roleHierarchy := make(map[string]int)
	for _, role := range roles {
		roleHierarchy[role.Name] = role.Level
	}

	roleHierarchyCache = roleHierarchy
	return roleHierarchyCache, nil
}

func hasRequiredRole(userRole, requiredRole string, roleHierarchy map[string]int) bool {
	userLevel, okUser := roleHierarchy[userRole]
	requiredLevel, okReq := roleHierarchy[requiredRole]

	if !okUser || !okReq {
		return false
	}

	return userLevel >= requiredLevel
}
