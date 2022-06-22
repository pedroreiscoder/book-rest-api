package data

import (
	"log"

	"github.com/pedroreiscoder/book-rest-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:P3dr0@123@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Book{})
}
