package reminders

import (
	base_vms "github.com/fikri240794/golib/models/vms"
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Delete a reminder by ID
//	@Description	Delete a reminder by ID
//	@Tags			Reminders
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	base_vms.ResponseVM{data=bool}
//	@Router			/reminders/{id} [delete]
func (c *ReminderController) Delete(ctx *fiber.Ctx) error {
	var (
		id    string
		resVM *base_vms.ResponseVM
		err   error
	)

	id = ctx.Params("id")
	err = c.reminderService.Delete(ctx.UserContext(), id)

	if err != nil {
		resVM = base_vms.NewResponseVM().
			SetErrorFromError(err)
		ctx.Status(resVM.Error.Code)

		return ctx.JSON(resVM)
	}

	resVM = base_vms.NewResponseVM().
		SetData(true)

	return ctx.JSON(resVM)
}
