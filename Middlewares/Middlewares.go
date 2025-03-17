package Middlewares

import (
	// "encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Validations/Rules"
)

func ValidationMiddleware(schema map[string]([](Rules.ValidationRule))) (func(c *fiber.Ctx) error) {
    return func(c *fiber.Ctx) error {
        for field_name, rules := range schema {
            for _, rule := range rules {
                passed, message, flags, err := rule(c, field_name)
                if err != nil {
                    return c.Status(500).JSON(fiber.Map{
                        "message": "Internal server error !",
                    })
                }
                if passed && (flags["itsnull"]) {
                    break
                } else if (!passed) {
                    return c.Status(400).JSON(fiber.Map{
                        "message": message,
                    })
                }
            }
        }
        return c.Next()
    }
}
