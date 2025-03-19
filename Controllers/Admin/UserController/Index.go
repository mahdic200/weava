package UserController

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Models/User"
	"github.com/mahdic200/weava/Resources/UserResource"
	"github.com/mahdic200/weava/Utils/Http"
)

func Index(c *fiber.Ctx) error {
    var users []Models.User
    tx := db.DB.Table("users")
    var metadata Http.PaginationMetadata
    if Http.GetQueryArg(c, "page") != nil {
        tx, metadata = User.Paginate(tx, c)
        tx.Find(&users)
        return c.Status(200).JSON(fiber.Map{
            "data": UserResource.Collection(users),
            "metadata": metadata,
        })
    }
    tx.Find(&users)
    return c.Status(200).JSON(fiber.Map{
        "data": UserResource.Collection(users),
    })
}
