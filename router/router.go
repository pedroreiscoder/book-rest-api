package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroreiscoder/book-rest-api/controllers"
)

func New() *gin.Engine {
	router := gin.Default()

	router.GET("/api/books", controllers.GetBooks)
	router.GET("/api/book/:id", controllers.GetBook)
	router.POST("/api/book", controllers.CreateBook)
	router.PUT("/api/book/:id", controllers.UpdateBook)

	return router
}
