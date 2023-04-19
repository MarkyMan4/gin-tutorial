package main

import (
	"github.com/MarkyMan4/gin-tutorial/controller"
	"github.com/MarkyMan4/gin-tutorial/service"
	"github.com/gin-gonic/gin"
)

var (
	bookService    service.BookService       = service.New()
	bookController controller.BookController = controller.New(bookService)
)

func main() {
	server := gin.Default()

	server.GET("/books", func(ctx *gin.Context) {
		ctx.JSON(200, bookController.FindAll())
	})

	server.POST("/books", func(ctx *gin.Context) {
		ctx.JSON(200, bookController.Save(ctx))
	})

	server.Run(":8080")
}
