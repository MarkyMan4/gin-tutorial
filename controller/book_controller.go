package controller

import (
	"github.com/MarkyMan4/gin-tutorial/entity"
	"github.com/MarkyMan4/gin-tutorial/service"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	FindAll() []entity.Book
	Save(ctx *gin.Context) entity.Book
}

type bookController struct {
	service service.BookService
}

func New(service service.BookService) BookController {
	return &bookController{service: service}
}

func (c *bookController) FindAll() []entity.Book {
	return c.service.FindAll()
}

func (c *bookController) Save(ctx *gin.Context) entity.Book {
	var book entity.Book
	ctx.BindJSON(&book)
	c.service.Save(book)

	return book
}
