package main

import (
	"github.com/gin-gonic/gin"
	"golangAPI/controllers"
	"golangAPI/models"
)

func main() {
	models.ConnectDataBase()

	router := gin.Default()
	public := router.Group("/api")

	public.POST("/register", controllers.Register)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
