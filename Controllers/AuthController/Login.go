package AuthController

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Providers"
	"github.com/mahdic200/weava/Providers/Response"
	"github.com/mahdic200/weava/Utils"
	"github.com/mahdic200/weava/Utils/Http"
)

type ResponseStruct struct {
	Token       string    `json:"token"`
	Expire_time time.Time `json:"expire_time"`
}

func Login(c *fiber.Ctx) error {
	tx := Config.DB
	/* Parsing body */
	data, err := Http.BodyParser(c)
	if err != nil {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}
	var user Models.User
	tx.Model(&user).Where("deleted_at IS NULL").Where("email = ?", data["email"]).Find(&user)
	message, _ := Response.Message("user", "notFound")
	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": message,
		})
	}
	message, _ = Response.Message("email", "wrongCredentials")
	if err := Utils.VerifyPassword(data["password"], user.Password); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": message,
		})
	}
	remember_me := false
	if data["remember_me"] == "true" {
		remember_me = true
	}
	now := time.Now()
	new_token := Models.Session{
		User_id:    user.Id,
		Created_at: &now,
	}
	session := tx.Create(&new_token)
	if session.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	token, expire_time, err := Utils.CreateToken(int64(new_token.Id), "sessions", remember_me)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	if session.Update("token_string", token).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	return c.Status(200).JSON(ResponseStruct{
		Token:       token,
		Expire_time: expire_time,
	})
}
