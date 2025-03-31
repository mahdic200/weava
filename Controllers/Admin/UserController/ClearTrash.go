package UserController

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Providers"
	"github.com/mahdic200/weava/Providers/Response"
	"github.com/mahdic200/weava/Utils/File"
)

func ClearTrash(c *fiber.Ctx) error {
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

	var users []Models.User
	/* This part is super important , we can't take the risk of doing some operation
	on many number of records simultaneously , unless we take the risk of downing our
	server and it means downtime or unexpected problems such as memory leak , or high
	pressure on CPU , because its not just about deleting records , it's about removing
	files and that means working with IO , and IO operations are always kind of pain
	so with limiting the number of records , we prevent such problems */
	err := tx.Table("users").Where("deleted_at IS NOT NULL").Limit(10).Scan(&users).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}

	for _, user := range users {
		/* old_file is absolute paths of files and theres no need of using
		File.PublicPath() function for it */
		old_file := File.PublicPath(user.Image)
		if err := tx.Exec("DELETE FROM users WHERE id = ?", user.Id).Error; err != nil {
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
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	message, _ := Response.Message("users", "clearedTrash")
	return c.Status(200).JSON(fiber.Map{
		"message": message,
	})
}
