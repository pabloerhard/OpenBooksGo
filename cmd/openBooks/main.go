package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pabloerhard/openBooksGo/internal/middleware"
	"github.com/pabloerhard/openBooksGo/internal/routes"
	"os"
)

func main() {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	routes.SetUpBooksRouter(router)
	routes.SetUpUserRouter(router)
	err := router.Run(os.Getenv("ADDRESS"))
	if err != nil {
		return
	}
}
