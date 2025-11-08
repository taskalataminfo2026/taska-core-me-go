package services_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/config"
	"github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/constants"
	"github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/models"
	"github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/services"
)

func TestJwtServices(t *testing.T) {
	ctx := context.Background()
	s := &services.JwtServices{}

	config.Config.JWT.SecretKey = "test_secret"
	config.Config.JWT.AccessExpiry = time.Minute * 5
	config.Config.JWT.RefreshExpiry = time.Hour * 24

	t.Run("GenerateAccessToken_Success", func(t *testing.T) {
		token, err := s.GenerateAccessToken(ctx, 1, "wilson", "wilson@test.com", constants.TypeUserUser)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("GenerateRefreshToken_Success", func(t *testing.T) {
		token, err := s.GenerateRefreshToken(ctx, 2, "leonel", "leonel@test.com", constants.TypeUserUser)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("ValidateToken_Success", func(t *testing.T) {
		token, _ := s.GenerateAccessToken(ctx, 10, "testuser", "user@test.com", constants.TypeUserUser)
		claims, err := s.ValidateToken(ctx, token)
		assert.NoError(t, err)
		assert.Equal(t, int64(10), claims.UserID)
		assert.Equal(t, "testuser", claims.UserName)
	})

	t.Run("ValidateToken_InvalidSignatureMethod", func(t *testing.T) {
		badToken := jwt.NewWithClaims(jwt.SigningMethodRS256, &models.CustomClaims{
			UserID:   1,
			UserName: "baduser",
			Email:    "bad@test.com",
		})

		signed, _ := badToken.SignedString([]byte("wrong_key"))

		_, err := s.ValidateToken(ctx, signed)

		assert.Error(t, err)

		errMsg := err.Error()
		if !strings.Contains(errMsg, constants.ErrMsgUnexpectedSigningMethod) &&
			!strings.Contains(errMsg, constants.InvalidCredentialsMessage) {
			t.Errorf("Error inesperado, se esperaba que contuviera '%s' o '%s', pero fue: %v",
				constants.ErrMsgUnexpectedSigningMethod,
				constants.InvalidCredentialsMessage,
				errMsg)
		}
	})

	t.Run("ValidateToken_Expired", func(t *testing.T) {
		expiredClaims := &models.CustomClaims{
			UserID:   1,
			UserName: "expireduser",
			Email:    "expired@test.com",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Minute)),
			},
		}
		expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
		signed, _ := expiredToken.SignedString([]byte(config.Config.JWT.SecretKey))
		_, err := s.ValidateToken(ctx, signed)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), constants.SessionExpiredMessage)
	})

	t.Run("RefreshToken_Success", func(t *testing.T) {
		refresh, _ := s.GenerateRefreshToken(ctx, 5, "renewuser", "renew@test.com", constants.TypeUserUser)
		newAccess, newRefresh, err := s.RefreshToken(ctx, refresh)
		assert.NoError(t, err)
		assert.NotEmpty(t, newAccess)
		assert.NotEmpty(t, newRefresh)
	})

	t.Run("RefreshToken_Invalid", func(t *testing.T) {
		invalidToken := "invalid.token.value"
		_, _, err := s.RefreshToken(ctx, invalidToken)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), constants.ErrMsgInvalidRefreshToken)
	})

	t.Run("GenerateToken_DifferentExpiry", func(t *testing.T) {
		access, err1 := s.GenerateAccessToken(ctx, 1, "access", "a@test.com", constants.TypeUserUser)
		refresh, err2 := s.GenerateRefreshToken(ctx, 1, "refresh", "r@test.com", constants.TypeUserUser)

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.NotEqual(t, access, refresh)
	})
}
