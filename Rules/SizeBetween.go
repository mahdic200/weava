package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
)

func SizeBetween(min uint, max uint) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value, err := c.FormFile(field_name)
		if err != nil {
			return false, "", nil, fmt.Errorf("error , SizeBetween Rule : %s", err)
		}
		otherKeys := make(map[string]string)
		otherKeys["{min}"] = fmt.Sprintf("%v", min)
		otherKeys["{max}"] = fmt.Sprintf("%v", max)
		message, err = Validation.ErrorMessageProvider(field_name, "sizeBetween", otherKeys)
		if err != nil {
			return false, "", nil, err
		}
		file_size := value.Size / 1000
		if int(file_size) > int(max) || int(file_size) < int(min) {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
