package UserValidator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Middlewares"
	"github.com/mahdic200/weava/Validations/Rules"
)

// map[string]([](Rules.ValidationRule))
func Store() func(c *fiber.Ctx) error {
    rules := make(map[string]([](Rules.ValidationRule)))
    rules["first_name"] = [](Rules.ValidationRule){ Rules.Required }
    rules["last_name"] = [](Rules.ValidationRule){ Rules.Nullable }
    rules["email"] = [](Rules.ValidationRule){ Rules.Email }
    return Middlewares.ValidationMiddleware(rules)
}

func UpdateValidator() [](func()) {
    return [](func()){
        
    }
}
