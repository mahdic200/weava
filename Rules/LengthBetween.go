package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
)

func LengthBetween(min uint, max uint) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)
		otherKeys := make(map[string]string)
		otherKeys["{min}"] = fmt.Sprintf("%v", min)
		otherKeys["{max}"] = fmt.Sprintf("%v", max)
		message, err = Validation.ErrorMessageProvider(field_name, "lengthBetween", otherKeys)
		if err != nil {
			return false, "", nil, err
		}
		if len(value) > int(max) || len(value) < int(min) {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
