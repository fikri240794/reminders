package reminders

import "github.com/fikri240794/reminders/internal/reminders/models/dtos"

type UpdateReminderVM struct {
	ID   string `params:"id"`
	Note string `json:"note"`
}

func NewUpdateReminderVM() *UpdateReminderVM {
	return &UpdateReminderVM{}
}

func (vm *UpdateReminderVM) MapToUpdateReminderDTO() *dtos.UpdateReminderDTO {
	var dto *dtos.UpdateReminderDTO = dtos.NewUpdateReminderDTO()

	dto.ID = vm.ID
	dto.Note = vm.Note

	return dto
}
