package book

import (
	"github.com/gofiber/fiber/v2"

	"github.com/copy-ninja1/library/config/database"
	"github.com/copy-ninja1/library/models"
)

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book models.Book
	db.Find(&book, id)
	return c.JSON(book)
}
