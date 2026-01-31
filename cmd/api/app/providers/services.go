package providers

import (
	"taska-core-me-go/cmd/api/repositories"
	"taska-core-me-go/cmd/api/services"
)

func JwtService() *services.JwtServices {
	return &services.JwtServices{}
}

func CategoriesServices(cr repositories.ICategoriesRepository) *services.CategoriesServices {
	return &services.CategoriesServices{CategoriesRepository: cr}
}

func SkillsServices(sr repositories.ISkillsRepository) *services.SkillsServices {
	return &services.SkillsServices{SkillsRepository: sr}
}

func TaskerServices(sr repositories.ISkillsRepository) *services.TaskerServices {
	return &services.TaskerServices{SkillsRepository: sr}
}
