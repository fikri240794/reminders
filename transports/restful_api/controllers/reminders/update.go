package reminders

import (
	base_enums "github.com/fikri240794/golib/models/enums"
	base_vms "github.com/fikri240794/golib/models/vms"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	vms "github.com/fikri240794/reminders/transports/restful_api/models/vms/reminders"
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Update a reminder by ID
//	@Description	Update a reminder by ID
//	@Tags			Reminder
//	@Param			id		path		string					true	"id"
//	@Param			request	body		vms.UpdateReminderVM	true	"request"
//	@Success		200		{object}	base_vms.ResponseVM{data=vms.ReminderVM}
//	@Router			/reminders/{id} [put]
func (c *ReminderController) Update(ctx *fiber.Ctx) error {
	var (
		reqDataVM *vms.UpdateReminderVM
		paramsDTO *dtos.UpdateReminderDTO
		resultDTO *dtos.ReminderDTO
		resDataVM *vms.ReminderVM
		resVM     *base_vms.ResponseVM
		err       error
	)

	reqDataVM = vms.NewUpdateReminderVM()
	err = ctx.BodyParser(reqDataVM)

	if err != nil {
		resVM = base_vms.NewResponseVM().
			SetError(base_vms.NewResponseErrorVM().
				SetCode(base_enums.INTERNAL_SERVER_ERROR.ToHttpStatusCode()).
				SetMessage(err.Error()))
		ctx.Status(resVM.Error.Code)

		return ctx.JSON(resVM)
	}

	err = ctx.ParamsParser(reqDataVM)

	if err != nil {
		resVM = base_vms.NewResponseVM().
			SetError(base_vms.NewResponseErrorVM().
				SetCode(base_enums.INTERNAL_SERVER_ERROR.ToHttpStatusCode()).
				SetMessage(err.Error()))
		ctx.Status(resVM.Error.Code)

		return ctx.JSON(resVM)
	}

	paramsDTO = reqDataVM.MapToUpdateReminderDTO()
	resultDTO, err = c.reminderService.Update(ctx.UserContext(), paramsDTO)

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
