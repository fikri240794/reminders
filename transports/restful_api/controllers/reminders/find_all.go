package reminders

import (
	base_enums "github.com/fikri240794/golib/models/enums"
	base_vms "github.com/fikri240794/golib/models/vms"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	vms "github.com/fikri240794/reminders/transports/restful_api/models/vms/reminders"
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Find all reminder
//	@Description	Find all reminder
//	@Tags			Reminders
//	@Param			request	query		vms.ReminderPaginationRequestVM	true	"request"
//	@Success		200		{object}	base_vms.ResponseVM{data=vms.ReminderPaginationResponseVM}
//	@Failure		400		{object}	base_vms.ResponseVM{}
//	@Failure		404		{object}	base_vms.ResponseVM{}
//	@Failure		500		{object}	base_vms.ResponseVM{}
//	@Router			/reminders [get]
func (c *ReminderController) FindAll(ctx *fiber.Ctx) error {
	var (
		reqDataVM *vms.ReminderPaginationRequestVM
		paramsDTO *dtos.ReminderPaginationParamsDTO
		resultDTO *dtos.ReminderPaginationResultDTO
		resDataVM *vms.ReminderPaginationResponseVM
		resVM     *base_vms.ResponseVM
		err       error
	)

	reqDataVM = vms.NewReminderPaginationRequestVM()
	err = ctx.QueryParser(reqDataVM)

	if err != nil {
		resVM = base_vms.NewResponseVM().
			SetError(base_vms.NewResponseErrorVM().
				SetCode(base_enums.INTERNAL_SERVER_ERROR.ToHttpStatusCode()).
				SetMessage(err.Error()))
		ctx.Status(resVM.Error.Code)

		return ctx.JSON(resVM)
	}

	paramsDTO = reqDataVM.MapToReminderPaginationParamsDTO()
	resultDTO, err = c.reminderService.FindAll(ctx.UserContext(), paramsDTO)

	if err != nil {
		resVM = base_vms.NewResponseVM().
			SetErrorFromError(err)
		ctx.Status(resVM.Error.Code)

		return ctx.JSON(resVM)
	}

	resDataVM = vms.NewReminderPaginationResponseVM().
		MapFromReminderPaginationResultDTO(resultDTO)
	resVM = base_vms.NewResponseVM().
		SetData(resDataVM)

	return ctx.JSON(resVM)
}
