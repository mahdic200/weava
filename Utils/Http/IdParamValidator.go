package Http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func IdParamValidator(c *fiber.Ctx) (int, fiber.Map) {
	id_param := c.Params("id")
	id, err := strconv.ParseUint(id_param, 10, 64)
	if err != nil || id == 0 {
		return 0, fiber.Map{
			"message": "Id must be a valid non-negative and non-zero integer",
		}
	}
	return int(id), nil
}
