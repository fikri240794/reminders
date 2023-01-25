package swagger

import (
	_ "github.com/fikri240794/reminders/docs"
	"github.com/gofiber/fiber/v2"
)

type ISwaggerController interface {
	Index(ctx *fiber.Ctx) error
}
