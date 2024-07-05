// package router mengatur rute untuk aplikasi
package router

import (
	"assignment_1/handler"
	"assignment_1/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter menginisialisasi dan mengatur rute untuk aplikasi
func SetupRouter(r *gin.Engine,
	userHandler handler.IUserHandler,
	submissionsHandler handler.ISubmissionHandler,
) {
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/:id", userHandler.GetUser)
	usersPublicEndpoint.GET("", userHandler.GetAllUsers)
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)
	usersPublicEndpoint.POST("", middleware.AuthMiddleware, userHandler.CreateUser)
	usersPublicEndpoint.POST("/", middleware.AuthMiddleware, userHandler.CreateUser)
	usersPublicEndpoint.PUT("/:id", middleware.AuthMiddleware, userHandler.UpdateUser)
	usersPublicEndpoint.DELETE("/:id", middleware.AuthMiddleware, userHandler.DeleteUser)

	submissionsPublicEndpoint := r.Group("/submissions")
	submissionsPublicEndpoint.GET("/:id", submissionsHandler.GetSubmission)
	submissionsPublicEndpoint.GET("", submissionsHandler.GetAllSubmissions)
	submissionsPublicEndpoint.GET("/", submissionsHandler.GetAllSubmissions)
	submissionsPublicEndpoint.POST("", middleware.AuthMiddleware, submissionsHandler.CreateSubmission)
	submissionsPublicEndpoint.POST("/", middleware.AuthMiddleware, submissionsHandler.CreateSubmission)
	submissionsPublicEndpoint.DELETE("/:id", middleware.AuthMiddleware, submissionsHandler.DeleteSubmission)
}
