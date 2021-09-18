package book

import (
	"github.com/gofiber/fiber/v2"

	"github.com/copy-ninja1/library/config/database"
	"github.com/copy-ninja1/library/models"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []models.Book
	db.Find(&books)
	return c.JSON(books)
}
