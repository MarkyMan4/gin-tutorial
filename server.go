package main

import (
	"net/http"

	"github.com/MarkyMan4/gin-tutorial/controller"
	"github.com/gin-gonic/gin"
)

var (
	bookController *controller.BookController = controller.New()
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
