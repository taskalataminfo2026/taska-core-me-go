package services

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"taska-core-me-go/cmd/api/config"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
)

//go:generate mockgen -destination=../mocks/services/$GOFILE -package=mservices -source=./$GOFILE

type IJWTServices interface {
	GenerateAccessToken(ctx context.Context, userID int64, username, email, role string) (string, error)
	GenerateRefreshToken(ctx context.Context, userID int64, username, email, role string) (string, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, string, error)
	ValidateToken(ctx context.Context, tokenString string) (*models.CustomClaims, error)
}

type JwtServices struct{}

func (s *JwtServices) GenerateAccessToken(ctx context.Context, userID int64, username, email, role string) (string, error) {
	accessToken, err := s.GenerateToken(userID, username, email, role, false)
	if err != nil {
		logger.StandardError(ctx, constants.LayerService, constants.ModuleJwt, constants.FunctionGenerateAccessToken, err, "Error al generar access token",
			zap.Int64("userID", userID),
			zap.String("username", username))
		return "", fmt.Errorf("%s: %w", constants.ErrMsgAccessTokenGenFailed, err)
	}
	return accessToken, nil
}

func (s *JwtServices) GenerateRefreshToken(ctx context.Context, userID int64, username, email, role string) (string, error) {
	refreshToken, err := s.GenerateToken(userID, username, email, role, true)
	if err != nil {
		logger.StandardError(ctx, constants.LayerService, constants.ModuleJwt, constants.FunctionGenerateRefreshToken, err, "Error al generar refresh token",
			zap.Int64("userID", userID),
			zap.String("username", username))
		return "", fmt.Errorf("%s: %w", constants.ErrMsgRefreshTokenGenFailed, err)
	}
	return refreshToken, nil
}

func (s *JwtServices) ValidateToken(ctx context.Context, tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&models.CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.StandardError(ctx, constants.LayerService, constants.ModuleJwt, constants.FunctionValidateToken, nil, "Método de firma inesperado en token")
				return nil, response_capture.NewErrorME(ctx, http.StatusInternalServerError, nil, constants.ErrMsgUnexpectedSigningMethod)
			}
			return []byte(config.Config.JWT.SecretKey), nil
		},
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, response_capture.NewErrorME(ctx, http.StatusUnauthorized, err, constants.SessionExpiredMessage)
		}

		return nil, response_capture.NewErrorME(ctx, http.StatusUnauthorized, err, constants.InvalidCredentialsMessage)
	}

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, response_capture.NewErrorME(ctx, http.StatusInternalServerError, nil, constants.InvalidCredentialsMessage)
}

func (s *JwtServices) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	claims, err := s.ValidateToken(ctx, refreshToken)
	if err != nil {
		logger.StandardError(ctx, constants.LayerService, constants.ModuleJwt, constants.FunctionRefreshToken, err, "Error al validar el refresh token",
			zap.String("refresh_token", refreshToken))
		return "", "", fmt.Errorf("%s: %w", constants.ErrMsgInvalidRefreshToken, err)
	}

	accessToken, err := s.GenerateAccessToken(ctx, claims.UserID, claims.UserName, claims.Email, claims.Role)
	if err != nil {
		logger.StandardError(ctx, constants.LayerService, constants.ModuleJwt, constants.FunctionRefreshToken, err, "Error al generar el nuevo access token",
			zap.Int64("user_id", claims.UserID),
			zap.String("email", claims.Email))
		return "", "", err
	}

	newRefreshToken, err := s.GenerateRefreshToken(ctx, claims.UserID, claims.UserName, claims.Email, claims.Role)
	if err != nil {
		logger.StandardError(ctx, constants.LayerService, constants.ModuleJwt, constants.FunctionRefreshToken, err, "Error al generar el nuevo refresh token",
			zap.Int64("user_id", claims.UserID),
			zap.String("email", claims.Email))
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func (s *JwtServices) GenerateToken(userID int64, username, email, role string, isRefresh bool) (string, error) {
	expiryTime := config.Config.JWT.AccessExpiry
	if isRefresh {
		expiryTime = config.Config.JWT.RefreshExpiry
	}

	claims := &models.CustomClaims{
		UserID:   userID,
		UserName: username,
		Email:    email,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiryTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "taska-auth-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.JWT.SecretKey))
}
