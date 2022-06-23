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

func CreateBook(book *models.Book) error {
	return db.Create(book).Error
}

func UpdateBook(id uint64, newBook models.Book) error {
	var book models.Book
	result := db.First(&book, id)

	if result.Error != nil {
		return result.Error
	}

	book.Title = newBook.Title
	book.Author = newBook.Author
	book.Quantity = newBook.Quantity

	result = db.Save(&book)
	return result.Error
}

func DeleteBook(book *models.Book) error {
	return db.Delete(book).Error
}
