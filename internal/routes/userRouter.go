package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pabloerhard/openBooksGo/internal/controllers"
	"github.com/pabloerhard/openBooksGo/internal/middleware"
)

func SetUpUserRouter(router *gin.Engine) {
	userGroup := router.Group("/user")
	// RequireAuth is our middleware, it validates the credentials of the user
	userGroup.POST("/create", controllers.InsertUser)
	userGroup.POST("/login", controllers.LoginUser)
	userGroup.GET("/validate", middleware.RequireAuth, controllers.Validate)
	userGroup.GET("/get/:userId", controllers.FindUser)
	userGroup.PUT("/update/want/:googleId/:userId", controllers.InsertWantToReadBook)
	userGroup.PUT("/update/read/:googleId/:userId", controllers.InsertReadBook)
	userGroup.PUT("/update/reading/:googleId/:userId", controllers.InsertReadingBook)
}
