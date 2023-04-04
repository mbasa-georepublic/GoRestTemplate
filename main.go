package main

import (
	"net/http"

	"example/webservice/controllers"
	"example/webservice/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World",
			"addr": c.Request.Host})
	})

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.PATCH("/books/:id", controllers.UpdateBook)

	err := r.Run()

	if err == nil {
		return
	}
}
