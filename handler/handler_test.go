package handler_test

import (
	"errors"
	"go_prais/handler"
	"go_prais/model"
	mock_services "go_prais/test/mock/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserHandler_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockIUserService(ctrl)
	userHandler := handler.NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)

	t.Run("valid create handler", func(t *testing.T) {
		mockService.EXPECT().CreateUser(gomock.Any(), &model.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
		}).Return(model.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
		}, nil)

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"John Doe","email":"john@example.com","password":"password"}`))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		r := gin.Default()
		r.POST("/", userHandler.CreateUser)

		r.ServeHTTP(res, req)

		require.Equal(t, http.StatusOK, res.Code)
		require.JSONEq(t, `{"id":0,"name":"John Doe","email":"john@example.com","password":"password","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, res.Body.String())
	})
	t.Run("InvalidPayload_MissingName", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"email":"john@example.com","password":"password"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/", userHandler.CreateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.JSONEq(t, `{"error":"Name tidak boleh kosong!"}`, resp.Body.String())
	})

	t.Run("InvalidPayload_MissingEmail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"john","password":"password"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/", userHandler.CreateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusBadRequest, resp.Code)
		require.JSONEq(t, `{"error":"E-mail tidak boleh kosong!"}`, resp.Body.String())
	})

	t.Run("ServiceError", func(t *testing.T) {
		mockService.EXPECT().CreateUser(gomock.Any(), &model.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
		}).Return(model.User{}, errors.New("some service error"))

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"John Doe","email":"john@example.com","password":"password"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/", userHandler.CreateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
	})
}

func TestUserHandler_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockIUserService(ctrl)
	userHandler := handler.NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)

	t.Run("InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/abc", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/:id", userHandler.GetUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusNotFound, resp.Code)
		require.JSONEq(t, `{"error":"Invalid ID"}`, resp.Body.String())
	})

	t.Run("UserNotFound", func(t *testing.T) {
		mockService.EXPECT().GetUserByID(gomock.Any(), 1).Return(model.User{}, errors.New("user not found"))

		req := httptest.NewRequest(http.MethodGet, "/1", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/:id", userHandler.GetUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusNotFound, resp.Code)
		require.JSONEq(t, `{"error":"User not found"}`, resp.Body.String())
	})
}

func TestUserHandler_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockIUserService(ctrl)
	userHandler := handler.NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)

	t.Run("InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/abc", strings.NewReader(`{"name":"John Doe","email":"john@example.com"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/:id", userHandler.UpdateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusNotFound, resp.Code)
		require.JSONEq(t, `{"error":"Invalid ID"}`, resp.Body.String())
	})

	t.Run("InvalidPayload", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/1", strings.NewReader(`{"email":"john@example.com"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/:id", userHandler.UpdateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusNotFound, resp.Code)
		require.JSONEq(t, `{"error":"Name tidak boleh kosong!"}`, resp.Body.String())
	})

	t.Run("ServiceError", func(t *testing.T) {
		mockService.EXPECT().UpdateUser(gomock.Any(), 1, model.User{
			Name:  "John Doe",
			Email: "john@example.com",
		}).Return(model.User{}, errors.New("some service error"))

		req := httptest.NewRequest(http.MethodPut, "/1", strings.NewReader(`{"name":"John Doe","email":"john@example.com"}`))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/:id", userHandler.UpdateUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusNotFound, resp.Code)
		require.JSONEq(t, `{"error":"some service error"}`, resp.Body.String())
	})
}

func TestUserHandler_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockIUserService(ctrl)
	userHandler := handler.NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)

	t.Run("InvalidID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/abc", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.DELETE("/:id", userHandler.DeleteUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusNotFound, resp.Code)
		require.JSONEq(t, `{"error":"Invalid ID"}`, resp.Body.String())
	})

	t.Run("ServiceError", func(t *testing.T) {
		mockService.EXPECT().DeleteUser(gomock.Any(), 1).Return(errors.New("some service error"))

		req := httptest.NewRequest(http.MethodDelete, "/1", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.DELETE("/:id", userHandler.DeleteUser)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusNotFound, resp.Code)
		require.JSONEq(t, `{"error":"some service error"}`, resp.Body.String())
	})
}

func TestUserHandler_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockIUserService(ctrl)
	userHandler := handler.NewUserHandler(mockService)

	gin.SetMode(gin.TestMode)

	t.Run("ValidRequest", func(t *testing.T) {
		mockService.EXPECT().GetAllUsers(gomock.Any()).Return([]model.User{
			{Id: 1, Name: "John Doe", Email: "john@example.com"},
			{Id: 2, Name: "Jane Doe", Email: "jane@example.com"},
		}, nil)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/", userHandler.GetAllUsers)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
		require.JSONEq(t, `[{"id":1,"name":"John Doe","email":"john@example.com","password":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},{"id":2,"name":"Jane Doe","email":"jane@example.com","password":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`, resp.Body.String())
	})

	t.Run("ServiceError", func(t *testing.T) {
		mockService.EXPECT().GetAllUsers(gomock.Any()).Return(nil, errors.New("some service error"))

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/", userHandler.GetAllUsers)

		router.ServeHTTP(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
	})
}
