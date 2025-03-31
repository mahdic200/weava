package UserController

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Providers"
	"github.com/mahdic200/weava/Providers/Response"
	"github.com/mahdic200/weava/Utils/File"
	"github.com/mahdic200/weava/Utils/Http"
)

func ForceDelete(c *fiber.Ctx) error {
	var id int
	if param, err := Http.IdParamValidator(c); err != nil {
		return c.Status(400).JSON(err)
	} else {
		id = param
	}

	tx := Config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(tx.Error),
		})
	}

	var user Models.User
	tx.Table("users").Where("deleted_at IS NOT NULL").Where("id = ?", id).Find(&user)
	if user.Id == 0 {
		message, _ := Response.Message("user", "notFound")
		return c.Status(404).JSON(fiber.Map{
			"message": message,
		})
	}

	/* old_file is absolute paths of files and theres no need of using
	File.PublicPath() function for it */
	old_file := File.PublicPath(user.Image)

	if err := tx.Exec("DELETE FROM users WHERE id = ?", id).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}

	if old_file != "" && File.Exists(old_file) {
		if err := os.Remove(old_file); err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"message": Providers.ErrorProvider(err),
			})
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	message, _ := Response.Message("user", "deleted")
	return c.Status(200).JSON(fiber.Map{
		"message": message,
	})
}
