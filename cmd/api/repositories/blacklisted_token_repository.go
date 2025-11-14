package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/models"
	models2 "taska-core-me-go/cmd/api/repositories/models"
	"time"
)

//go:generate mockgen -destination=../mocks/repositories/$GOFILE -package=mrepositories -source=./$GOFILE

type IBlacklistedTokenRepository interface {
	FirstByToken(ctx context.Context, token string) (models.BlacklistedToken, error)
	FirstByTokenNil(ctx context.Context, token string) (models.BlacklistedToken, error)
	DeleteExpired(ctx context.Context, before time.Time) (int64, error)
	Save(ctx context.Context, body models.BlacklistedToken) (models.BlacklistedToken, error)
}

const blacklistedTableName = "blacklisted_tokens"

type BlacklistedTokenRepository struct {
	Conn *gorm.DB
}

func (repository *BlacklistedTokenRepository) FirstByToken(ctx context.Context, token string) (models.BlacklistedToken, error) {
	var entity models2.BlacklistedTokenDb

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionFirstByToken, "Buscando token en lista negra", zap.String("token", token))

	res := repository.Conn.
		WithContext(ctx).
		Table(blacklistedTableName).
		Where("token = ?", token).
		First(&entity)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			logger.StandardWarn(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionFirstByToken, "Token no encontrado", zap.String("token", token))
			return entity.ToDomainModel(), response_capture.NewErrorME(ctx, http.StatusNotFound, res.Error, constants.ErrorMessageTokenNotFound)
		}
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionFirstByToken, res.Error, "Error consultando token en base de datos", zap.String("token", token))
		return entity.ToDomainModel(), response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingToken)
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionFirstByToken, "Token encontrado exitosamente", zap.String("token", token))
	return entity.ToDomainModel(), nil
}

func (repository *BlacklistedTokenRepository) FirstByTokenNil(ctx context.Context, token string) (models.BlacklistedToken, error) {
	var entity models2.BlacklistedTokenDb

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionFirstByTokenNil, "Buscando token (retorno nil si no existe)", zap.String("token", token))

	res := repository.Conn.
		WithContext(ctx).
		Table(blacklistedTableName).
		Where("token = ?", token).
		First(&entity)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			logger.StandardWarn(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionFirstByTokenNil, "Token no encontrado (retornando nil)", zap.String("token", token))
			return models.BlacklistedToken{}, nil
		}
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionFirstByTokenNil, res.Error, "Error consultando token", zap.String("token", token))
		return models.BlacklistedToken{}, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingToken)
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionFirstByTokenNil, "Token encontrado exitosamente", zap.String("token", token))
	return entity.ToDomainModel(), nil
}

func (repository *BlacklistedTokenRepository) DeleteExpired(ctx context.Context, before time.Time) (int64, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionDeleteExpired, "Eliminando tokens expirados", zap.Time("fecha_limite", before))

	result := repository.Conn.
		WithContext(ctx).
		Where("expires_at < ?", before).
		Delete(&models.BlacklistedToken{})

	if result.Error != nil {
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionDeleteExpired, result.Error, "Error al eliminar tokens expirados")
		return 0, result.Error
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionDeleteExpired, "Tokens expirados eliminados exitosamente", zap.Int64("filas_eliminadas", result.RowsAffected))
	return result.RowsAffected, nil
}

func (repository *BlacklistedTokenRepository) Save(ctx context.Context, body models.BlacklistedToken) (models.BlacklistedToken, error) {
	var (
		entity models2.BlacklistedTokenDb
		err    error
	)

	entity.Load(body)
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionSave, "Guardando token en base de datos", zap.Any("body", body))

	tx := repository.Conn.WithContext(ctx).Begin()

	if entity.ID > 0 {
		err = tx.Save(&entity).Error
	} else {
		err = tx.Create(&entity).Error
	}

	if err != nil {
		tx.Rollback()
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionSave, err, "Error guardando token", zap.Any("body", body))
		return entity.ToDomainModel(), response_capture.NewErrorME(ctx, http.StatusBadRequest, err, fmt.Sprintf(constants.ErrorMessageSavingToken, err.Error()))
	}

	tx.Commit()
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleBlacklistedToken, constants.FunctionSave, "Token guardado correctamente", zap.Int64("token_id", entity.ID))
	return entity.ToDomainModel(), nil
}
