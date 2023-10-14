package controllers

import (
	"LoginAndRegisterApiJWT/auth"
	"LoginAndRegisterApiJWT/database"
	"LoginAndRegisterApiJWT/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"RefreshToken"`
}

func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}
	result := database.GlobalDB.Where("email = ?", payload.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	signedtoken, err := jwtWrapper.RefreshToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	tokenResponse := LoginResponse{
		Token:        signedToken,
		RefreshToken: signedtoken,
	}
	c.JSON(200, tokenResponse)
}

func Signup(contex *gin.Context) {
	var user models.User
	err := contex.ShouldBindJSON(&user)

	if err != nil {
		log.Println(err)
		contex.JSON(400, gin.H{
			"Error": "Invalid Inputs !",
		})
		contex.Abort()
		return
	}

	err = user.HashPassword(user.Password)

	if err != nil {
		log.Println(err.Error())
		contex.JSON(500, gin.H{
			"Error": "Error Hashing Password !",
		})
		contex.Abort()
		return
	}

	err = user.CreateUserRecord()

	if err != nil {
		log.Println(err)
		contex.JSON(500, gin.H{
			"Error": "Error Creating User",
		})
		contex.Abort()
		return
	}

	contex.JSON(200, gin.H{
		"Message": "Sucessfully Register",
	})
}

func CreateBook(contex *gin.Context) {
	var book models.Book
	err := contex.ShouldBindJSON(&book)

	if err != nil {
		log.Println(err)

		contex.JSON(400, gin.H{
			"Error": "Invalid Input !",
		})
		contex.Abort()
		return
	}

	err = book.CreateBookRecord()

	if err != nil {
		log.Println(err)
		contex.JSON(500, gin.H{
			"Error": "Internal Server Error !",
		})
	}

	contex.JSON(200, gin.H{
		"Message": "Successfully added !",
	})
}

func CreateBookCat(contex *gin.Context) {
	var category models.BookCategory
	err := contex.ShouldBindJSON(&category)

	if err != nil {
		log.Println(err)
		contex.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		contex.Abort()
		return
	}

	err = category.CreateBookCategoryRecord()

	if err != nil {
		log.Println(err)
		contex.JSON(500, gin.H{
			"Error": "Internal Server Error !",
		})
		contex.Abort()
		return
	}

	contex.JSON(200, gin.H{
		"Message": "Book Successfully Added !",
	})

}
