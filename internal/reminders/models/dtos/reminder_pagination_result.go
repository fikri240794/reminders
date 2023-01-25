package dtos

import (
	base_dtos "github.com/fikri240794/golib/models/dtos"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

type ReminderPaginationResultDTO struct {
	*base_dtos.PaginationResultDTO
	Collection []*ReminderDTO
}

func NewReminderPaginationResultDTO() *ReminderPaginationResultDTO {
	return &ReminderPaginationResultDTO{
		PaginationResultDTO: base_dtos.NewPaginationResultDTO(),
	}
}

func (dto *ReminderPaginationResultDTO) SetCount(count int64) *ReminderPaginationResultDTO {
	dto.PaginationResultDTO = dto.PaginationResultDTO.SetCount(count)

	return dto
}

func (dto *ReminderPaginationResultDTO) SetCollection(entities []*entities.ReminderEntity) *ReminderPaginationResultDTO {
	dto.Collection = []*ReminderDTO{}

	for i := 0; i < len(entities); i++ {
		dto.Collection = append(dto.Collection, NewReminderDTO().
			MapFromReminderEntity(entities[i]))
	}

	return dto
}
