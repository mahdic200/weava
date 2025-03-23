package Rules

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers/Validation"
	"github.com/mahdic200/weava/Services/FileService"
)

func Mimes(array ...string) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value, err := c.FormFile(field_name)
		if err != nil {
			return false, "", nil, fmt.Errorf("error , Mines Rule : %s", err)
		}
		otherKeys := make(map[string]string)
		otherKeys["{options}"] = strings.Join(array, " , ")
		message, err = Validation.ErrorMessageProvider(field_name, "mimes", otherKeys)
		if err != nil {
			return false, "", nil, err
		}
		if !slices.Contains(array, strings.ToLower(FileService.GetFileExtension(value.Filename))) {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
