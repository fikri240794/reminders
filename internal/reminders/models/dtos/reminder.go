package dtos

import (
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

type ReminderDTO struct {
	ID        string
	Note      string
	CreatedAt int64
	UpdatedAt int64
}

func NewReminderDTO() *ReminderDTO {
	return &ReminderDTO{}
}

func (dto *ReminderDTO) MapFromReminderEntity(e *entities.ReminderEntity) *ReminderDTO {
	dto.ID = e.ID
	dto.Note = e.Note
	dto.CreatedAt = e.CreatedAt
	dto.UpdatedAt = e.UpdatedAt

	return dto
}
