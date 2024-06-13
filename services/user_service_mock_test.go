package services

import (
	"errors"
	"go_prais/model"
	"time"
)

type MockRepository struct {
	users []*model.User
}

func (m *MockRepository) CreateUser(user model.User) *model.User {
	user.Id = len(m.users)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	m.users = append(m.users, &user)
	return &user
}

func (m *MockRepository) GetUser(id int) (*model.User, error) {
	for _, user := range m.users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *MockRepository) UpdateUser(id int, user model.User) (*model.User, error) {
	for i, u := range m.users {
		if u.Id == id {
			user.Id = id
			user.CreatedAt = u.CreatedAt
			user.UpdatedAt = time.Now()
			m.users[i] = &user
			return &user, nil
		}
	}
	return nil, errors.New("cant update user")
}

func (m *MockRepository) DeleteUser(id int) error {
	for i, user := range m.users {
		if user.Id == id {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return nil
		}
	}
	return errors.New("cant delete user")
}

func (m *MockRepository) GetAllUsers() []*model.User {
	return m.users
}
