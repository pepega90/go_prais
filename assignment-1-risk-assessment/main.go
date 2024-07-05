package main

import (
	"assignment_1/handler"
	"assignment_1/repository/postgres_gorm"
	"assignment_1/router"
	"assignment_1/service"

	"github.com/gin-gonic/gin"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup gorm connection
	dsn := "postgresql://prais:prais@localhost:5432/db_prais"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	// setup repository
	userRepo := postgres_gorm.NewUserRepository(gormDB)
	submissionRepo := postgres_gorm.NewSubmissionRepository(gormDB)

	// service and handler declaration
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	submissionService := service.NewSubmissionService(submissionRepo)
	submissionHandler := handler.NewSubmissionHandler(submissionService)

	// Routes
	router.SetupRouter(r, userHandler, submissionHandler)

	// Run the server
	log.Println("Running server on port 4000")
	r.Run(":4000")
}
