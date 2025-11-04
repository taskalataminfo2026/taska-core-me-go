package utils

import (
	"github.com/labstack/echo/v4"
	"net"
	"strings"
)

func GetClientIP(c echo.Context) string {
	req := c.Request()
	// 1) X-Forwarded-For (puede contener lista: "client, proxy1, proxy2")
	if xff := strings.TrimSpace(req.Header.Get("X-Forwarded-For")); xff != "" {
		// tomar la primera entrada
		parts := strings.Split(xff, ",")
		if len(parts) > 0 {
			ip := strings.TrimSpace(parts[0])
			if parsed := net.ParseIP(ip); parsed != nil {
				return ip
			}
		}
	}

	// 2) X-Real-IP
	if xr := strings.TrimSpace(req.Header.Get("X-Real-Ip")); xr != "" {
		if parsed := net.ParseIP(xr); parsed != nil {
			return xr
		}
	}

	// 3) RemoteAddr
	host, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr))
	if err == nil {
		if parsed := net.ParseIP(host); parsed != nil {
			return host
		}
	}

	// Si no se pudo obtener, devolver vacío
	return ""
}

// Optional: NormalizeIP corta la IP a 45 chars y la limpia
func NormalizeIP(ip string) string {
	if ip == "" {
		return ""
	}
	if len(ip) > 45 {
		return ip[:45]
	}
	return ip
}
