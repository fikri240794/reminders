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
	"github.com/google/wire"
)

var MiddlewareModule wire.ProviderSet = wire.NewSet(
	compresses.NewCompressMiddleware,
	wire.Bind(new(compresses.ICompressMiddleware), new(*compresses.CompressMiddleware)),
	loggers.NewLoggerMiddleware,
	wire.Bind(new(loggers.ILoggerMiddleware), new(*loggers.LoggerMiddleware)),
	recovers.NewRecoverMiddleware,
	wire.Bind(new(recovers.IRecoverMiddleware), new(*recovers.RecoverMiddleware)),
	favicons.NewFaviconMiddleware,
	wire.Bind(new(favicons.IFaviconMiddleware), new(*favicons.FaviconMiddleware)),
	corses.NewCORSMiddleware,
	wire.Bind(new(corses.ICORSMiddleware), new(*corses.CORSMiddleware)),
	request_timeouts.NewRequestTimeoutMiddleware,
	wire.Bind(new(request_timeouts.IRequestTimeoutDMiddleware), new(*request_timeouts.RequestTimeoutMiddleware)),
	etags.NewETagMiddleware,
	wire.Bind(new(etags.IETagMiddleware), new(*etags.ETagMiddleware)),
	requestids.NewRequestIDMiddleware,
	wire.Bind(new(requestids.IRequestIDMiddleware), new(*requestids.RequestIDMiddleware)),
	NewMiddleware,
)
