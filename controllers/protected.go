package controllers

import (
	"LoginAndRegisterApiJWT/database"
	"LoginAndRegisterApiJWT/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Profile(context *gin.Context) {
	var user models.User

	email, _ := context.Get("email")

	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		context.JSON(404, gin.H{
			"Error": "User Not Found !",
		})
		context.Abort()
		return
	}

	if result.Error != nil {
		context.JSON(500, gin.H{
			"Error": "Could Not Get User Profile",
		})
		context.Abort()
		return
	}

	user.Password = ""

	context.JSON(200, user)
}
