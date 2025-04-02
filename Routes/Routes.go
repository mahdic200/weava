package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Controllers/Admin/UserController"
	"github.com/mahdic200/weava/Controllers/AuthController"
	"github.com/mahdic200/weava/Middlewares"
	"github.com/mahdic200/weava/Validations/Admin/UserValidation"
	"github.com/mahdic200/weava/Validations/Auth"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/login", Auth.Login(), AuthController.Login)
	app.Post("/register", Auth.Register(), AuthController.Register)

	adminGroup := app.Group("/admin", Middlewares.AuthMiddleware)

	userGroup := adminGroup.Group("/user")
	userGroup.Get("/", UserController.Index).Name("admin.user.index")
	userGroup.Get("/show/:id", UserController.Show).Name("admin.user.show")
	userGroup.Post("/store", UserValidation.Store(), UserController.Store).Name("admin.user.store")
	userGroup.Post("/update/:id", UserValidation.Update(), UserController.Update).Name("admin.user.update")
	userGroup.Post("/delete/:id", UserController.Delete).Name("admin.user.delete")
	userGroup.Get("/restore/:id", UserController.Restore).Name("admin.user.restore")
	userGroup.Get("/trash", UserController.Trash).Name("admin.user.trash")
	userGroup.Post("/force-delete/:id", UserController.ForceDelete).Name("admin.user.force-delete")
	userGroup.Post("/clear-trash", UserController.ClearTrash).Name("admin.user.clear-trash")

	/* Static file rendering */
	app.Static("/", "public")

	/* Not found response */
	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Route Not found",
		})
	})
}
