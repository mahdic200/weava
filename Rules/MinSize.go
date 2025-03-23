package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
)

func MinSize(min uint) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value, err := c.FormFile(field_name)
		if err != nil {
			return false, "", nil, fmt.Errorf("error , MinSize Rule : %s", err)
		}
		otherKeys := make(map[string]string)
		otherKeys["{min}"] = fmt.Sprintf("%v", min)
		message, err = Validation.ErrorMessageProvider(field_name, "minSize", otherKeys)
		if err != nil {
			return false, "", nil, err
		}
		if int(value.Size/1000) < int(min) {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
