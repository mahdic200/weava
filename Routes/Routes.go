package Routes

import (
	"github.com/mahdic200/weava/Controllers/AuthController"
	"github.com/mahdic200/weava/Controllers/Admin/UserController"
	"github.com/mahdic200/weava/Validations/Admin/UserValidator"

	// "github.com/mahdic200/weava/Middlewares/Auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

    app.Post("/login", AuthController.Login).Name("app.user.index")
    // adminGroup := app.Group("/admin", Auth.AuthMiddleware)
    adminGroup := app.Group("/admin")

    userGroup := adminGroup.Group("user")
    userGroup.Get("/", UserController.Index).Name("app.user.index")
    userGroup.Get("/show/:id", UserController.Show).Name("app.user.show")
    userGroup.Post("/store", UserValidator.Store(), UserController.Store).Name("app.user.store")

    /* Static file rendering */
    app.Static("/", "public")

    /* Not found reponse */
    app.Use("*", func(c *fiber.Ctx) error {
        return c.Status(404).JSON(fiber.Map{
            "message": "Not found",
        })
    })
}


