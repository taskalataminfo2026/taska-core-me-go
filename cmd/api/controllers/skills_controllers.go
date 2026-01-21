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
}

type SkillsController struct {
	SkillsServices services.ISkillsServices
	Validator      validator.IValidator
}

// SkillsSearch Busquedad de habilidades.
// @Summary Busquedad de habilidades
// @Description Permite buscar las habilidades del usuario.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequestDTO true "Autentica al usuario"
// @Success 200 {object} dto.LoginResponseDTO "Inicio de sesión exitoso"
// @Router /v1/api/core/skills/search [get]
func (controller *SkillsController) SkillsSearch(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.SkillsRequestDto
		dto    dto.SkillsResponseDto
		data   []models.SkillsResponse
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionSkillsList, "Iniciando proceso de login",
		zap.String("endpoint", "/v1/api/core/skills/search"),
		zap.String("method", c.Request().Method),
		zap.String("ip", c.RealIP()),
	)

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

// SkillsList lista de habilidades.
// @Summary Inicio de sesión
// @Description Inicia sesión con credenciales válidas y genera tokens de acceso y actualización.
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} dto.LoginResponseDTO "Inicio de sesión exitoso"
// @Router /v1/api/core/skills/list [get]
func (controller *SkillsController) SkillsList(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		dto  dto.SkillsResponseDto
		data []models.SkillsResponse
		err  error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionSkillsList, "Iniciando proceso de login",
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
