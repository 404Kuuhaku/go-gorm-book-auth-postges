package main

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	// "log"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

func createUser(db *gorm.DB, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	result := db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func loginUser(db *gorm.DB, user *User) (string, error) {
	// get user from email
	selectedUser := new(User)
	result := db.Where("email = ?", user.Email).First(selectedUser)

	if result.Error != nil {
		return "", result.Error
	}

	// compare password
	err := bcrypt.CompareHashAndPassword(
		[]byte(selectedUser.Password),
		[]byte(user.Password),
	)

	if err != nil {
		return "", err
	}

	// if it's correct return jwt

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
