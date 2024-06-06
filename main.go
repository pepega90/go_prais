package main

import (
	"go_prais/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handler.GetAllUser)
	r.POST("/", handler.CreateUser)
	r.PUT("/", handler.UpdateUser)
	r.GET("/:id", handler.GetUser)
	r.DELETE("/:id", handler.DeleteUser)

	log.Println("Server start")
	r.Run(":4000")
}
