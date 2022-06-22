package data

import "github.com/pedroreiscoder/book-rest-api/models"

func GetBooks() []models.Book {
	var books []models.Book
	db.Find(&books)
	return books
}

func CreateBook(book *models.Book) {
	db.Create(book)
}
