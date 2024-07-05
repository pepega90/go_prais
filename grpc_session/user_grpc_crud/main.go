package main

import (
	"log"
	"net"

	grpcHandler "user_grpc_crud/handler/grpc"
	pb "user_grpc_crud/protos"
	"user_grpc_crud/repository/postgres_gorm"
	"user_grpc_crud/services"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgresql://prais:prais@localhost:5432/db_prais"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err.Error())
	}

	userRepo := postgres_gorm.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := grpcHandler.NewUserHandler(userService)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Running grpc server in port :50051")
	_ = grpcServer.Serve(lis)
}
