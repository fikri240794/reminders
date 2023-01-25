package dtos

import (
	"time"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
	"github.com/google/uuid"
)

type CreateReminderDTO struct {
	Note string
}

func NewCreateReminderDTO() *CreateReminderDTO {
	return &CreateReminderDTO{}
}

func (dto *CreateReminderDTO) Validate() error {
	if dto.Note == "" {
		return errors.NewError(base_enums.BAD_REQUEST, base_enums.BAD_REQUEST.ToString(), errors.NewErrorField("note", "required field"))
	}

	return nil
}

func (dto *CreateReminderDTO) MapToReminderEntity() *entities.ReminderEntity {
	var e *entities.ReminderEntity = entities.NewReminderEntity()

	e.ID = uuid.NewString()
	e.Note = dto.Note
	e.CreatedAt = time.Now().UnixMilli()
	e.UpdatedAt = e.CreatedAt

	return e
}
