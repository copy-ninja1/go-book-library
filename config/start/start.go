package start

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/copy-ninja1/library/config/database"
	"github.com/copy-ninja1/library/models"
)

func InitializeDataBase() {
	var err error

	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{}) //gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&models.Book{})
	fmt.Println("Database migrated")

}
