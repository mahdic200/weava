package UserController

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
)

func Index(c *fiber.Ctx) error {
    var users []Models.User
    db.DB.Table("users").Find(&users)
    return c.Status(200).JSON(fiber.Map{
        "data": users,
    })
}
