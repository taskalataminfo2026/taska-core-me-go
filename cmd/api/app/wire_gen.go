package app

import (
	"github.com/labstack/echo/v4"
	"taska-core-me-go/cmd/api/app/providers"
)

func Start() (*echo.Echo, error) {
	echoEcho := providers.ProviderRouter()
	return echoEcho, nil
}
