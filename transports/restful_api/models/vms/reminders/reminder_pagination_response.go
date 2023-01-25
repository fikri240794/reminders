package reminders

import (
	base_vms "github.com/fikri240794/golib/models/vms"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
)

type ReminderPaginationResponseVM struct {
	*base_vms.PaginationResponseVM
	Collection []*ReminderVM `json:"collection"`
}

func NewReminderPaginationResponseVM() *ReminderPaginationResponseVM {
	return &ReminderPaginationResponseVM{
		PaginationResponseVM: base_vms.NewPaginationResponseVM(),
	}
}

func (vm *ReminderPaginationResponseVM) MapFromReminderPaginationResultDTO(dto *dtos.ReminderPaginationResultDTO) *ReminderPaginationResponseVM {
	vm.PaginationResponseVM = vm.SetCount(dto.Count)
	vm.Collection = []*ReminderVM{}

	for i := 0; i < len(dto.Collection); i++ {
		vm.Collection = append(vm.Collection, NewReminderVM().
			MapFromReminderDTO(dto.Collection[i]))
	}

	return vm
}
