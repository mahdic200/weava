package Auth

import (
	"strings"

	"github.com/mahdic200/weava/Utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tokenSplit := strings.Split(token, " ")

	if len(tokenSplit) != 2 || strings.ToLower(tokenSplit[0]) != "bearer" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token type, expected bearer",
		})
	}

	entity_id, table_name, err := Utils.VerifyToken(tokenSplit[1])
	if err != nil || table_name != "users" {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Locals("user_id", entity_id)

	return c.Next()
}
