package compresses

import "github.com/gofiber/fiber/v2"

type ICompressMiddleware interface {
	Compress(ctx *fiber.Ctx) error
}
