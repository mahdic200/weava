package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
)

func MaxLength(max uint) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)
		otherKeys := make(map[string]string)
		otherKeys["{max}"] = fmt.Sprintf("%v", max)
		message, err = Validation.ErrorMessageProvider(field_name, "maxLength", otherKeys)
		if err != nil {
			return false, "", nil, err
		}
		if len(value) > int(max) {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
