package UserController

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Models/User"
	"github.com/mahdic200/weava/Providers"
	"github.com/mahdic200/weava/Providers/Response"
	"github.com/mahdic200/weava/Services/FileService"
	"github.com/mahdic200/weava/Utils"
	"github.com/mahdic200/weava/Utils/File"
	"github.com/mahdic200/weava/Utils/Http"
)

func Update(c *fiber.Ctx) error {
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
	User.Find(tx, id, &user)
	if user.Id == 0 {
		message, _ := Response.Message("user", "notFound")
		return c.Status(404).JSON(fiber.Map{
			"message": message,
		})
	}

	/* Parsing body */
	/* We immune ourselves of any kind of bad-data-type-format from client's request here
	   so in the next section for file uploading , we are sure that just one type of error
	   might happen which is file not found error */
	data, err := Http.BodyParser(c)
	if err != nil {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	/* old_file and new_file are absolute paths of files and theres no need of using
	File.PublicPath() function for them */
	old_file := File.PublicPath(user.Image)
	new_file := ""
	file, err := c.FormFile("image")
	if file != nil && err == nil {
		fs := FileService.New(file)
		if err := fs.SaveToPublic("uploads", "images", "user-profiles"); err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to save file !",
			})
		}
		new_file = fs.GetFinalPath()
		data["image"] = fs.GetRelativePath()
	}

	if data["password"] != "" {
		data["password"], err = Utils.GenerateHashPassword(data["password"])

		if err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"message": Providers.ErrorProvider(err),
			})
		}
	}

	if err := User.Update(id, tx, data).Error; err != nil {
		tx.Rollback()
		if new_file != "" {
			os.Remove(new_file)
		}
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		if new_file != "" {
			os.Remove(new_file)
		}
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	if new_file != "" && File.Exists(old_file) {
		os.Remove(old_file)
	}
	User.Find(Config.DB, id, &user)
	message, _ := Response.Message("user", "updated")
	return c.Status(200).JSON(fiber.Map{
		"message": message,
	})
}
