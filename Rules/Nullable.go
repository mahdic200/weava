package Rules

import "github.com/gofiber/fiber/v2"

func Nullable(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
	value := c.FormValue(field_name)
	flags = &Flags{ IsNull: true }
	if value == "" {
		return true, "", flags, nil
	}
	return true, "", nil, nil
}
