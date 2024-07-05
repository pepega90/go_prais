package grpc

import (
	"context"
	"fmt"
	"log"
	"user_grpc_crud/entity"
	pb "user_grpc_crud/protos"
	"user_grpc_crud/services"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(ctx context.Context, _ *emptypb.Empty) (*pb.GetUsersResponse, error) {
	users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var userProtoList []*pb.User
	for _, v := range users {
		userProtoList = append(userProtoList, &pb.User{
			Id:        int32(v.ID),
			Name:      v.Name,
			Emai:      v.Email,
			Password:  v.Password,
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
		})
	}

	return &pb.GetUsersResponse{
		Users: userProtoList,
	}, nil
}

func (h *UserHandler) GetUserByID(c context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	u, err := h.userService.GetUserByID(c, int(req.GetId()))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:        int32(u.ID),
			Name:      u.Name,
			Emai:      u.Email,
			Password:  u.Password,
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt),
		},
	}

	return res, nil
}

func (h *UserHandler) CreateUser(c context.Context, req *pb.CreateUserRequest) (*pb.MutationResponse, error) {
	u, err := h.userService.CreateUser(c, &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success created user with ID %d", u.ID),
	}, nil
}

func (h *UserHandler) UpdateUser(c context.Context, req *pb.UpdateUserRequest) (*pb.MutationResponse, error) {
	updatedUser, err := h.userService.UpdateUser(c, int(req.GetId()), entity.User{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("success update user with id %v", updatedUser.ID),
	}, nil
}

func (h *UserHandler) DeleteUser(c context.Context, req *pb.DeleteUserRequest) (*pb.MutationResponse, error) {
	err := h.userService.DeleteUser(c, int(req.GetId()))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.MutationResponse{
		Message: fmt.Sprintf("success delete user with id %v", req.GetId()),
	}, nil
}
