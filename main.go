package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/copy-ninja1/library/config/database"
	"github.com/copy-ninja1/library/config/routes"
	"github.com/copy-ninja1/library/config/start"
)

func main() {

	app := fiber.New()
	// start.OpenDataBase()

	start.InitializeDataBase()

	sqlDB, _ := database.DBConn.DB()
	defer sqlDB.Close()
	routes.Routes_v1(app)
	// setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
