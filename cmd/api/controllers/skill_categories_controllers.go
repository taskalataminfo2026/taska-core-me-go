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
	CategoriesServices services.ICategoriesServices
	Validator          validator.IValidator
}

// Save guarda una categoría del marketplace.
// @Summary Guardar categoría
// @Description Crea o actualiza una categoría del marketplace para organización, navegación y filtros.
// @Tags Skills
// @Accept json
// @Produce json
// @Success 200 {object} dto.ListSkillsResponseDto "Categoría guardada correctamente"
// @Router /v1/api/core/skills-categories/save [post]
func (controller *SkillsCategoriesController) Save(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.ParamsCategorySaveDto
		data   models.Category
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionCategoryList, "Guardar habilidades",
		zap.String("endpoint", "/v1/api/core/skills-categories/save"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

	err = utils.BindAndValidate(c, controller.Validator, &entity)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Initializing")
	data, err = controller.CategoriesServices.Save(ctx, entity.ToModel())
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, dto.CategoryToDto(data))
}

// Update actualiza una categoría del marketplace.
// @Summary Actualizar categoría
// @Description Actualiza los datos de una categoría existente del marketplace para organización, navegación y filtros.
// @Tags Skills
// @Accept json
// @Produce json
// @Param id path int true "ID de la categoría"
// @Param request body dto.ParamsCategorySaveDto true "Datos de la categoría"
// @Success 200 {object} dto.CategoryResponseDto "Categoría actualizada correctamente"
// @Router /v1/api/core/skills-categories/:id [put]
func (controller *SkillsCategoriesController) Update(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.ParamsCategorySaveDto
		path   dto.ParamsCategoryRequestDTO
		data   models.Category
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionCategoryList, "Actualizar habilidades",
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
	data, err = controller.CategoriesServices.Update(ctx, path.ID, entity.ToModel())
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, dto.CategoryToDto(data))
}
