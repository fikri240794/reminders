package reminders

import "github.com/fikri240794/reminders/internal/reminders/models/dtos"

type ReminderVM struct {
	ID        string `json:"id"`
	Note      string `json:"note"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func NewReminderVM() *ReminderVM {
	return &ReminderVM{}
}

func (vm *ReminderVM) MapFromReminderDTO(dto *dtos.ReminderDTO) *ReminderVM {
	vm.ID = dto.ID
	vm.Note = dto.Note
	vm.CreatedAt = dto.CreatedAt
	vm.UpdatedAt = dto.UpdatedAt

	return vm
}
