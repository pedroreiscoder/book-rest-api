package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroreiscoder/book-rest-api/data"
)

func GetBooks(c *gin.Context) {
	books := data.GetBooks()
	c.IndentedJSON(http.StatusOK, books)
}
