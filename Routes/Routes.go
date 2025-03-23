package Routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
    /* Static file rendering */
    app.Static("/", "public")

    /* Not found response */
    app.Use("*", func(c *fiber.Ctx) error {
        return c.Status(404).JSON(fiber.Map{
            "message": "Not found",
        })
    })
}
