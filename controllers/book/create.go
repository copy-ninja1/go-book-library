package book

import (
	"github.com/gofiber/fiber/v2"

	"github.com/copy-ninja1/library/config/database"
	"github.com/copy-ninja1/library/models"
)

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&book)
	return c.JSON(book)
}
