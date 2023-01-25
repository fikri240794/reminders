package reminders

import (
	base_vms "github.com/fikri240794/golib/models/vms"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
)

type ReminderPaginationRequestVM struct {
	*base_vms.PaginationRequestVM
	Keyword string `json:"keyword" query:"keyword"`
}

func NewReminderPaginationRequestVM() *ReminderPaginationRequestVM {
	return &ReminderPaginationRequestVM{
		PaginationRequestVM: base_vms.NewPaginationRequestVM(),
	}
}

func (vm *ReminderPaginationRequestVM) MapToReminderPaginationParamsDTO() *dtos.ReminderPaginationParamsDTO {
	var dto *dtos.ReminderPaginationParamsDTO = dtos.NewReminderPaginationParamsDTO().
		SetKeyword(vm.Keyword)

	dto.PaginationParamsDTO = dto.PaginationParamsDTO.SetSkip(vm.Skip).
		SetTake(vm.Take)

	return dto
}
