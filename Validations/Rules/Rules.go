package Rules

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Providers/Validation"
)

type ValidationRule func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error)


/* general validation rules */

func Nullable(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
    value := c.FormValue(field_name)
    flags = make(map[string]bool)
    flags["itsnull"] = true
    if value == "" {
        return true, "", flags, nil
    }
    return true, "", nil, nil
}

func Sometimes(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
    form , err := c.MultipartForm()
    if err != nil {
        return false, "", nil, err
    }
    value := form.Value[field_name]
    message, err = Validation.ErrorMessageProvider(field_name, "sometimes", nil)
    if err != nil {
        return false, "", nil, err
    }
    fmt.Printf("value : %#v\n", value)
    flags = make(map[string]bool)
    flags["itsnull"] = true
    if value == nil {
        return true, "", flags, nil
    } else if value[0] == "" {
        return false, message, nil, nil
    }
    return true, "", nil, nil
}

func Required(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
    form, err := c.MultipartForm()
    if err != nil {
        return false, "", nil, err
    }
    value := form.Value[field_name]
    message, err = Validation.ErrorMessageProvider(field_name, "required", nil)
    if err != nil {
        return false, "", nil, err
    }
    if value == nil || value[0] == "" {
        return false, message, nil, nil
    }
    return true, "", nil, nil
}

func Unique(column string, table string) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value := c.FormValue(field_name)
        message, err = Validation.ErrorMessageProvider(field_name, "unique", nil)
        if err != nil {
            return false, "", nil, err
        }
        var record map[string]any
        Config.DB.Table(table).Select(column).Where(column + " = ?", value).Find(&record)
        if record != nil {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

/* foreign keys */

func Exists(column string, table string) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value := c.FormValue(field_name)
        message, err = Validation.ErrorMessageProvider(field_name, "exists", nil)
        if err != nil {
            return false, "", nil, err
        }
        var record map[string]any
        Config.DB.Table(table).Select(column).Where(column + " = ?", value).Find(&record)
        if record == nil {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

func InArray(array []string) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
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

func NotInArray(array []string) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value := c.FormValue(field_name)
        otherKeys := make(map[string]string)
        otherKeys["{options}"] = strings.Join(array, " , ")
        message, err = Validation.ErrorMessageProvider(field_name, "notInArray", otherKeys)
        if err != nil {
            return false, "", nil, err
        }
        if slices.Contains(array, value) {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

/* string validation rules */

func Email(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
    value := c.FormValue(field_name)
    r := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    message, err = Validation.ErrorMessageProvider(field_name, "email", nil)
    if err != nil {
        return false, "", nil, err
    }
    fmt.Printf("%#v\n", value)
    fmt.Printf("%#v\n", r.MatchString(value))
    if !r.MatchString(value) {
        return false, message, nil, nil
    }
    return true, "", nil, nil
}
