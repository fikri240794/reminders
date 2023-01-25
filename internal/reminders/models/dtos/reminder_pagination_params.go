package dtos

import base_dtos "github.com/fikri240794/golib/models/dtos"

type ReminderPaginationParamsDTO struct {
	*base_dtos.PaginationParamsDTO
	Keyword string
}

func NewReminderPaginationParamsDTO() *ReminderPaginationParamsDTO {
	return &ReminderPaginationParamsDTO{
		PaginationParamsDTO: base_dtos.NewPaginationParamsDTO(),
	}
}

func (dto *ReminderPaginationParamsDTO) SetKeyword(keyword string) *ReminderPaginationParamsDTO {
	dto.Keyword = keyword

	return dto
}

func (dto *ReminderPaginationParamsDTO) Validate() error {
	return dto.PaginationParamsDTO.Validate()
}
