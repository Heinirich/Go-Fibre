package routes

import (
	"github.com/Heinirich/gorm-mysql/config"
	"github.com/Heinirich/gorm-mysql/controllers"
	"github.com/Heinirich/gorm-mysql/repository"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookRoutes(app *fiber.App) {
	bookRepo := repository.NewBookRepository(config.DB)

	controller := &controllers.BookController{
		Repo: bookRepo,
	}

	api := app.Group("/api")
	api.Post("/books", controller.CreateBook)
	api.Get("/books", controller.GetBooks)
	api.Get("/books/:id", controller.GetBook)
	api.Put("/books/:id", controller.UpdateBook)
	api.Delete("/books/:id", controller.DeleteBook)
}
