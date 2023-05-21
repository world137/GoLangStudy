package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json: "price"`
}

var books = []Book{
	{ID: "001", Name: "hello", Price: "100"},
	{ID: "002", Name: "world", Price: "130"},
	{ID: "003", Name: "golang", Price: "150"},
}

func main() {
	r := gin.New()

	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	r.Run()
}
