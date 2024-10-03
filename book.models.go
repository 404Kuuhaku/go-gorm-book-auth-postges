package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	PublisherID uint
	Publisher   Publisher
	Authors     []Author `gorm:"many2many:author_books;"`
}

type Publisher struct {
	gorm.Model
	Details string
	Name    string
}

type Author struct {
	gorm.Model
	Name  string
	Books []Book `gorm:"many2many:author_books;"`
}

type AuthorBook struct {
	AuthorID uint
	Author   Author
	BookID   uint
	Book     Book
}

// ================ //
// PUBLISHER FUNCTION
// ================ //
func getPublishers(db *gorm.DB) []Publisher {
	var publishers []Publisher
	result := db.Find(&publishers)
	if result.Error != nil {
		log.Fatalf("Error Get Publisher : %v", result.Error)
	}
	fmt.Println("Get book successful!")
	return publishers
}

func createPublisher(db *gorm.DB, publisher *Publisher) error {
	result := db.Create(publisher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func updatePublisher(db *gorm.DB, publisher *Publisher) error {
	result := db.Model(&publisher).Updates(publisher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func deletePublisher(db *gorm.DB, id uint) error {
	var publisher Publisher
	result := db.Delete(&publisher, id) //SOFT DELETE
	// result := db.Unscoped().Delete(&publisher, id) //HARD DELETE
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ================ //
// AUTHOR FUNCTION
// ================ //
func getAuthors(db *gorm.DB) []Author {
	var authors []Author

	result := db.Find(&authors)

	if result.Error != nil {
		log.Fatalf("Error Get Author : %v", result.Error)
	}

	fmt.Println("Get book successful!")
	return authors
}

func createAuthor(db *gorm.DB, author *Author) error {
	result := db.Create(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func updateAuthor(db *gorm.DB, author *Author) error {
	result := db.Model(&author).Updates(author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func deleteAuthor(db *gorm.DB, id uint) error {
	var author Author
	result := db.Delete(&author, id) //SOFT DELETE
	// result := db.Unscoped().Delete(&author, id) //HARD DELETE
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ================ //
// BOOK WITH AUTHOR FUNCTION
// ================ //
func getBookWithAuthors(db *gorm.DB, bookID uint) (*Book, error) {
	var book Book
	result := db.Preload("Authors").First(&book, bookID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func getListBooksOfAuthor(db *gorm.DB, authorID uint) ([]Book, error) {
	var books []Book
	result := db.Joins("JOIN author_books on author_books.book_id = books.id").
		Where("author_books.author_id = ?", authorID).
		Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func createBookWithAuthor(db *gorm.DB, book *Book, authorIDs []uint) error {
	if err := db.Create(book).Error; err != nil {
		return err
	}

	return nil
}

func updateBookWithAuthor(db *gorm.DB, bookID uint, newAuthors []Author) error {
	// Step 1: Find the book you want to update
	var book Book
	if err := db.Preload("Authors").First(&book, bookID).Error; err != nil {
		return err
	}

	// Step 2: Update the book's authors with the new list
	// This will update the `author_books` join table accordingly
	err := db.Model(&book).Association("Authors").Replace(newAuthors)
	if err != nil {
		return err
	}
	return nil
}

func deleteBookWithAuthor(db *gorm.DB, bookID uint, authorsToDelete []Author) error {
	// Step 1: Find the book you want to update
	var book Book
	if err := db.Preload("Authors").First(&book, bookID).Error; err != nil {
		return err
	}

	// Step 2: Delete the specified authors from the book's associations
	err := db.Model(&book).Association("Authors").Delete(authorsToDelete)
	if err != nil {
		return err
	}

	return nil
}

// ================ //
// BOOK WITH PUBLISHER FUNCTION
// ================ //

func getBookWithPublisher(db *gorm.DB, bookID uint) (*Book, error) {
	var book Book
	result := db.Preload("Publishers").First(&book, bookID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func updateBookWithPublisher(db *gorm.DB, bookID uint, publishersToDelete []Publisher) error {
	var book Book
	if err := db.Preload("Publishers").First(&book, bookID).Error; err != nil {
		return err
	}

	err := db.Model(&book).Association("Publishers").Delete(publishersToDelete)
	if err != nil {
		return err
	}
	return nil
}

func deleteBookWithPublisher(db *gorm.DB, bookID uint) (*Book, error) {
	var book Book
	result := db.Preload("Publisher").First(&book, bookID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

// func getBooks(db *gorm.DB) []Book {
// 	var book []Book
// 	result := db.Find(&book)

// 	if result.Error != nil {
// 		log.Fatalf("Error  creating book : %v", result.Error)
// 	}

// 	fmt.Println("Get book successful!")
// 	return book
// }

// func getBook(db *gorm.DB, id uint) *Book {
// 	var book Book
// 	result := db.First(&book, id)

// 	if result.Error != nil {
// 		log.Fatalf("Error  creating book : %v", result.Error)
// 	}

// 	fmt.Println("Get book successful!")
// 	return &book
// }

// func searchBook(db *gorm.DB, bookName string) (*Book, error) {
// 	var book Book
// 	result := db.Where("name = ?", bookName).First(&book)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	fmt.Println("Search book successful!")

// 	return &book, nil

// }

// func searchBooks(db *gorm.DB, bookName string, order string) ([]Book, error) {
// 	var book []Book
// 	result := db.Where("name = ?", bookName).Order(order).Find(&book)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	fmt.Println("Search books successful!")

// 	return book, nil

// }

// func creatBook(db *gorm.DB, book *Book) error {
// 	result := db.Create(book)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil

// }

// func updateBook(db *gorm.DB, book *Book) error {
// 	// result := db.Save(book) //this medthod will update everything in field, and we don't want to update Create At field, becuase it's never change
// 	result := db.Model(&book).Updates(book) //use this instead Save becuse it' did not change Creat At Field

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil

// }

// func deleteBook(db *gorm.DB, id uint) error {
// 	var book Book
// 	result := db.Delete(&book, id) //SOFT DELETE
// 	// result := db.Unscoped().Delete(&book, id) //HARD DELETE

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil

// }
