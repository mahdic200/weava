package UserValidation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Middlewares"
	"github.com/mahdic200/weava/Rules"
)

func Store() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "first_name",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.LengthBetween(2, 255)},
		},
		{
			FieldName: "last_name",
			Rules:     []Rules.ValidationRule{Rules.Nullable, Rules.LengthBetween(2, 255)},
		},
		{
			FieldName: "image",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.File, Rules.Mimes("jpg", "jpeg", "png"), Rules.MaxSize(1000)},
		},
		{
			FieldName: "email",
			Rules:     []Rules.ValidationRule{Rules.Required, Rules.Email, Rules.LengthBetween(2, 255), Rules.Unique("email", "users")},
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

func Update() func(c *fiber.Ctx) error {
	return Middlewares.ValidationMiddleware([]Rules.FieldRules{
		{
			FieldName: "first_name",
			Rules:     []Rules.ValidationRule{Rules.Sometimes, Rules.LengthBetween(2, 255)},
		},
		{
			FieldName: "last_name",
			Rules:     []Rules.ValidationRule{Rules.Sometimes, Rules.LengthBetween(2, 255)},
		},
		{
			FieldName: "image",
			Rules:     []Rules.ValidationRule{Rules.Sometimes, Rules.File, Rules.Mimes("jpg", "jpeg", "png"), Rules.MaxSize(1000)},
		},
		{
			FieldName: "email",
			Rules:     []Rules.ValidationRule{Rules.Sometimes, Rules.Email, Rules.LengthBetween(2, 255), Rules.Unique("email", "users")},
		},
		{
			FieldName: "phone",
			Rules:     []Rules.ValidationRule{Rules.Sometimes, Rules.Regex(`09\d{9}`), Rules.Unique("phone", "users")},
		},
		{
			FieldName: "password",
			Rules:     []Rules.ValidationRule{Rules.Sometimes, Rules.LengthBetween(8, 16)},
		},
	})
}
