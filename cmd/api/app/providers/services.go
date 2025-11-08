package providers

import (
	"github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/services"
)

func JwtService() *services.JwtServices {
	return &services.JwtServices{}
}
