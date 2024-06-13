package services

import "go_prais/model"

type IUserService interface {
	GetAllUsers() []*model.User
	CreateUser(user model.User) *model.User
	UpdateUser(id int, user model.User) (*model.User, error)
	DeleteUser(id int) error
	GetUser(id int) (*model.User, error)
}

type IUserRepository interface {
	GetAllUsers() []*model.User
	CreateUser(user model.User) *model.User
	UpdateUser(id int, user model.User) (*model.User, error)
	DeleteUser(id int) error
	GetUser(id int) (*model.User, error)
}

type userService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) *userService {
	return &userService{userRepo}
}

func (s *userService) GetAllUsers() []*model.User {
	users := s.userRepo.GetAllUsers()
	return users
}

func (s *userService) CreateUser(req model.User) *model.User {
	return s.userRepo.CreateUser(req)
}

func (s *userService) GetUser(id int) (*model.User, error) {
	user, err := s.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(id int, req model.User) (*model.User, error) {
	updatedUser, err := s.userRepo.UpdateUser(id, req)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (s *userService) DeleteUser(id int) error {
	err := s.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
