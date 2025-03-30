package AuthController

import "github.com/gofiber/fiber/v2"

func AdminLogin(c *fiber.Ctx) error {
	// !TODO
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello from AdminLogin",
	})
}
