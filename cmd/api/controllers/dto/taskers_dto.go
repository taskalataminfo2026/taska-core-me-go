package dto

import (
	"github.com/labstack/echo/v4"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/utils"
)

type TaskerDto struct {
	ID int64 `json:"id"`
}

type ParamsProfileDto struct {
	ID int64 `json:"id"`
}

func (v *ParamsProfileDto) BindSkillsSearchFilter(c echo.Context) error {
	if err := utils.GetInt64Query(c, "id", &v.ID); err != nil {
		return err
	}

	return nil
}

func (v *ParamsProfileDto) ToModel() models.ParamsProfile {
	return models.ParamsProfile{
		ID: v.ID,
	}
}

func (s *TaskerDto) FromModel(tasker models.Tasker) TaskerDto {
	var taskerDto TaskerDto

	taskerDto.ID = tasker.ID

	return taskerDto
}
