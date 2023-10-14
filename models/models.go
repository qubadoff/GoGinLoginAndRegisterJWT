package models

import (
	"LoginAndRegisterApiJWT/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

type Book struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Author      string `json:"author"`
	CatID       int    `json:"catID" binding:"required"`
}

type BookCategory struct {
	ID   int    `gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	gorm.Model
}

func (user *User) CreateUserRecord() error {
	result := database.GlobalDB.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (book *Book) CreateBookRecord() error {
	result := database.GlobalDB.Create(&book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (bookCat *BookCategory) CreateBookCategoryRecord() error {
	result := database.GlobalDB.Create(&bookCat)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

func (user *User) CheckPassword(providePassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providePassword))

	if err != nil {
		return err
	}

	return nil
}
