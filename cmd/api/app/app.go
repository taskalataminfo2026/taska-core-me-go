//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/repositories"
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

// 🔹 Services
var ServicesRouterSet = wire.NewSet(
	providers.JwtService,
	wire.Bind(new(services.IJWTServices), new(*services.JwtServices)),
)

// 🔹 Repositories
var RepositoryRouterSet = wire.NewSet(
	providers.RolesTokenRepository,
	wire.Bind(new(repositories.IRolesRepository), new(*repositories.RolesRepository)),
)

// 🔹 Validators
var ValidatorRouterSet = wire.NewSet(
	providers.Validator,
	wire.Bind(new(validator.IValidator), new(*validator.Validator)),
)

// 🔹 Router
var RouterSet = wire.NewSet(
	//ControllerRouterSet,
	ServicesRouterSet,
	RepositoryRouterSet,
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
