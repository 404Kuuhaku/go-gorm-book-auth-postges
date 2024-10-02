package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func SetupDatabase() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error!")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Database Connected!")

	return db
}

func main() {
	db = SetupDatabase()
	db.AutoMigrate(&Book{})

	app := fiber.New()
	app.Get("/books", getAllBooksHandler)
	app.Get("/book/:id", getBookHandler)
	app.Get("/search", searchBookHandler)
	app.Get("/searchs", searchBooksHandler)
	app.Post("/book", createBookHandler)
	app.Put("/book/:id", updateBookHandler)
	app.Delete("/book/:id", deleteBookHandler)
	app.Listen(":8080")
}
