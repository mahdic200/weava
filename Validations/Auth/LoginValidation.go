package Auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Middlewares"
	"github.com/mahdic200/weava/Rules"
)

func Login() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "email",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.Email},
		},
		{
			FieldName: "password",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.MinLength(8)},
		},
	})
}
