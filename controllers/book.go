package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pedroreiscoder/book-rest-api/data"
	"github.com/pedroreiscoder/book-rest-api/models"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {
	books := data.GetBooks()
	c.IndentedJSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	book, err := data.GetBook(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	err := c.BindJSON(&book)

	if err != nil {
		return
	}

	data.CreateBook(&book)

	c.IndentedJSON(http.StatusCreated, book)
}
