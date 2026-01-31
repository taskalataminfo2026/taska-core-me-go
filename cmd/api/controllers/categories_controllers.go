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

type ICategoriesController interface {
	Search(c echo.Context) error
	List(c echo.Context) error
	Save(c echo.Context) error
	Update(c echo.Context) error
}

type CategoriesController struct {
	CategoriesServices services.ICategoriesServices
	Validator          validator.IValidator
}

// Search realiza la búsqueda de habilidades activas del marketplace.
// @Summary Buscar habilidades
// @Description Permite buscar habilidades activas del marketplace usando texto, categoría y paginación. Se utiliza para exploración, filtros y matching.
// @Tags Category
// @Accept json
// @Produce json
// @Param q query string false "Texto de búsqueda (nombre o descripción)"
// @Param category_id query int false "ID de la categoría"
// @Param limit query int false "Cantidad máxima de resultados"
// @Param offset query int false "Desplazamiento para paginación"
// @Success 200 {object} dto.ListSkillsResponseDto "Listado de habilidades"
// @Router /v1/api/core/category/search [get]
func (controller *CategoriesController) Search(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		param  dto.ParamsCategorySearchDto
		entity dto.CategoryDto
		data   []models.Category
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionSkillsSearch, "Busqueda de habilidades",
		zap.String("endpoint", "/v1/api/core/category/search"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

	err = param.BindCategorySearchFilter(c)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	err = utils.BindAndValidate(c, controller.Validator, &param)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Initializing")
	data, err = controller.CategoriesServices.Search(ctx, param.ToModel())
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, entity.FromModel(data))
}

// List lista las categorías activas del marketplace.
// @Summary Listar categorías
// @Description Devuelve un listado de categorías activas del marketplace. Se utiliza para navegación, filtros y clasificación de habilidades.
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} dto.ListSkillsResponseDto "Listado de habilidades"
// @Router /v1/api/core/category/list [get]
func (controller *CategoriesController) List(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.CategoryDto
		data   []models.Category
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionCategoryList, "Listar habilidades",
		zap.String("endpoint", "/v1/api/core/category/list"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

	logger.Info(ctx, "Initializing")
	data, err = controller.CategoriesServices.List(ctx)
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, entity.FromModel(data))
}

// Save guarda una categoría del marketplace.
// @Summary Guardar categoría
// @Description Crea o actualiza una categoría del marketplace para organización, navegación y filtros.
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} dto.ListSkillsResponseDto "Categoría guardada correctamente"
// @Router /v1/api/core/category/save [post]
func (controller *CategoriesController) Save(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.ParamsCategorySaveDto
		data   models.Category
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionCategoryList, "Guardar habilidades",
		zap.String("endpoint", "/v1/api/core/category/save"),
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
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "ID de la categoría"
// @Param request body dto.ParamsCategorySaveDto true "Datos de la categoría"
// @Success 200 {object} dto.CategoryResponseDto "Categoría actualizada correctamente"
// @Router /v1/api/core/category/:id [put]
func (controller *CategoriesController) Update(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.ParamsCategorySaveDto
		path   dto.ParamsCategoryRequestDTO
		data   models.Category
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionCategoryList, "Actualizar habilidades",
		zap.String("endpoint", "/v1/api/core/category/:id"),
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
