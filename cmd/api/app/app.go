//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"taska-core-me-go/cmd/api/app/providers"
	"taska-core-me-go/cmd/api/clients/rusty"
	"taska-core-me-go/cmd/api/controllers"
	"taska-core-me-go/cmd/api/repositories"
	"taska-core-me-go/cmd/api/services"
	"taska-core-me-go/cmd/api/validator"
)

// 🔹 Database
var DatabaseSet = wire.NewSet(
	providers.DatabaseConnectionPostgres,
)

// 🔹 Clients
var ClientSet = wire.NewSet(
	providers.GetRustyClient,
	wire.Bind(new(rusty.IRustyClient), new(*rusty.RustyClient)),
)

// 🔹 Controllers
var ControllerRouterSet = wire.NewSet(
	providers.CategoriesController,
	wire.Bind(new(controllers.ICategoriesController), new(*controllers.CategoriesController)),
	providers.SkillsController,
	wire.Bind(new(controllers.ISkillsController), new(*controllers.SkillsController)),
	providers.TaskerController,
	wire.Bind(new(controllers.ITaskerController), new(*controllers.TaskerController)),
)

// 🔹 Services
var ServicesRouterSet = wire.NewSet(
	providers.JwtService,
	wire.Bind(new(services.IJWTServices), new(*services.JwtServices)),
	providers.CategoriesServices,
	wire.Bind(new(services.ICategoriesServices), new(*services.CategoriesServices)),
	providers.SkillsServices,
	wire.Bind(new(services.ISkillsServices), new(*services.SkillsServices)),
	providers.TaskerServices,
	wire.Bind(new(services.ITaskerServices), new(*services.TaskerServices)),
)

// 🔹 Repositories
var RepositoryRouterSet = wire.NewSet(
	providers.CategoriesRepository,
	wire.Bind(new(repositories.ICategoriesRepository), new(*repositories.CategoriesRepository)),
	providers.RolesTokenRepository,
	wire.Bind(new(repositories.IRolesRepository), new(*repositories.RolesRepository)),
	providers.BlacklistedTokenRepository,
	wire.Bind(new(repositories.IBlacklistedTokenRepository), new(*repositories.BlacklistedTokenRepository)),
	providers.SkillsRepository,
	wire.Bind(new(repositories.ISkillsRepository), new(*repositories.SkillsRepository)),
)

// 🔹 Validators
var ValidatorRouterSet = wire.NewSet(
	providers.Validator,
	wire.Bind(new(validator.IValidator), new(*validator.Validator)),
)

// 🔹 Router
var RouterSet = wire.NewSet(
	ControllerRouterSet,
	ServicesRouterSet,
	RepositoryRouterSet,
	//GatewayRouterSet,
	ValidatorRouterSet,
	providers.ProviderRouter,
)

// 🔹 Start app
func Start() (*echo.Echo, error) {
	panic(wire.Build(
		DatabaseSet,
		RouterSet,
	))
	return nil, nil
}
