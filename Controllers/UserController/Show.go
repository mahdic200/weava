package UserController

import (
	db "github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"

	"github.com/gofiber/fiber/v2"
)

func Show(c *fiber.Ctx) error {
    id := c.Params("id")
    var user Models.User
    db.DB.Select("id,name,password").Where("id=?", id).First(&user)
    if user.Id == 0 {
        return c.Status(404).JSON(fiber.Map{
            "message": "! کاربر مورد نظر یافت نشد",
        })
    }
    // user_data = make(map[string]interface {})
    return c.Status(200).JSON(fiber.Map{
        "message": "this is your user !",
        "data": user,
    })
}