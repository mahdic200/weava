package UserController

import (
	"crypto/rand"
	"fmt"
	"strings"

	db "github.com/mahdic200/weava/Config"
	models "github.com/mahdic200/weava/Models"

	"time"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
    return c.Status(200).JSON(fiber.Map{
        "message": "hello pretty !",
    })
}

func Store(c *fiber.Ctx) error {
    tx := db.DB.Begin()
    data := new(models.User)
    err := c.BodyParser(data)
    if err != nil {
        fmt.Print(err)
        return c.Status(400).JSON(fiber.Map{
            "message": "Invalid data",
        })
    }
    file, err := c.FormFile("image")
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Failed to get the file!",
        })
    }

    fileExtension := strings.Split(file.Filename, ".")[1]
    newName := rand.Text()
    fmt.Println(fileExtension)
    filePath := fmt.Sprintf("public/%s.%s", newName, fileExtension)
    if err := c.SaveFile(file, filePath); err != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Failed to save file !",
            "err": err,
        })
    }

    user := models.User {
        First_name: data.First_name,
        Password: data.Password,
        Created_at: time.Time{},
    }
    tx.Create(&user)
    return c.Status(200).JSON(fiber.Map{
        "message": "user created successfully !",
    })
}

func Show(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    db.DB.Select("id,name,password").Where("id=?", id).First(&user)
    if user.Id == 0 {
        return c.Status(404).JSON(fiber.Map{
            "message": "! کاربر مورد نظر یافت نشد",
        })
    }
    // user_data = make(map[string]interface {})
    return c.Status(200).JSON(fiber.Map{
        "message": "this is your user !",
        "data": user,
    })
}

func Update(c *fiber.Ctx) error {
    // id := c.Params("id")
    // var user models.User

    db.DB.Select("id,name")
    return c.Status(200).JSON(fiber.Map{
        "message": "",
    })
}
