package Rules

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
)

func Email(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
	value := c.FormValue(field_name)
	r := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	message, err = Validation.ErrorMessageProvider(field_name, "email", nil)
	if err != nil {
		return false, "", nil, err
	}
	if !r.MatchString(value) {
		return false, message, nil, nil
	}
	return true, "", nil, nil
}
