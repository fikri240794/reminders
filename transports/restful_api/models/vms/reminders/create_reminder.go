package reminders

import "github.com/fikri240794/reminders/internal/reminders/models/dtos"

type CreateReminderVM struct {
	Note string `json:"note"`
}

func NewCreateReminderVM() *CreateReminderVM {
	return &CreateReminderVM{}
}

func (vm *CreateReminderVM) MapToCreateReminderDTO() *dtos.CreateReminderDTO {
	var dto *dtos.CreateReminderDTO = dtos.NewCreateReminderDTO()

	dto.Note = vm.Note

	return dto
}
