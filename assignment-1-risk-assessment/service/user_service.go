package service

import (
	"assignment_1/entity"
	"context"
	"log"
)

// IUserService mendefinisikan interface untuk layanan pengguna
type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context, limit, offset int) ([]entity.User, error)
	GetUserCount(ctx context.Context) (int, error)
}

// IUserRepository mendefinisikan interface untuk repository pengguna
type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetUserCount(ctx context.Context) (int, error)
	GetAllUsers(ctx context.Context, limit, offset int) ([]entity.User, error)
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type userService struct {
	userRepo IUserRepository
}

// NewUserService membuat instance baru dari userService
func NewUserService(userRepo IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	createdUser, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		log.Printf("error create user: %v", err.Error())
		return entity.User{}, err
	}
	return createdUser, nil
}

func (u *userService) GetUserCount(ctx context.Context) (int, error) {
	count, err := u.userRepo.GetUserCount(ctx)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return count, nil
}

func (u *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	user, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("error get user with id = %v", id)
		return entity.User{}, err
	}
	return user, nil
}

func (u *userService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	updatedUser, err := u.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		log.Printf("error updating user: %v", err.Error())
		return entity.User{}, err
	}
	return updatedUser, nil
}

func (u *userService) DeleteUser(ctx context.Context, id int) error {
	err := u.userRepo.DeleteUser(ctx, id)
	if err != nil {
		log.Printf("error deleting user: %v", err.Error())
		return err
	}
	return nil
}

func (u *userService) GetAllUsers(ctx context.Context, limit, offset int) ([]entity.User, error) {
	listUser, err := u.userRepo.GetAllUsers(ctx, limit, offset)
	if err != nil {
		log.Printf("error get all users: %v", err.Error())
		return nil, err
	}
	return listUser, nil
}
