package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wallravit/hexagonal-architecture/core/handlers/bookhandler"
	"github.com/wallravit/hexagonal-architecture/core/services/booksservice"
	"github.com/wallravit/hexagonal-architecture/repositories/bookrepository"
)

func main() {
	bookRepository := bookrepository.NewPgSQL()
	bookService := booksservice.New(bookRepository)
	bookHandler := bookhandler.NewHTTPHandler(bookService)

	route := gin.New()

	route.GET("/books/:id", bookHandler.Get)
	route.POST("/books/add", bookHandler.Add)
	route.POST("books/:id/occupie", bookHandler.Occupie)
	route.POST("books/:id/return", bookHandler.Return)

	route.Run(":8000")
}
