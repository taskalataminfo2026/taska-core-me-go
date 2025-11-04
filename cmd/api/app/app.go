package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"taska-core-me-go/cmd/api/app/providers"
	"taska-core-me-go/cmd/api/clients/rusty"
	"taska-core-me-go/cmd/api/validator"
)

// 🔹 Database
var DatabaseSet = wire.NewSet(
	providers.DatabaseConnectionPostgres(),
)

// 🔹 Clients
var ClientSet = wire.NewSet(
	providers.GetRustyClient,
	wire.Bind(new(rusty.IRustyClient), new(*rusty.RustyClient)),
)

// 🔹 Validators
var ValidatorRouterSet = wire.NewSet(
	providers.Validator,
	wire.Bind(new(validator.IValidator), new(*validator.Validator)),
)

// 🔹 Router
var RouterSet = wire.NewSet(
	//ControllerRouterSet,
	//ServicesRouterSet,
	//RepositoryRouterSet,
	//GatewayRouterSet,
	ValidatorRouterSet,
	providers.ProviderRouter,
)

func Start() (*echo.Echo, error) {
	panic(wire.Build(
		ClientSet,
		DatabaseSet,
		RouterSet,
	))
	return nil, nil
}
