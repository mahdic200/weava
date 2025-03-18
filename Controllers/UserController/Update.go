package UserController

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/mahdic200/weava/Config"
)


func Update(c *fiber.Ctx) error {
    // id := c.Params("id")
    // var user models.User

    db.DB.Select("id,name")
    return c.Status(200).JSON(fiber.Map{
        "message": "",
    })
}
