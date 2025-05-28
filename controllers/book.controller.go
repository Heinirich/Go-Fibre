package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Heinirich/gorm-mysql/models"
	"github.com/Heinirich/gorm-mysql/protocol"
	"github.com/Heinirich/gorm-mysql/repository"
	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	Repo repository.BookRepository
}

func (c *BookController) CreateBook(ctx *fiber.Ctx) error {
	var protoBook protocol.Book
	if err := ctx.BodyParser(&protoBook); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	book := models.FromProto(&protoBook)
	if err := c.Repo.Create(book); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "DB error"})
	}

	return ctx.Status(201).JSON(book.ToProto())
}

func (c *BookController) GetBooks(ctx *fiber.Ctx) error {
	time_start := time.Now()
	books, err := c.Repo.FindAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "DB error"})
	}

	var protoBooks []*protocol.Book
	for _, b := range books {
		protoBooks = append(protoBooks, b.ToProto())
	}

	timeTaken := time.Since(time_start)
	fmt.Println(timeTaken)

	return ctx.JSON(protoBooks)
}

func (c *BookController) GetBook(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	book, err := c.Repo.FindByID(uint(id))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	return ctx.JSON(book.ToProto())
}

func (c *BookController) UpdateBook(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var protoBook protocol.Book
	if err := ctx.BodyParser(&protoBook); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	book := models.FromProto(&protoBook)
	book.ID = uint(id) // override ID to match URL

	if err := c.Repo.Update(book); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "DB error"})
	}

	return ctx.JSON(book.ToProto())
}

func (c *BookController) DeleteBook(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := c.Repo.Delete(uint(id)); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "DB error"})
	}

	return ctx.SendStatus(200)
}
