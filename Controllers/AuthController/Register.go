package AuthController

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Models/User"
	"github.com/mahdic200/weava/Providers"
	"github.com/mahdic200/weava/Resources/UserResource"
	"github.com/mahdic200/weava/Utils"
	"github.com/mahdic200/weava/Utils/Http"
)

type RRS struct {
	Token       string            `json:"token"`
	Expire_time time.Time         `json:"expire_time"`
	User        UserResource.User `json:"user"`
}

func Register(c *fiber.Ctx) error {

	tx := Config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(tx.Error),
		})
	}
	/* Parsing body */
	data, err := Http.BodyParser(c)
	if err != nil {
		tx.Rollback()
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}
	data["password"], err = Utils.GenerateHashPassword(data["password"])
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}

	if err := User.Create(tx, data).Error; err != nil {
		tx.Rollback()
		fmt.Printf("%#v\n", data)
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	var user Models.User
	tx.Table("users").Where("email = ?", data["email"]).Find(&user)

	remember_me := false
	if data["remember_me"] == "true" {
		remember_me = true
	}
	now := time.Now()
	new_token := Models.Session{
		User_id:    user.Id,
		Created_at: &now,
	}
	if err := tx.Create(&new_token).Error; err != nil {
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
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"message": Providers.ErrorProvider(err),
		})
	}
	return c.Status(200).JSON(RRS{
		Token:       token,
		Expire_time: expire_time,
		User:        UserResource.Single(user),
	})
}
