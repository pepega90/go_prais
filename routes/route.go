package routes

import (
	"go_prais/handler"
	"go_prais/repository/slice"
	"go_prais/services"

	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine) {
	userRepo := slice.NewSliceRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.GET("/", userHandler.GetAllUsers)
	r.POST("/", userHandler.CreateUser)
	r.PUT("/:id", userHandler.UpdateUser)
	r.GET("/:id", userHandler.GetUser)
	r.DELETE("/:id", userHandler.DeleteUser)
}
