package book

import (
	"github.com/gofiber/fiber/v2"

	"github.com/copy-ninja1/library/config/database"
	"github.com/copy-ninja1/library/models"
)

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book models.Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(404).Send([]byte("No book found with given ID"))
		return nil
	}
	db.Delete(&book)
	return c.SendString("Book deleted successfully")
}
