package UserController

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	db "github.com/mahdic200/weava/Config"
	models "github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Services/FileService"
	"github.com/mahdic200/weava/Utils"
)

func Store(c *fiber.Ctx) error {
    tx := db.DB.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()
    if tx.Error != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Internal server error",
        })
    }
    data := new(models.User)
    err := c.BodyParser(data)
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

    fs := FileService.NewFileService(file)

    if err := fs.SaveToPublic("uploads", "images", "user-profiles"); err != nil {
        tx.Rollback()
        return c.Status(500).JSON(fiber.Map{
            "message": "Failed to save file !",
        })
    }
    data.Image = fs.GetRelativePath()

    data.Password, err = Utils.GenerateHashPassword(data.Password)
    if err != nil {
        tx.Rollback()
        return c.Status(500).JSON(fiber.Map{
            "message": "Internal server error",
        })
    }
    if err := tx.Exec("INSERT INTO users (first_name, last_name, email, phone, image, password, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)", data.First_name, data.Last_name, data.Email, data.Phone, data.Image, data.Password, time.Time{}).Error; err != nil {
        tx.Rollback()
        os.Remove(fs.GetFinalPath())
        return c.Status(500).JSON(fiber.Map{
            "message": "Internal server error !",
        })
    }
    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        os.Remove(fs.GetFinalPath())
        return c.Status(500).JSON(fiber.Map{
            "message": "Internal server error",
        })
    }
    return c.Status(200).JSON(fiber.Map{
        "message": "user created successfully !",
    })
}
