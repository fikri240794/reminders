//go:generate rm -rf wire_gen.go
//go:generate go run github.com/google/wire/cmd/wire

package servers

import (
	"github.com/fikri240794/reminders/configs"
	"github.com/gofiber/fiber/v2"
)

type IRESTfulAPIServer interface {
	Serve()
}

type RESTfulAPIServer struct {
	cfg    *configs.Config
	server *fiber.App
}
