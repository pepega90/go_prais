package routes

import (
	"go_prais/handler"

	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine) {
	r.GET("/", handler.GetAllUser)
	r.POST("/", handler.CreateUser)
	r.PUT("/", handler.UpdateUser)
	r.GET("/:id", handler.GetUser)
	r.DELETE("/:id", handler.DeleteUser)
}
