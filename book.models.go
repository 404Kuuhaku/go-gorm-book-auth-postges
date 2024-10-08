package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Publisher   string `json:"publisher"`
}

func getBooks(db *gorm.DB) []Book {
	var book []Book
	result := db.Find(&book)

	if result.Error != nil {
		log.Fatalf("Error  creating book : %v", result.Error)
	}

	fmt.Println("Get book successful!")
	return book
}

func getBook(db *gorm.DB, id uint) *Book {
	var book Book
	result := db.First(&book, id)

	if result.Error != nil {
		log.Fatalf("Error  creating book : %v", result.Error)
	}

	fmt.Println("Get book successful!")
	return &book
}

func searchBook(db *gorm.DB, bookName string) (*Book, error) {
	var book Book
	result := db.Where("name = ?", bookName).First(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("Search book successful!")

	return &book, nil

}

func searchBooks(db *gorm.DB, bookName string, order string) ([]Book, error) {
	var book []Book
	result := db.Where("name = ?", bookName).Order(order).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("Search books successful!")

	return book, nil

}

func creatBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func updateBook(db *gorm.DB, book *Book) error {
	// result := db.Save(book) //this medthod will update everything in field, and we don't want to update Create At field, becuase it's never change
	result := db.Model(&book).Updates(book) //use this instead Save becuse it' did not change Creat At Field

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func deleteBook(db *gorm.DB, id uint) error {
	var book Book
	result := db.Delete(&book, id) //SOFT DELETE
	// result := db.Unscoped().Delete(&book, id) //HARD DELETE

	if result.Error != nil {
		return result.Error
	}

	return nil

}
