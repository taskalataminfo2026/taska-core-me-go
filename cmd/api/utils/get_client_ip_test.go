package utils_test

//
//import (
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//
//	"taska-core-me-go/cmd/api/utils"
//)
//
//func TestGetClientIP(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//
//	t.Run("Debe obtener IP desde X-Forwarded-For", func(t *testing.T) {
//		c, _ := gin.CreateTestContext(httptest.NewRecorder())
//		req := httptest.NewRequest(http.MethodGet, "/", nil)
//		req.Header.Set("X-Forwarded-For", "203.0.113.5, 70.41.3.18, 150.172.238.178")
//		c.Request = req
//
//		ip := utils.GetClientIP(c)
//		assert.Equal(t, "203.0.113.5", ip)
//	})
//
//	t.Run("Debe obtener IP desde X-Real-IP", func(t *testing.T) {
//		c, _ := gin.CreateTestContext(httptest.NewRecorder())
//		req := httptest.NewRequest(http.MethodGet, "/", nil)
//		req.Header.Set("X-Real-IP", "198.51.100.23")
//		c.Request = req
//
//		ip := utils.GetClientIP(c)
//		assert.Equal(t, "198.51.100.23", ip)
//	})
//
//	t.Run("Debe obtener IP desde RemoteAddr cuando no hay headers", func(t *testing.T) {
//		c, _ := gin.CreateTestContext(httptest.NewRecorder())
//		req := httptest.NewRequest(http.MethodGet, "/", nil)
//		req.RemoteAddr = "192.0.2.44:12345"
//		c.Request = req
//
//		ip := utils.GetClientIP(c)
//		assert.Equal(t, "192.0.2.44", ip)
//	})
//
//	t.Run("Debe devolver vacío si no hay IP válida", func(t *testing.T) {
//		c, _ := gin.CreateTestContext(httptest.NewRecorder())
//		req := httptest.NewRequest(http.MethodGet, "/", nil)
//		req.RemoteAddr = ""
//		c.Request = req
//
//		ip := utils.GetClientIP(c)
//		assert.Equal(t, "", ip)
//	})
//}
