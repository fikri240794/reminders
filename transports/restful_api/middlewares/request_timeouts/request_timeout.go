package request_timeouts

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	base_vms "github.com/fikri240794/golib/models/vms"
	"github.com/gofiber/fiber/v2"
)

func (m *RequestTimeoutMiddleware) RequestTimeout(ctx *fiber.Ctx) error {
	var (
		rtoCtx          context.Context
		cancel          context.CancelFunc
		timeoutDuration time.Duration
		resVM           *base_vms.ResponseVM
		err             error
	)

	timeoutDuration = time.Duration(m.cfg.RESTfulAPIServer.RequestTimeoutInSec) * time.Second
	rtoCtx, cancel = context.WithTimeout(ctx.UserContext(), timeoutDuration)

	defer cancel()

	ctx.SetUserContext(rtoCtx)
	err = ctx.Next()

	if err != nil && (errors.Is(err, context.DeadlineExceeded) ||
		strings.Contains(strings.ToLower(err.Error()), context.DeadlineExceeded.Error())) {
		resVM = base_vms.NewResponseVM().
			SetError(base_vms.NewResponseErrorVM().
				SetCode(http.StatusRequestTimeout).
				SetMessage(http.StatusText(http.StatusRequestTimeout)))
		ctx.Status(resVM.Error.Code)

		return ctx.JSON(resVM)
	}

	if errors.Is(ctx.UserContext().Err(), context.DeadlineExceeded) &&
		strings.Contains(strings.ToLower(string(ctx.Response().Body())), context.DeadlineExceeded.Error()) {
		resVM = base_vms.NewResponseVM().
			SetError(base_vms.NewResponseErrorVM().
				SetCode(http.StatusRequestTimeout).
				SetMessage(http.StatusText(http.StatusRequestTimeout)))
		ctx.Status(resVM.Error.Code)

		return ctx.JSON(resVM)
	}

	return nil
}
