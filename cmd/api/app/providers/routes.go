package providers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	middlewares_lib "github.com/taskalataminfo2026/tool-kit-lib-go/pkg/middlewares"
	"go.uber.org/zap"
	"net/http"
	"os"
	"taska-core-me-go/cmd/api/controllers"
)

func ProviderRouter(
	skillsController *controllers.SkillsController,
	taskerController *controllers.TaskerController,
) *echo.Echo {
	router := echo.New()
	logger, _ := zap.NewProduction()

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	// Swagger.
	router.GET("/swagger/*", func(c echo.Context) error {
		c.Response().Header().Set(
			"Content-Security-Policy",
			"default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'",
		)
		return echoSwagger.WrapHandler(c)
	})
	router.GET("", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})

	frontendEnv := os.Getenv("FRONTEND_URL")
	middlewares_lib.RegisterBaseMiddlewares(router, logger, frontendEnv)

	router.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "OK",
			"service": "taska-core-me-go",
			"version": "1.0.0",
		})
	})

	core := router.Group("/v1/api/core")
	{
		// Verificación de cuenta (Skills).
		skills := core.Group("/skills")
		skills.GET("/search", skillsController.SkillsSearch)
		skills.GET("/List", skillsController.SkillsList)

		// Verificación de cuenta (Tasker).
		tasker := core.Group("/tasker")
		tasker.GET("/{id_user}/skills", taskerController.TaskerProfile)
	}

	return router
}
