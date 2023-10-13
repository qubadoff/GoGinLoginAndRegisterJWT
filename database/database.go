package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var GlobalDB *gorm.DB

func InitDatabase() (err error) {
	config, err := godotenv.Read()

	if err != nil {
		log.Fatal("Error Reading the .env file !")
	}

	// DB Connect variables
	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DATABASE_HOST"],
		config["DB_DATABASE"],
	)

	//db connect
	GlobalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}
	return
}
