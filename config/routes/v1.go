package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/copy-ninja1/library/controllers/book"
)

func Routes_v1(app *fiber.App) {
	api := app.Group("/api") // /api
	v1 := api.Group("/v1")
	v1.Get("/book", book.GetBooks)
	v1.Get("/book/:id", book.GetBook)
	v1.Post("/book", book.NewBook)
	v1.Delete("/book/:id", book.DeleteBook)

}
