package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroreiscoder/book-rest-api/controllers"
)

func New() *gin.Engine {
	router := gin.Default()

	router.GET("/api/books", controllers.GetBooks)

	return router
}