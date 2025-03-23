package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Controllers/Admin/UserController"
)

func SetupRoutes(app *fiber.App) {

    adminGroup := app.Group("/admin")

    userGroup := adminGroup.Group("/user")
    userGroup.Get("/", UserController.Index).Name("admin.user.index")

    /* Static file rendering */
    app.Static("/", "public")

    /* Not found response */
    app.Use("*", func(c *fiber.Ctx) error {
        return c.Status(404).JSON(fiber.Map{
            "message": "Not found",
        })
    })
}
