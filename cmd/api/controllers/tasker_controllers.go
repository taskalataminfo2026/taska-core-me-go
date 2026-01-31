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

type ITaskerController interface {
	TaskerProfile(c echo.Context) error
}

type TaskerController struct {
	TaskerServices services.ITaskerServices
	Validator      validator.IValidator
}

// TaskerProfile obtiene las habilidades de un tasker.
// @Summary Habilidades de un tasker
// @Description Devuelve las habilidades activas asociadas a un tasker, incluyendo nivel, verificación y métricas. Endpoint base para perfil público y matching.
// @Tags Tasker
// @Accept json
// @Produce json
// @Param id_user path int true "ID del tasker"
// @Success 200 {object} dto.ListSkillsResponseDto
// @Router /v1/api/core/tasker/{id_user}/skills [get]
func (controller *TaskerController) TaskerProfile(c echo.Context) error {
	ctx := utils.CreateRequestContext(c)

	var (
		entity dto.ParamsProfileDto
		dto    dto.TaskerDto
		data   models.Tasker
		err    error
	)

	logger.StandardInfo(ctx, constants.LayerController, constants.ModuleSkills, constants.FunctionSkillsList, "Habilidades de un tasker",
		zap.String("endpoint", "/v1/api/core/tasker/{id_user}/skills"),
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
	data, err = controller.TaskerServices.GetTasker(ctx, entity.ToModel())
	if err != nil {
		return response_capture.RespondError(c, err)
	}

	logger.Info(ctx, "Finalized")
	return response_capture.HandleResponse(c, http.StatusOK, constants.StatusOk, dto.FromModel(data))
}
