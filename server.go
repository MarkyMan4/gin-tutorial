package main

import (
	"net/http"

	"github.com/MarkyMan4/gin-tutorial/controller"
	"github.com/MarkyMan4/gin-tutorial/db"
	"github.com/gin-gonic/gin"
)

var (
	bookDb         *db.BookDb                 = db.New("books.db")
	bookController *controller.BookController = controller.New(bookDb)
)

func main() {
	server := gin.Default()

	server.GET("/books", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, bookController.FindAll())
	})

	server.POST("/books", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, bookController.Save(ctx))
	})

	server.Run(":8080")
}
