package Rules

import (
	"github.com/gofiber/fiber/v2"
)

type ValidationRule func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error)

type FieldRules struct {
	FieldName string
	Rules     []ValidationRule
}

type Flags struct {
	IsNull bool
}
