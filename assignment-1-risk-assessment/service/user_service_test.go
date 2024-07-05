package service_test

import (
	"assignment_1/entity"
	"assignment_1/service"
	mock_service "assignment_1/test/mock/services"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setupTest(t *testing.T) (context.Context, *gomock.Controller, *mock_service.MockIUserRepository, service.IUserService) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_service.NewMockIUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)
	ctx := context.Background()

	return ctx, ctrl, mockRepo, userService
}

func Test_CreateUser(t *testing.T) {
	ctx, ctrl, mockRepo, userService := setupTest(t)
	defer ctrl.Finish()

	user := &entity.User{
		Name:      "pepeg",
		Email:     "pepeg@handsome.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("successfully create user", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, user).Return(*user, nil)

		createdUser, err := userService.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.Equal(t, *user, createdUser)
	})

	t.Run("cant create user", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, user).Return(entity.User{}, errors.New("error create user"))

		createdUser, err := userService.CreateUser(ctx, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error create user")
		assert.Equal(t, entity.User{}, createdUser)
	})
}

func Test_GetUserCount(t *testing.T) {
	ctx, ctrl, mockRepo, userService := setupTest(t)
	defer ctrl.Finish()

	t.Run("get count user", func(t *testing.T) {
		mockRepo.EXPECT().GetUserCount(ctx).Return(2, nil)

		count, err := userService.GetUserCount(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 2, count)
	})

	t.Run("cant get user count", func(t *testing.T) {
		mockRepo.EXPECT().GetUserCount(ctx).Return(0, errors.New("cant get user count"))

		count, err := userService.GetUserCount(ctx)
		assert.Error(t, err)
		assert.Equal(t, 0, count)
		assert.Contains(t, err.Error(), "cant get user count")
	})
}

func Test_GetUserByID(t *testing.T) {
	ctx, ctrl, mockRepo, userService := setupTest(t)
	defer ctrl.Finish()

	user := &entity.User{
		ID:    1,
		Name:  "pepeg",
		Email: "pepeg@handsome.com",
	}

	t.Run("successfully get user by id", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByID(ctx, 1).Return(*user, nil)

		getUser, err := userService.GetUserByID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, *user, getUser)
	})

	t.Run("cant get user by id", func(t *testing.T) {
		mockRepo.EXPECT().GetUserByID(ctx, 0).Return(entity.User{}, errors.New("cant user with id 0"))

		getUser, err := userService.GetUserByID(ctx, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cant user with id 0")
		assert.Equal(t, entity.User{}, getUser)
	})
}

func Test_UpdateUser(t *testing.T) {
	ctx, ctrl, mockRepo, userService := setupTest(t)
	defer ctrl.Finish()

	user := &entity.User{
		ID:    1,
		Name:  "pepeg",
		Email: "pepeg@handsome.com",
	}

	t.Run("successfully update user", func(t *testing.T) {
		updatedReq := entity.User{
			Name:  "pepeg update",
			Email: "pepeg_update@handsome.com",
		}
		mockRepo.EXPECT().UpdateUser(ctx, user.ID, updatedReq).Return(updatedReq, nil)

		updatedUser, err := userService.UpdateUser(ctx, user.ID, updatedReq)
		assert.NoError(t, err)
		assert.Equal(t, updatedReq, updatedUser)
	})

	t.Run("cant update user", func(t *testing.T) {
		mockRepo.EXPECT().UpdateUser(ctx, 0, entity.User{}).Return(entity.User{}, errors.New("user not found"))

		updatedUser, err := userService.UpdateUser(ctx, 0, entity.User{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
		assert.Equal(t, entity.User{}, updatedUser)
	})
}

func Test_DeleteUser(t *testing.T) {
	ctx, ctrl, mockRepo, userService := setupTest(t)
	defer ctrl.Finish()

	user := entity.User{
		ID:    1,
		Name:  "pepeg",
		Email: "pepeg@handsome.com",
	}

	t.Run("successfully delete user", func(t *testing.T) {
		mockRepo.EXPECT().DeleteUser(ctx, user.ID).Return(nil)

		err := userService.DeleteUser(ctx, user.ID)
		assert.NoError(t, err)
	})

	t.Run("error while delete user", func(t *testing.T) {
		mockRepo.EXPECT().DeleteUser(ctx, 0).Return(errors.New("cant delete user with id 0"))

		err := userService.DeleteUser(ctx, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cant delete user with id 0")
	})
}

func Test_GetAllUsers(t *testing.T) {
	ctx, ctrl, mockRepo, userService := setupTest(t)
	defer ctrl.Finish()

	listUser := []entity.User{
		{
			ID:    1,
			Name:  "pepeg",
			Email: "pepeg@handsome.com",
		},
	}

	t.Run("successfully get all users", func(t *testing.T) {
		mockRepo.EXPECT().GetAllUsers(ctx, 1, 1).Return(listUser, nil)

		userList, err := userService.GetAllUsers(ctx, 1, 1)
		assert.NoError(t, err)
		assert.Len(t, userList, 1)
		assert.Equal(t, listUser, userList)
	})

	t.Run("cant get all users", func(t *testing.T) {
		mockRepo.EXPECT().GetAllUsers(ctx, 0, 0).Return([]entity.User(nil), errors.New("user not found"))

		userList, err := userService.GetAllUsers(ctx, 0, 0)
		assert.Error(t, err)
		assert.Len(t, userList, 0)
		assert.Equal(t, []entity.User(nil), userList)
		assert.Contains(t, err.Error(), "user not found")
	})
}
