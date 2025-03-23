package Rules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
)

func Sometimes(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return false, "", nil, err
	}
	value := form.Value[field_name]
	file := form.File[field_name]
	message, err = Validation.ErrorMessageProvider(field_name, "sometimes", nil)
	if err != nil {
		return false, "", nil, err
	}
	flags = &Flags{IsNull: true}
	if value == nil && file == nil {
		return true, "", flags, nil
	} else if value != nil && value[0] == "" {
		return false, message, nil, nil
	}
	return true, "", nil, nil
}
