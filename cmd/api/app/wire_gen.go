package app

import (
	"github.com/labstack/echo/v4"
	"taska-core-me-go/cmd/api/app/providers"
)

func Start() (*echo.Echo, error) {
	db, err := providers.DatabaseConnectionPostgres()
	if err != nil {
		return nil, err
	}
	skillsRepository := providers.SkillsRepository(db)
	skillsServices := providers.SkillsServices(skillsRepository)
	validator := providers.Validator()
	skillsController := providers.SkillsController(skillsServices, validator)
	echoEcho := providers.ProviderRouter(skillsController)
	return echoEcho, nil
}
