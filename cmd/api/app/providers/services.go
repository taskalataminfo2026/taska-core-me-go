package providers

import (
	"taska-core-me-go/cmd/api/repositories"
	"taska-core-me-go/cmd/api/services"
)

func JwtService() *services.JwtServices {
	return &services.JwtServices{}
}

func SkillsServices(sr repositories.ISkillsRepository) *services.SkillsServices {
	return &services.SkillsServices{SkillsRepository: sr}
}
