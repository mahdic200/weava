package Rules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Providers/Validation"
)

func Exists(column string, table string) ValidationRule {
	return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags *Flags, err error) {
		value := c.FormValue(field_name)
		message, err = Validation.ErrorMessageProvider(field_name, "exists", nil)
		if err != nil {
			return false, "", nil, err
		}
		var record map[string]any
		Config.DB.Table(table).Select(column).Where(column+" = ?", value).Find(&record)
		if record == nil {
			return false, message, nil, nil
		}
		return true, "", nil, nil
	}
}
