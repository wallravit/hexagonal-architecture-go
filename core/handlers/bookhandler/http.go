package bookhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/wallravit/hexagonal-architecture/core/models"
	"github.com/wallravit/hexagonal-architecture/core/ports"
)

type HTTPHandler struct {
	bookService ports.BookService
}

func NewHTTPHandler(bookService ports.BookService) *HTTPHandler {
	return &HTTPHandler{
		bookService: bookService,
	}
}

func (handler *HTTPHandler) Get(ctx *gin.Context) {
	bookID := ctx.Param("id")
	book, err := handler.bookService.Get(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, book)
}

func (handler *HTTPHandler) Add(ctx *gin.Context) {
	newBook := models.NewBook{}
	err := ctx.BindJSON(&newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
		return
	}
	book, err := handler.bookService.Add(newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, book)
}

func (handler *HTTPHandler) Occupie(ctx *gin.Context) {
	bookID := ctx.Param("id")
	book, err := handler.bookService.Occupie(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, book)
}

func (handler *HTTPHandler) Return(ctx *gin.Context) {
	bookID := ctx.Param("id")
	book, err := handler.bookService.Return(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, book)
}
