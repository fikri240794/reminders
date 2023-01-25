package reminders

import (
	"net/http"

	base_enums "github.com/fikri240794/golib/models/enums"
	base_vms "github.com/fikri240794/golib/models/vms"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	vms "github.com/fikri240794/reminders/transports/restful_api/models/vms/reminders"
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Create a reminder
//	@Description	Create a reminder
//	@Tags			Reminders
//	@Param			CreateReminderVM	body		vms.CreateReminderVM	true	"CreateReminderVM"
//	@Success		201					{object}	base_vms.ResponseVM{data=vms.ReminderVM}
//	@Router			/reminders [post]
func (c *ReminderController) Create(ctx *fiber.Ctx) error {
	var (
		reqDataVM *vms.CreateReminderVM
		paramsDTO *dtos.CreateReminderDTO
		resultDTO *dtos.ReminderDTO
		resDataVM *vms.ReminderVM
		resVM     *base_vms.ResponseVM
		err       error
	)

	reqDataVM = vms.NewCreateReminderVM()
	err = ctx.BodyParser(reqDataVM)

	if err != nil {
		resVM = base_vms.NewResponseVM().
			SetError(base_vms.NewResponseErrorVM().
				SetCode(base_enums.INTERNAL_SERVER_ERROR.ToHttpStatusCode()).
				SetMessage(err.Error()))
		ctx.Status(resVM.Error.Code)

		return ctx.JSON(resVM)
	}

	paramsDTO = reqDataVM.MapToCreateReminderDTO()
	resultDTO, err = c.reminderService.Create(ctx.UserContext(), paramsDTO)

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
	ctx.Status(http.StatusCreated)

	return ctx.JSON(resVM)
}
