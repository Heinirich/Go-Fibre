package main

import (
	"github.com/Heinirich/gorm-mysql/config"
	"github.com/Heinirich/gorm-mysql/models"
	"github.com/Heinirich/gorm-mysql/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(
		fiber.Config{
			Prefork: true,
		})

	config.ConnectDB()
	config.DB.AutoMigrate(&models.BookModel{})

	routes.RegisterBookRoutes(app)

	app.Listen(":3000")
}
