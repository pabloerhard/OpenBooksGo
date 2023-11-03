package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pabloerhard/openBooksGo/internal/models"
	"github.com/pabloerhard/openBooksGo/internal/services"
	"log"
	"net/http"
)

// InsertOneBookController Insert Book Controller
func InsertOneBookController(c *gin.Context) {
	var book models.Book

	if err := c.BindJSON(&book); err != nil {
		log.Fatal(err)
	}

	err := services.InsertOneBook(book)

	if err != nil {
		c.IndentedJSON(http.StatusConflict, err)
	} else {
		c.IndentedJSON(http.StatusCreated, book)
	}
}
