package reminders

import (
	base_vms "github.com/fikri240794/golib/models/vms"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	vms "github.com/fikri240794/reminders/transports/restful_api/models/vms/reminders"
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Find a reminder by ID
//	@Description	Find a reminder by ID
//	@Tags			Reminders
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	base_vms.ResponseVM{data=vms.ReminderVM}
//	@Router			/reminders/{id} [get]
func (c *ReminderController) FindById(ctx *fiber.Ctx) error {
	var (
		id        string
		resultDTO *dtos.ReminderDTO
		resDataVM *vms.ReminderVM
		resVM     *base_vms.ResponseVM
		err       error
	)

	id = ctx.Params("id")
	resultDTO, err = c.reminderService.FindById(ctx.UserContext(), id)

	if err != nil {
		resVM = base_vms.NewResponseVM().
			SetErrorFromError(err)
		ctx.Status(resVM.Error.Code)

		return ctx.JSON(resVM)
	}

	resDataVM = vms.NewReminderVM().
		MapFromReminderDTO(resultDTO)
	resVM = base_vms.NewResponseVM().
		SetData(resDataVM)

	return ctx.JSON(resVM)
}
