package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyort/go-fiber-api/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json.title`
	Author string `json.author`
	Rating int    `json.rating`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}
func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&book)

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	change := false
	var book Book
	db.Find(&book, id)
	newbook := new(Book)
	if err := c.BodyParser(newbook); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	if book.Title != newbook.Title {
		book.Title = newbook.Title
		change = true
	}
	if book.Author != newbook.Author {
		book.Author = newbook.Author
		change = true
	}
	if book.Rating != newbook.Rating {
		book.Rating = newbook.Rating
		change = true
	}

	if change {
		db.Save(&book)
		return c.SendString("Book successfully updated")
	} else {
		return c.SendString("No changes made")
	}

}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No book found with given ID")
	}
	db.Delete(&book)

	return c.SendString("Book successfully deleted")
}
