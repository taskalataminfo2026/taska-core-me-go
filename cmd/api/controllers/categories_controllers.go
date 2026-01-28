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
	List(c echo.Context) error
}

type CategoriesController struct {
	CategoriesServices services.ICategoriesServices
	Validator          validator.IValidator
}

// List lista de habilidades del marketplace.
// @Summary Listar habilidades
// @Description Devuelve un listado paginado de habilidades activas disponibles en el marketplace. Permite alimentar búsqueda, filtros y matching de taskers.
// @Tags Skills
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
