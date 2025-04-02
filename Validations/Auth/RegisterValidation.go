package Auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Middlewares"
	"github.com/mahdic200/weava/Rules"
)

func Register() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "first_name",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(3, 255)},
		},
		{
			FieldName: "last_name",
			Rules:     []Rules.ValidationRule{Rules.Sometimes, Rules.LengthBetween(3, 255)},
		},
		{
			FieldName: "email",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.Email, Rules.MaxLength(255), Rules.Unique("email", "users")},
		},
		{
			FieldName: "phone",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.Regex(`09\d{9}`), Rules.Unique("phone", "users")},
		},
		{
			FieldName: "password",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(8, 16)},
		},
	})
}
