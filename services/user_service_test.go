package services

import (
	"go_prais/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockRepo := &MockRepository{}
	userService := NewUserService(mockRepo)

	t.Run("create user", func(t *testing.T) {
		u := model.User{Name: "pepeg", Email: "pepeg@handsome.com", Password: "pass"}
		createdUser := userService.CreateUser(u)

		assert.Equal(t, "pepeg", createdUser.Name)
		assert.Equal(t, "pepeg@handsome.com", createdUser.Email)
	})
}

func TestGetUser(t *testing.T) {
	mockRepo := &MockRepository{}
	userService := NewUserService(mockRepo)

	createdUser := model.User{Name: "pepeg", Email: "pepeg@handsome.com", Password: "pass"}
	userService.CreateUser(createdUser)

	u, err := userService.GetUser(0)

	assert.NoError(t, err)
	assert.Equal(t, "pepeg", u.Name)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := &MockRepository{}
	userService := NewUserService(mockRepo)

	createdUser := model.User{Name: "pepeg", Email: "pepeg@handsome.com", Password: "pass"}
	userService.CreateUser(createdUser)

	updated := model.User{Name: "aji"}
	updatedUser, err := userService.UpdateUser(0, updated)

	assert.NoError(t, err)
	assert.Equal(t, "aji", updatedUser.Name)
}

func TestDelete(t *testing.T) {
	mockRepo := &MockRepository{}
	userService := NewUserService(mockRepo)

	createdUser := model.User{Name: "pepeg", Email: "pepeg@handsome.com", Password: "pass"}
	userService.CreateUser(createdUser)

	err := userService.DeleteUser(0)
	assert.NoError(t, err)
}
