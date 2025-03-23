package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
)

func MaxSize(max uint) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value, err := c.FormFile(field_name)
		if err != nil {
			return false, "", nil, fmt.Errorf("error , MaxSize Rule : %s", err)
		}
		otherKeys := make(map[string]string)
		otherKeys["{max}"] = fmt.Sprintf("%v", max)
		message, err = Validation.ErrorMessageProvider(field_name, "maxSize", otherKeys)
		if err != nil {
			return false, "", nil, err
		}
		if int(value.Size/1000) > int(max) {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
