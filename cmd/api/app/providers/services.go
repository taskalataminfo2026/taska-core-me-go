package providers

import (
	"taska-core-me-go/cmd/api/services"
)

func JwtService() *services.JwtServices {
	return &services.JwtServices{}
}
