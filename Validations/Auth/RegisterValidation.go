package Auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Middlewares"
	"github.com/mahdic200/weava/Rules"
)

func Store() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		// !TODO
	})
}
