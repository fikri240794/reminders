package reminders

import (
	"github.com/gofiber/fiber/v2"
)

type IReminderController interface {
	FindAll(ctx *fiber.Ctx) error
	FindOne(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
