package reminders

import (
	base_vms "github.com/fikri240794/golib/models/vms"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	vms "github.com/fikri240794/reminders/transports/restful_api/models/vms/reminders"
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Find a reminder
//	@Description	Find a reminder
//	@Tags			Reminders
//	@Param			keyword	query		string	false	"keyword"
//	@Success		200		{object}	base_vms.ResponseVM{data=vms.ReminderVM}
//	@Router			/nreminders/find [get]
func (c *ReminderController) FindOne(ctx *fiber.Ctx) error {
	var (
		keyword   string
		resultDTO *dtos.ReminderDTO
		resDataVM *vms.ReminderVM
		resVM     *base_vms.ResponseVM
		err       error
	)

	keyword = ctx.Query("keyword")
	resultDTO, err = c.reminderService.FindOne(ctx.UserContext(), keyword)

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
