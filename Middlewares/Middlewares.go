package Middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Providers"
	"github.com/mahdic200/weava/Rules"
)

/*
	Our middleware for validating form-data requests , it iterates all over on your defined validators in

Validations folder, see the docs to see how to use validators
*/
func ValidationMiddleware(schema []Rules.FieldRules) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		/* This part of the code is so important , because if you don't place it , it will spit an ugly
		   error in your face :) don't trust me ? comment out this section and see this happens => "request
		   Content-Type has bad boundary or is not multipart/form-data" */
		body := c.Request().Body()
		if len(body) == 0 {
			return c.Status(400).JSON(fiber.Map{
				"message": "Empty form-data request",
			})
		}
		/* This is the main body of our middleware to loop through field rules and then start calling them
		   to see is there any errors or not , according to our previous lines of code , it's very rare to
		   encounter any technical errors , unless to tweak the rules or define yours */
		for _, field_rules := range schema {
			for _, rule := range field_rules.Rules {
				passed, message, flags, err := rule(c, field_rules.FieldName)
				if err != nil {
					return c.Status(500).JSON(fiber.Map{
						"message": Providers.ErrorProvider(err),
					})
				}
				/* The it's null flag is so important and if the Nullable rules exists in the array, which
				   it must be the first rule in the list , we simply skip all other rules in the list for
				   current field_name, rules that can issue this flag are Sometimes and Nullable */
				if passed && (flags != nil && flags.IsNull) {
					break
				} else if !passed {
					return c.Status(400).JSON(fiber.Map{
						"message": message,
					})
				}
			}
		}
		return c.Next()
	}
}
