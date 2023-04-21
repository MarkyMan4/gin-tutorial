package controller

import (
	"github.com/MarkyMan4/gin-tutorial/db"
	"github.com/MarkyMan4/gin-tutorial/model"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookDb *db.BookDb
	books  []model.Book
}

func New(bookDb *db.BookDb) *BookController {
	return &BookController{bookDb: bookDb}
}

func (c *BookController) FindAll() []model.Book {
	var books []model.Book
	result := c.bookDb.Db.Find(&books)

	if result.Error != nil {
		panic("failed to retrieve books")
	}

	return books
}

func (c *BookController) Save(ctx *gin.Context) model.Book {
	var book model.Book
	ctx.BindJSON(&book)
	c.bookDb.Db.Create(&book)

	return book
}
