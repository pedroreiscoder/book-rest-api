package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroreiscoder/book-rest-api/data"
	"github.com/pedroreiscoder/book-rest-api/models"
)

func GetBooks(c *gin.Context) {
	books := data.GetBooks()
	c.IndentedJSON(http.StatusOK, books)
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
