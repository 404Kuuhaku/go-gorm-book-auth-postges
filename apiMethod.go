package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getAllBooksHandler(c *fiber.Ctx) error {
	books := getBooks(db)
	return c.JSON(books)
}

func getBookHandler(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	uintBookId := uint(bookId)

	book := getBook(db, uintBookId)
	return c.JSON(book)
}

func createBookHandler(c *fiber.Ctx) error {
	bookCreate := new(Book)

	if err := c.BodyParser(bookCreate); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := creatBook(db, bookCreate)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Create book successful!",
	})

}

func updateBookHandler(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	updatedBook := new(Book)

	if err := c.BodyParser(updatedBook); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	updatedBook.ID = uint(bookId)

	err = updateBook(db, updatedBook)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// return c.JSON(book)
	return c.JSON(fiber.Map{
		"message": "Update book successful!",
	})

}

func deleteBookHandler(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	uintBookId := uint(bookId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = deleteBook(db, uintBookId)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Delete book successful!",
	})

}

func searchBookHandler(c *fiber.Ctx) error {
	bookName := c.Query("name")
	fmt.Println("Searching for book with name:", bookName)

	book, err := searchBook(db, bookName)

	if err != nil {
		fmt.Println("Error:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(book)

}

func searchBooksHandler(c *fiber.Ctx) error {
	bookName := c.Query("name")
	order := c.Query("order")

	books, err := searchBooks(db, bookName, order)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(books)

}
