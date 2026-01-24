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

type ISkillsController interface {
	SkillsSearch(c echo.Context) error
	SkillsList(c echo.Context) error
	SkillsProfile(c echo.Context) error
}

type SkillsController struct {
	SkillsServices services.ISkillsServices
	Validator      validator.IValidator
}

// SkillsSearch Busqueda de habilidades.
// @Summary Busqueda de habilidades
// @Description Permite buscar habilidades activas del marketplace
// @Tags Skills
// @Accept json
// @Produce json
// @Param q query string false "Texto de búsqueda"
// @Param category_id query int false "Categoría"
// @Param limit query int false "Límite"
// @Param offset query int false "Offset"
// @Success 200 {object} dto.ListSkillsResponseDto
// @Router /v1/api/core/skills/search [get]
func (controller *SkillsController) SkillsSearch(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.ParamsSkillsSearchDto
		dto    dto.SkillsResponseDto
		data   []models.Skills
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionSkillsList, "Busqueda de habilidades",
		zap.String("endpoint", "/v1/api/core/skills/search"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

	err = entity.BindSkillsSearchFilter(c)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	err = utils.BindAndValidate(c, controller.Validator, &entity)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Initializing")
	data, err = controller.SkillsServices.SkillsSearch(ctx, entity.ToModel())
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, dto.FromModel(data))
}

// SkillsList lista de habilidades del marketplace.
// @Summary Listar habilidades
// @Description Devuelve un listado paginado de habilidades activas disponibles en el marketplace. Permite alimentar búsqueda, filtros y matching de taskers.
// @Tags Skills
// @Accept json
// @Produce json
// @Success 200 {object} dto.ListSkillsResponseDto "Listado de habilidades"
// @Router /v1/api/core/skills/list [get]
func (controller *SkillsController) SkillsList(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		dto  dto.SkillsResponseDto
		data []models.Skills
		err  error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionSkillsList, "Listar habilidades",
		zap.String("endpoint", "/v1/api/core/skills/list"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

	logger.Info(ctx, "Initializing")
	data, err = controller.SkillsServices.SkillsList(ctx)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, dto.FromModel(data))
}
