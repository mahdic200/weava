package Http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func BodyParser(c *fiber.Ctx) (map[string]string, error) {
	parsed := make(map[string]string)
	form, err := c.MultipartForm()
	if err != nil {
		return nil, fmt.Errorf("BodyParser helper : %s", err.Error())
	}
	for key, value := range form.Value {
		parsed[key] = value[0]
	}
	return parsed, nil
}
