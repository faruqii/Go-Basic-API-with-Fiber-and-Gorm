package main

import (
	"gowebservices/book"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	bookHandler := book.NewHandler(db)

	app := fiber.New()

	app.Get("/books", bookHandler.GetBooks)
	app.Post("/books", bookHandler.NewBooks)

	app.Listen(":3000")
}