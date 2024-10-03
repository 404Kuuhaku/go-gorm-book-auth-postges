package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ================ //
// PUBLISHER API
// ================ //
func getAllPublishersHandler(c *fiber.Ctx) error {
	publishers := getPublishers(db)
	return c.JSON(publishers)
}

func createPublisherHandler(c *fiber.Ctx) error {
	newPublisher := new(Publisher)

	if err := c.BodyParser(newPublisher); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := createPublisher(db, newPublisher)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Create Publisher Successful!",
	})

}

func updatePublisherHandler(c *fiber.Ctx) error {
	publisherId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newPublisher := new(Publisher)

	if err := c.BodyParser(newPublisher); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newPublisher.ID = uint(publisherId)

	err = updatePublisher(db, newPublisher)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Update Publisher Successful!",
	})

}

func deletePublisherHandler(c *fiber.Ctx) error {
	publisherId, err := strconv.Atoi(c.Params("id"))
	uintPublisherId := uint(publisherId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = deletePublisher(db, uintPublisherId)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Delete Publisher Successful!",
	})

}

// ================ //
// AUTHOR API
// ================ //
func getAllAuthorsHandler(c *fiber.Ctx) error {
	authors := getAuthors(db)
	return c.JSON(authors)
}

func createAuthorHandler(c *fiber.Ctx) error {
	newAuthor := new(Author)

	if err := c.BodyParser(newAuthor); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := createAuthor(db, newAuthor)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Create Author Successful!",
	})

}

func updateAuthorHandler(c *fiber.Ctx) error {
	authorId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newAuthor := new(Author)

	if err := c.BodyParser(newAuthor); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newAuthor.ID = uint(authorId)

	err = updateAuthor(db, newAuthor)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Update Author Successful!",
	})

}

func deleteAuthorHandler(c *fiber.Ctx) error {
	authorId, err := strconv.Atoi(c.Params("id"))
	uintAuthorId := uint(authorId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = deleteAuthor(db, uintAuthorId)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Delete Author Successful!",
	})

}

// func getAllBooksHandler(c *fiber.Ctx) error {
// 	books := getBooks(db)
// 	return c.JSON(books)
// }

// func getBookHandler(c *fiber.Ctx) error {
// 	bookId, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}
// 	uintBookId := uint(bookId)

// 	book := getBook(db, uintBookId)
// 	return c.JSON(book)
// }

// func createBookHandler(c *fiber.Ctx) error {
// 	bookCreate := new(Book)

// 	if err := c.BodyParser(bookCreate); err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	err := creatBook(db, bookCreate)

// 	if err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Create book successful!",
// 	})

// }

// func updateBookHandler(c *fiber.Ctx) error {
// 	bookId, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	updatedBook := new(Book)

// 	if err := c.BodyParser(updatedBook); err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	updatedBook.ID = uint(bookId)

// 	err = updateBook(db, updatedBook)

// 	if err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Update book successful!",
// 	})

// }

// func deleteBookHandler(c *fiber.Ctx) error {
// 	bookId, err := strconv.Atoi(c.Params("id"))
// 	uintBookId := uint(bookId)
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	err = deleteBook(db, uintBookId)

// 	if err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Delete book successful!",
// 	})

// }

// func searchBookHandler(c *fiber.Ctx) error {
// 	bookName := c.Query("name")

// 	book, err := searchBook(db, bookName)

// 	if err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	return c.JSON(book)

// }

// func searchBooksHandler(c *fiber.Ctx) error {
// 	bookName := c.Query("name")
// 	order := c.Query("order")

// 	books, err := searchBooks(db, bookName, order)

// 	if err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	return c.JSON(books)

// }
