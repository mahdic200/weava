package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Providers"
	"github.com/mahdic200/weava/Providers/Response"
	"github.com/mahdic200/weava/Utils/Http"
)

func Delete(c *fiber.Ctx) error {
	var id int
	if param, err := Http.IdParamValidator(c); err != nil {
		return c.Status(400).JSON(err)
	} else {
		id = param
	}

	tx := Config.DB

	var user Models.User
	tx.Table("users").Where("id = ?", id).Where("deleted_at IS NULL").Find(&user)
	if user.Id == 0 {
		message, _ := Response.Message("user", "notFound")
		return c.Status(404).JSON(fiber.Map{
			"message": message,
		})
	}

	if err := tx.Exec("UPDATE users SET deleted_at = NOW() WHERE id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}

	message, _ := Response.Message("user", "trashed")
	return c.Status(200).JSON(fiber.Map{
		"message": message,
	})
}
