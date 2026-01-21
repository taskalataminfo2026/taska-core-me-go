package providers

import (
	"taska-core-me-go/cmd/api/controllers"
	"taska-core-me-go/cmd/api/services"
	"taska-core-me-go/cmd/api/validator"
)

func SkillsController(ss services.ISkillsServices, validator validator.IValidator) *controllers.SkillsController {
	return &controllers.SkillsController{SkillsServices: ss, Validator: validator}
}
