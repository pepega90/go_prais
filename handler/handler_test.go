package handler

import (
	"bytes"
	"encoding/json"
	"go_prais/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUserService := &MockUserService{}
	userHandler := NewUserHandler(mockUserService)

	r := gin.Default()
	r.GET("/", userHandler.GetAllUsers)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var users []*model.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	require.NoError(t, err)
	require.Equal(t, 2, len(users))
}

func TestCreateHandler(t *testing.T) {
	t.Run("success create user", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockService := &MockUserService{}
		userHandler := NewUserHandler(mockService)

		r := gin.Default()
		r.POST("/", userHandler.CreateUser)

		user := model.User{Name: "test user", Email: "test@gmail.com", Password: "test"}
		jsonUser, _ := json.Marshal(user)

		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonUser))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var createdUser model.User
		err := json.Unmarshal(w.Body.Bytes(), &createdUser)
		require.NoError(t, err)
		require.Equal(t, user.Name, createdUser.Name)
		require.Equal(t, user.Email, createdUser.Email)
	})

	t.Run("gagal create user", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		mockService := &MockUserService{}
		userHandler := NewUserHandler(mockService)

		r := gin.Default()
		r.POST("/", userHandler.CreateUser)

		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockUserService{}
	userHandler := NewUserHandler(mockService)

	r := gin.Default()
	r.POST("/:id", userHandler.GetUser)

	req, _ := http.NewRequest(http.MethodPost, "/1", nil)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var u model.User
	err := json.Unmarshal(res.Body.Bytes(), &u)
	require.NoError(t, err)
	require.Equal(t, 1, u.Id)
}

func TestUpdateUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockUserService{}
	userHandler := NewUserHandler(mockService)

	r := gin.Default()
	r.PUT("/:id", userHandler.UpdateUser)

	reqUpdate := model.User{
		Name: "pepeg",
	}
	jsonUser, _ := json.Marshal(reqUpdate)

	req, _ := http.NewRequest(http.MethodPut, "/1", bytes.NewBuffer(jsonUser))

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	require.Equal(t, http.StatusOK, res.Code)
	var updated model.User
	err := json.Unmarshal(res.Body.Bytes(), &updated)
	require.NoError(t, err)
	require.Equal(t, reqUpdate.Name, updated.Name)
}

func TestDeleteUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockUserService{}
	userHandler := NewUserHandler(mockService)

	r := gin.Default()
	r.DELETE("/:id", userHandler.DeleteUser)

	req, _ := http.NewRequest(http.MethodDelete, "/1", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	require.Equal(t, http.StatusOK, res.Code)
}
