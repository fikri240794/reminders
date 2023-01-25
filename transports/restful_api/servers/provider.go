//go:build wireinject
// +build wireinject

package servers

import (
	"github.com/bytedance/sonic"
	"github.com/fikri240794/reminders/configs"
	"github.com/fikri240794/reminders/datasources"
	"github.com/fikri240794/reminders/internal"
	"github.com/fikri240794/reminders/transports/restful_api/controllers"
	"github.com/fikri240794/reminders/transports/restful_api/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/rs/zerolog"
)

func newHTTPServer(cfg *configs.Config) *fiber.App {
	return fiber.New(fiber.Config{
		AppName:           cfg.RESTfulAPIServer.Name,
		EnablePrintRoutes: cfg.RESTfulAPIServer.PrintRoutes,
		Prefork:           cfg.RESTfulAPIServer.Prefork,
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
	})
}

func newRESTfulAPIServer(cfg *configs.Config, server *fiber.App, middleware *middlewares.Middleware, controller *controllers.Controller) *RESTfulAPIServer {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	if cfg.Environment != "production" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	return &RESTfulAPIServer{
		cfg:    cfg,
		server: server,
	}
}

func NewRESTfulAPIServer() *RESTfulAPIServer {
	wire.Build(
		configs.Read,
		datasources.DatasourceModule,
		internal.PersistenceModule,
		internal.RepositoryModule,
		internal.ServiceModule,
		newHTTPServer,
		middlewares.MiddlewareModule,
		controllers.ControllerModule,
		newRESTfulAPIServer,
	)

	return nil
}
