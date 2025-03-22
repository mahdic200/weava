package Rules

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Providers/Validation"
	"github.com/mahdic200/weava/Services/FileService"
)

type ValidationRule func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error)

type FieldRules struct {
    FieldName string
    Rules []ValidationRule
}


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

func Length(length uint) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value := c.FormValue(field_name)
        otherKeys := make(map[string]string)
        otherKeys["{length}"] = fmt.Sprintf("%v", length)
        message, err = Validation.ErrorMessageProvider(field_name, "length", otherKeys)
        if err != nil {
            return false, "", nil, err
        }
        if len(value) != int(length) {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

func MinLength(min uint) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value := c.FormValue(field_name)
        otherKeys := make(map[string]string)
        otherKeys["{min}"] = fmt.Sprintf("%v", min)
        message, err = Validation.ErrorMessageProvider(field_name, "minLength", otherKeys)
        if err != nil {
            return false, "", nil, err
        }
        if len(value) < int(min) {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

func MaxLength(max uint) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value := c.FormValue(field_name)
        otherKeys := make(map[string]string)
        otherKeys["{max}"] = fmt.Sprintf("%v", max)
        message, err = Validation.ErrorMessageProvider(field_name, "maxLength", otherKeys)
        if err != nil {
            return false, "", nil, err
        }
        if len(value) > int(max) {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

func LengthBetween(min uint, max uint) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value := c.FormValue(field_name)
        otherKeys := make(map[string]string)
        otherKeys["{min}"] = fmt.Sprintf("%v", min)
        otherKeys["{max}"] = fmt.Sprintf("%v", max)
        message, err = Validation.ErrorMessageProvider(field_name, "lengthBetween", otherKeys)
        if err != nil {
            return false, "", nil, err
        }
        if len(value) > int(max) || len(value) < int(min) {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

func Email(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
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

func Regex(regex string) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value := c.FormValue(field_name)
        r := regexp.MustCompile(regex)
        message, err = Validation.ErrorMessageProvider(field_name, "regex", nil)
        if err != nil {
            return false, "", nil, err
        }
        if !r.MatchString(value) {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

/* file validation */

func File(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
    form, err := c.MultipartForm()
    if err != nil {
        return false, "", nil, fmt.Errorf("Error , File Rule : %s", err)
    }
    message, err = Validation.ErrorMessageProvider(field_name, "file", nil)
    if form.File[field_name] == nil {
        return false, message, nil, nil
    }
    return true, "", nil, nil
}

func Mimes(array ...string) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value, err := c.FormFile(field_name)
        if err != nil {
            return false, "", nil, fmt.Errorf("Error , Mines Rule : %s", err)
        }
        otherKeys := make(map[string]string)
        otherKeys["{mimes}"] = strings.Join(array, " , ")
        message, err = Validation.ErrorMessageProvider(field_name, "mimes", otherKeys)
        if err != nil {
            return false, "", nil, err
        }
        if !slices.Contains(array, FileService.GetFileExtension(value.Filename)) {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

func FileSize(size uint) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value, err := c.FormFile(field_name)
        if err != nil {
            return false, "", nil, fmt.Errorf("Error , FileSize Rule : %s", err)
        }
        otherKeys := make(map[string]string)
        otherKeys["{size}"] = fmt.Sprintf("%v", size)
        message, err = Validation.ErrorMessageProvider(field_name, "fileSize", otherKeys)
        if err != nil {
            return false, "", nil, err
        }
        if int(value.Size / 1000) != int(size) {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

func MinSize(min uint) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value, err := c.FormFile(field_name)
        if err != nil {
            return false, "", nil, fmt.Errorf("Error , MinSize Rule : %s", err)
        }
        otherKeys := make(map[string]string)
        otherKeys["{min}"] = fmt.Sprintf("%v", min)
        message, err = Validation.ErrorMessageProvider(field_name, "minSize", otherKeys)
        if err != nil {
            return false, "", nil, err
        }
        if int(value.Size / 1000) < int(min) {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

func MaxSize(max uint) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value, err := c.FormFile(field_name)
        if err != nil {
            return false, "", nil, fmt.Errorf("Error , MaxSize Rule : %s", err)
        }
        otherKeys := make(map[string]string)
        otherKeys["{max}"] = fmt.Sprintf("%v", max)
        message, err = Validation.ErrorMessageProvider(field_name, "maxSize", otherKeys)
        if err != nil {
            return false, "", nil, err
        }
        if int(value.Size / 1000) > int(max) {
            return false, message, nil, nil
        }
        return true, "", nil, nil
    }
}

func SizeBetween(min uint, max uint) ValidationRule {
    return func(c *fiber.Ctx, field_name string) (passed bool, message string, flags map[string]bool, err error) {
        value, err := c.FormFile(field_name)
        if err != nil {
            return false, "", nil, fmt.Errorf("Error , SizeBetween Rule : %s", err)
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
