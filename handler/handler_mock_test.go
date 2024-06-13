package handler

import (
	"errors"
	"go_prais/model"
)

type MockUserService struct{}

var (
	mockUsers  []*model.User
	mockNextId int
)

func init() {
	mockUsers = []*model.User{
		{Id: 1, Name: "User 1", Email: "user1@example.com", Password: "pass1"},
		{Id: 2, Name: "User 2", Email: "user2@example.com", Password: "pass2"},
	}
	mockNextId = 3
}

func (s *MockUserService) CreateUser(req model.User) *model.User {
	mockUsers = append(mockUsers, &req)
	return &req
}

func (m *MockUserService) GetUser(id int) (*model.User, error) {
	for _, user := range mockUsers {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *MockUserService) UpdateUser(id int, user model.User) (*model.User, error) {
	for i, u := range mockUsers {
		if u.Id == id {
			user.Id = id
			mockUsers[i] = &user
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *MockUserService) DeleteUser(id int) error {
	for i, user := range mockUsers {
		if user.Id == id {
			mockUsers = append(mockUsers[:i], mockUsers[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (m *MockUserService) GetAllUsers() []*model.User {
	return mockUsers
}
