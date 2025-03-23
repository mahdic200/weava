package Rules

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
)

func InArray(array []string) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)
		otherKeys := make(map[string]string)
		otherKeys["{options}"] = strings.Join(array, " , ")
		message, err = Validation.ErrorMessageProvider(field_name, "inArray", otherKeys)
		if err != nil {
			return false, "", nil, err
		}
		if !slices.Contains(array, value) {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
