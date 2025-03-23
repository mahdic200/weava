package Rules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
)

func Required(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return false, "", nil, err
	}
	value := form.Value[field_name]
	file := form.File[field_name]
	message, err = Validation.ErrorMessageProvider(field_name, "required", nil)
	if err != nil {
		return false, "", nil, err
	}
	if (value == nil || value[0] == "") && (file == nil) {
		return false, message, nil, nil
	}
	return true, "", nil, nil
}
