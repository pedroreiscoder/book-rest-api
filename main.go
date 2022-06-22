package main

import (
	"github.com/pedroreiscoder/book-rest-api/data"
	"github.com/pedroreiscoder/book-rest-api/router"
)

func main() {
	data.Init()
	r := router.New()
	r.Run("localhost:8080")
}
