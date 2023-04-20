package controller

import (
	"github.com/MarkyMan4/gin-tutorial/model"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	books []model.Book
}

func New() *BookController {
	return &BookController{}
}

func (c *BookController) FindAll() []model.Book {
	return c.books
}

func (c *BookController) Save(ctx *gin.Context) model.Book {
	var book model.Book
	ctx.BindJSON(&book)
	c.books = append(c.books, book)

	return book
}
