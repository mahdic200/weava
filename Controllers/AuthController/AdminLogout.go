package AuthController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Providers"
)

func AdminLogout(c *fiber.Ctx) error {
	tx := Config.DB

	if err := tx.Delete(&Models.AdminSession{}, c.Locals("token_id")).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}

	return c.Status(200).JSON(map[string]string{
		"message": "You have been logged out",
	})
}
