package providers

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	middlewares_lib "github.com/taskalataminfo2026/tool-kit-lib-go/pkg/middlewares"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func ProviderRouter() *echo.Echo {
	router := echo.New()
	logger, _ := zap.NewProduction()

	// Swagger.
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	router.GET("", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})

	frontendEnv := os.Getenv("FRONTEND_URL")
	middlewares_lib.RegisterBaseMiddlewares(router, logger, frontendEnv)

	router.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "OK",
			"service": "taska-core-me-go",
			"version": "1.0.0"})
	})

	return router
}
