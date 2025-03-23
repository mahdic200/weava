package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Controllers/Admin/UserController"
	"github.com/mahdic200/weava/Validations/Admin/UserValidation"
)

func SetupRoutes(app *fiber.App) {

	adminGroup := app.Group("/admin")

	userGroup := adminGroup.Group("/user")
	userGroup.Get("/", UserController.Index).Name("admin.user.index")
	userGroup.Get("/show/:id", UserController.Show).Name("admin.user.show")
	userGroup.Post("/store", UserValidation.Store(), UserController.Store).Name("admin.user.store")

	/* Static file rendering */
	app.Static("/", "public")

	/* Not found response */
	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found",
		})
	})
}
