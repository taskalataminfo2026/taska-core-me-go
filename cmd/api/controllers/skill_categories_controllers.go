package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
	"go.uber.org/zap"
	"net/http"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/controllers/dto"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/services"
	"taska-core-me-go/cmd/api/utils"
	"taska-core-me-go/cmd/api/validator"
)

//go:generate mockgen -destination=../mocks/controllers/$GOFILE -package=mcontrollers -source=./$GOFILE

type ISkillsCategoriesController interface {
	Save(c echo.Context) error
	Update(c echo.Context) error
}

type SkillsCategoriesController struct {
	SkillsCategoriesServices services.ISkillsCategoriesServices
	Validator                validator.IValidator
}

// Save guarda la relación entre una skill y una categoría.
// @Summary Asociar skill a categoría
// @Description Crea o actualiza la relación entre una skill y una categoría del marketplace.
// @Tags SkillsCategories
// @Accept json
// @Produce json
// @Param request body dto.ParamsSkillsCategorySaveDto true "Datos de la relación skill-categoría"
// @Success 200 {object} dto.ListSkillsResponseDto "Categoría guardada correctamente"
// @Router /v1/api/core/skills-categories/save [post]
func (controller *SkillsCategoriesController) Save(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.ParamsSkillsCategorySaveDto
		data   models.SkillCategory
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkillsAndCategories, constants.FunctionSkillsCategoriesSave, "Solicitud para asociar skill a categoría",
		zap.String("endpoint", "/v1/api/core/skills-categories/save"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

	err = utils.BindAndValidate(c, controller.Validator, &entity)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Initializing")
	data, err = controller.SkillsCategoriesServices.Save(ctx, entity.ToModel())
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, dto.SkillCategoryToDto(data))
}

// Update actualiza la relación entre una skill y una categoría.
// @Summary Actualizar relación skill-categoría
// @Description Actualiza los datos de una relación existente entre una skill y una categoría del marketplace.
// @Tags SkillsCategories
// @Accept json
// @Produce json
// @Param id path int true "ID de la relación skill-categoría"
// @Param request body dto.ParamsSkillsCategorySaveDto true "Datos de la relación skill-categoría"
// @Success 200 {object} dto.SkillCategoryResponseDto "Relación skill-categoría actualizada correctamente"
// @Router /v1/api/core/skills-categories/:id [put]
func (controller *SkillsCategoriesController) Update(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.ParamsSkillsCategorySaveDto
		path   dto.ParamsSkillsCategoryRequestDTO
		data   models.SkillCategory
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkillsAndCategories, constants.FunctionSkillsCategoriesUpdate, "Solicitud para actualizar relación skill-categoría",
		zap.String("endpoint", "/v1/api/core/skills-categories/:id"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

	err = utils.BindAndValidate(c, controller.Validator, &entity)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	err = path.ParseIDFromParam(c)
	if err != nil {
		return response_capture.RespondError(c,
			response_capture.NewErrorME(ctx, http.StatusBadRequest, nil, err.Error()))
	}

	logger.Info(ctx, "Initializing")
	data, err = controller.SkillsCategoriesServices.Update(ctx, path.ID, entity.ToModel())
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, dto.SkillCategoryToDto(data))
}
