package Rules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
)

func File(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return false, "", nil, fmt.Errorf("error , File Rule : %s", err)
	}
	message, err = Validation.ErrorMessageProvider(field_name, "file", nil)
	if err != nil {
		return false, "", nil, fmt.Errorf("error , File Rule : %s", err)
	}
	if form.File[field_name] == nil {
		return false, message, nil, nil
	}
	return true, "", nil, nil
}
