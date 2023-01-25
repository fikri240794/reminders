package logger

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (m *LoggerMiddleware) Log(ctx *fiber.Ctx) error {
	var (
		startTime int64
		endTime   int64
		reqRes    *requestResponse
		err       error
	)

	startTime = time.Now().UnixMilli()
	err = ctx.Next()
	endTime = time.Now().UnixMilli()
	reqRes = newRequestResponse(ctx, startTime, endTime)

	switch ctx.Response().StatusCode() {
	case http.StatusOK, http.StatusCreated:
		log.Info().Interface("metadata", reqRes).Msg("")
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound:
		log.Warn().Interface("metadata", reqRes).Msg("")
	case http.StatusInternalServerError, http.StatusRequestTimeout, http.StatusTooManyRequests, http.StatusBadGateway, http.StatusGatewayTimeout:
		log.Error().Interface("metadata", reqRes).Msg("")
	}

	return err
}
