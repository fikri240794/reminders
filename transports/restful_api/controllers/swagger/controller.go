//go:generate rm -rf docs
//go:generate go run github.com/swaggo/swag/cmd/swag init -g ./../../../../main.go --pd --parseInternal
//go:generate go run github.com/swaggo/swag/cmd/swag fmt -g ./../../../../main.go -d ./../../../restful_api/

package swagger

import (
	_ "github.com/fikri240794/reminders/transports/restful_api/controllers/swagger/docs"
	"github.com/gofiber/fiber/v2"
)

type ISwaggerController interface {
	Index(ctx *fiber.Ctx) error
}
