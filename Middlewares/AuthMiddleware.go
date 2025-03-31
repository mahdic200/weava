package Middlewares

import (
	"strings"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Providers"
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
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token",
		})
	} else if table_name != "sessions" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tx := Config.DB

	var session Models.Session
	if err := tx.Table("sessions").Where("id = ?", entity_id).Find(&session).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	if session.Id == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var user Models.User
	if err := tx.Table("users").Where("deleted_at IS NULL").Where("id = ?", session.User_id).Find(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}

	c.Locals("user", user)

	return c.Next()
}
