package data

import "github.com/pedroreiscoder/book-rest-api/models"

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	result := db.Find(&books)
	return books, result.Error
}

func GetBook(id uint64) (models.Book, error) {
	var book models.Book
	result := db.First(&book, id)
	return book, result.Error
}

func CreateBook(book *models.Book) {
	db.Create(book)
}
