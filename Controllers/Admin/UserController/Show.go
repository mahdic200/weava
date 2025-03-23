package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Models/User"
	"github.com/mahdic200/weava/Providers/Response"
	"github.com/mahdic200/weava/Resources/UserResource"
	"github.com/mahdic200/weava/Utils/Http"
)

func Show(c *fiber.Ctx) error {
	id, id_err := Http.IdParamValidator(c)
	if id_err != nil {
		return c.Status(200).JSON(id_err)
	}
	var user Models.User
	User.Find(Config.DB, id, &user)
	message, _ := Response.Message("user", "notFound")
	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": message,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": UserResource.Single(user),
	})
}
