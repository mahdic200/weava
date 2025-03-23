package UserController

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models/User"
	"github.com/mahdic200/weava/Providers"
	"github.com/mahdic200/weava/Services/FileService"
	"github.com/mahdic200/weava/Utils"
	"github.com/mahdic200/weava/Utils/Http"
)

func Store(c *fiber.Ctx) error {
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
	/* Parsing body */
	data, err := Http.BodyParser(c)
	if err != nil {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}
	file, err := c.FormFile("image")
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get the file!",
		})
	}

	fs := FileService.New(file)

	if err := fs.SaveToPublic("uploads", "images", "user-profiles"); err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to save file !",
		})
	}
	data["image"] = fs.GetRelativePath()

	data["password"], err = Utils.GenerateHashPassword(data["password"])
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	if err := User.Create(tx, data).Error; err != nil {
		tx.Rollback()
		os.Remove(fs.GetFinalPath())
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		os.Remove(fs.GetFinalPath())
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "User created successfully !",
	})
}
