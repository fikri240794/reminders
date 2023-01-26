package middlewares

import (
	"github.com/fikri240794/reminders/transports/restful_api/middlewares/compresses"
	"github.com/fikri240794/reminders/transports/restful_api/middlewares/corses"
	"github.com/fikri240794/reminders/transports/restful_api/middlewares/etags"
	"github.com/fikri240794/reminders/transports/restful_api/middlewares/favicons"
	"github.com/fikri240794/reminders/transports/restful_api/middlewares/loggers"
	"github.com/fikri240794/reminders/transports/restful_api/middlewares/recovers"
	"github.com/fikri240794/reminders/transports/restful_api/middlewares/request_timeouts"
	"github.com/fikri240794/reminders/transports/restful_api/middlewares/requestids"
)

type Middleware struct {
	compressMiddleware       compresses.ICompressMiddleware
	loggerMiddleware         loggers.ILoggerMiddleware
	recoverMiddleware        recovers.IRecoverMiddleware
	faviconMiddleware        favicons.IFaviconMiddleware
	corsMiddleware           corses.ICORSMiddleware
	requestTimeoutMiddleware request_timeouts.IRequestTimeoutDMiddleware
	etagMiddleware           etags.IETagMiddleware
	requestIDMiddleware      requestids.IRequestIDMiddleware
}

func NewMiddleware(
	compressMiddleware compresses.ICompressMiddleware,
	loggerMiddleware loggers.ILoggerMiddleware,
	recoverMiddleware recovers.IRecoverMiddleware,
	faviconMiddleware favicons.IFaviconMiddleware,
	corsMiddleware corses.ICORSMiddleware,
	requestTimeoutMiddleware request_timeouts.IRequestTimeoutDMiddleware,
	etagMiddleware etags.IETagMiddleware,
	requestIDMiddleware requestids.IRequestIDMiddleware,
) *Middleware {
	return &Middleware{
		compressMiddleware:       compressMiddleware,
		loggerMiddleware:         loggerMiddleware,
		recoverMiddleware:        recoverMiddleware,
		faviconMiddleware:        faviconMiddleware,
		corsMiddleware:           corsMiddleware,
		requestTimeoutMiddleware: requestTimeoutMiddleware,
		etagMiddleware:           etagMiddleware,
		requestIDMiddleware:      requestIDMiddleware,
	}
}
