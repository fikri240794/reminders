package dtos

import (
	"time"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
	"github.com/google/uuid"
)

type UpdateReminderDTO struct {
	ID   string
	Note string
}

func NewUpdateReminderDTO() *UpdateReminderDTO {
	return &UpdateReminderDTO{}
}

func (dto *UpdateReminderDTO) Validate() error {
	var (
		errFields []errors.ErrorField
		err       error
	)

	errFields = []errors.ErrorField{}

	if dto.ID == "" {
		errFields = append(errFields, errors.NewErrorField("id", "required field"))
	}

	_, err = uuid.Parse(dto.ID)

	if err != nil {
		errFields = append(errFields, errors.NewErrorField("id", "value is not uuid"))
	}

	if dto.Note == "" {
		errFields = append(errFields, errors.NewErrorField("note", "required field"))
	}

	if len(errFields) > 0 {
		return errors.NewError(base_enums.BAD_REQUEST, base_enums.BAD_REQUEST.ToString(), errFields...)
	}

	return nil
}

func (dto *UpdateReminderDTO) MapToReminderEntity(e *entities.ReminderEntity) *entities.ReminderEntity {
	e.Note = dto.Note
	e.UpdatedAt = time.Now().UnixMilli()

	return e
}
