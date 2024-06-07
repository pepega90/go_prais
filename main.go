package main

import (
	"go_prais/middleware"
	"go_prais/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware)

	routes.Routing(r)

	log.Println("Server start")
	r.Run(":4000")
}
