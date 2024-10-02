package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var jwtSecretKey string = os.Getenv("JWT_SECRETKEY")

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

func authRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	// token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte(jwtSecretKey), nil
	// })

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claim := token.Claims.(*jwt.MapClaims)

	fmt.Println(claim)

	return c.Next()

}

func main() {
	db = SetupDatabase()
	db.AutoMigrate(&Book{}, &User{})

	app := fiber.New()
	// BOOK API
	apiBook := app.Group("/book-api", authRequired)
	apiBook.Get("/books", getAllBooksHandler)
	apiBook.Get("/book/:id", getBookHandler)
	apiBook.Get("/search", searchBookHandler)
	apiBook.Get("/searchs", searchBooksHandler)
	apiBook.Post("/book", createBookHandler)
	apiBook.Put("/book/:id", updateBookHandler)
	apiBook.Delete("/book/:id", deleteBookHandler)

	// app.Get("/books", getAllBooksHandler)
	// app.Get("/book/:id", getBookHandler)
	// app.Get("/search", searchBookHandler)
	// app.Get("/searchs", searchBooksHandler)
	// app.Post("/book", createBookHandler)
	// app.Put("/book/:id", updateBookHandler)
	// app.Delete("/book/:id", deleteBookHandler)

	// USER API
	app.Post("/register", createUserHandler)
	app.Post("/login", loginUserHandler)
	app.Listen(":8080")
}
