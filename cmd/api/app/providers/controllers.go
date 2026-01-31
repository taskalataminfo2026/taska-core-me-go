package providers

import (
	"taska-core-me-go/cmd/api/controllers"
	"taska-core-me-go/cmd/api/services"
	"taska-core-me-go/cmd/api/validator"
)

func CategoriesController(cs services.ICategoriesServices, validator validator.IValidator) *controllers.CategoriesController {
	return &controllers.CategoriesController{CategoriesServices: cs, Validator: validator}
}

func SkillsController(ss services.ISkillsServices, validator validator.IValidator) *controllers.SkillsController {
	return &controllers.SkillsController{SkillsServices: ss, Validator: validator}
}

func TaskerController(ts services.ITaskerServices, validator validator.IValidator) *controllers.TaskerController {
	return &controllers.TaskerController{TaskerServices: ts, Validator: validator}
}
