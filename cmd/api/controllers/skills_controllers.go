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
	Search(c echo.Context) error
	List(c echo.Context) error
	Save(c echo.Context) error
	Update(c echo.Context) error
}

type SkillsController struct {
	SkillsServices services.ISkillsServices
	Validator      validator.IValidator
}

// Search Busqueda de habilidades.
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
func (controller *SkillsController) Search(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		param  dto.ParamsSkillsSearchDto
		entity dto.SkillsDto
		data   []models.Skills
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionSkillsSearch, "Busqueda de habilidades",
		zap.String("endpoint", "/v1/api/core/skills/search"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

	err = param.BindSkillsSearchFilter(c)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	err = utils.BindAndValidate(c, controller.Validator, &param)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Initializing")
	data, err = controller.SkillsServices.Search(ctx, param.ToModel())
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, entity.FromModel(data))
}

// List lista de habilidades del marketplace.
// @Summary Listar habilidades
// @Description Devuelve un listado paginado de habilidades activas disponibles en el marketplace. Permite alimentar búsqueda, filtros y matching de taskers.
// @Tags Skills
// @Accept json
// @Produce json
// @Success 200 {object} dto.ListSkillsResponseDto "Listado de habilidades"
// @Router /v1/api/core/skills/list [get]
func (controller *SkillsController) List(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.SkillsDto
		data   []models.Skills
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionSkillsList, "Listar habilidades",
		zap.String("endpoint", "/v1/api/core/skills/list"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

	logger.Info(ctx, "Initializing")
	data, err = controller.SkillsServices.List(ctx)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, entity.FromModel(data))
}

// Save Crea una habilidad (skill) del marketplace.
// @Summary Crear skill
// @Description Crea una nueva habilidad (ej: slug).
// @Tags Skills
// @Accept json
// @Produce json
// @Param request body dto.ParamsSkillsUpsertDto true "Datos de la habilidad a crear"
// @Success 200 {object} dto.SkillsResponseDto "Habilidad creada correctamente"
// @Router /v1/api/core/skills/save [post]
func (controller *SkillsController) Save(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.ParamsSkillsUpsertDto
		data   models.Skills
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionSkillsSave, "Guardar habilidades",
		zap.String("endpoint", "/v1/api/core/skills/save"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

	err = utils.BindAndValidate(c, controller.Validator, &entity)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Initializing")
	data, err = controller.SkillsServices.Save(ctx, entity.ToModel())
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, dto.SkillToDto(data))
}

// Update actualiza una habilidad (skill) del marketplace.
// @Summary Actualizar skill
// @Description Actualiza una nueva habilidad según su identificador lógico (ej: slug).
// @Tags Skills
// @Accept json
// @Produce json
// @Param request body dto.ParamsSkillsUpsertDto true "Datos de la habilidad a actualizar"
// @Success 200 {object} dto.SkillsResponseDto "Habilidad actualizar correctamente"
// @Router /v1/api/core/skills/:id [put]
func (controller *SkillsController) Update(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.ParamsSkillsUpsertDto
		path   dto.ParamsSkillsRequestDTO
		data   models.Skills
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionSkillsUpdate, "Actualizar habilidades",
		zap.String("endpoint", "/v1/api/core/skills/:id"),
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
	data, err = controller.SkillsServices.Update(ctx, path.ID, entity.ToModel())
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, dto.SkillToDto(data))
}
