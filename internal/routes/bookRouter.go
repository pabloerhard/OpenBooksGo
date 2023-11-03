package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pabloerhard/openBooksGo/internal/controllers"
)

func SetUpBooksRouter(router *gin.Engine) {
	booksGroup := router.Group("/book")

	booksGroup.POST("/create", controllers.InsertOneBookController)
}
