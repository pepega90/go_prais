package routes

import (
	"go_prais/handler"
	postgresgorm "go_prais/repository/postgres_gorm"
	"go_prais/services"
	"go_prais/utils"

	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine) {
	// userRepo := slice.NewSliceRepository()
	// userService := services.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)
	// db := utils.DBPostgre() => PostgreSQL with pgx
	db := utils.DBPostgreGorm() // => PostgreSQL with GORM ORM
	userRepo := postgresgorm.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.GET("/", userHandler.GetAllUsers)
	r.POST("/", userHandler.CreateUser)
	r.PUT("/:id", userHandler.UpdateUser)
	r.GET("/:id", userHandler.GetUser)
	r.DELETE("/:id", userHandler.DeleteUser)
}
