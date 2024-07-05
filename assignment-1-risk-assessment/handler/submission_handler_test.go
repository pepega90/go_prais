package handler_test

import (
	"assignment_1/entity"
	"assignment_1/handler"
	mock_service "assignment_1/test/mock/services"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setupTestSubs(t *testing.T) (context.Context, *gomock.Controller, *mock_service.MockISubmissionService, handler.ISubmissionHandler) {
	ctrl := gomock.NewController(t)
	mockService := mock_service.NewMockISubmissionService(ctrl)
	subsHandler := handler.NewSubmissionHandler(mockService)
	ctx := context.Background()

	return ctx, ctrl, mockService, subsHandler
}

func TestSubHandler_CreateSubmission(t *testing.T) {
	_, ctrl, mockService, subsHandler := setupTestSubs(t)
	defer ctrl.Finish()

	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.POST("/submissions", subsHandler.CreateSubmission)

	t.Run("successfully create submission", func(t *testing.T) {
		reqBody := map[string]interface{}{
			"user_id": 1,
			"answers": []entity.Answer{
				{QuestionID: 1, Answer: "Answer 1"},
				{QuestionID: 2, Answer: "Answer 2"},
			},
		}

		reqJSON, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/submissions", bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		// Setup the mock expectation
		mockService.EXPECT().CreateSubmission(gomock.Any(), gomock.Any()).Return(nil).Times(1)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "created successfully")
	})

	t.Run("cant create submission", func(t *testing.T) {
		reqBody := map[string]interface{}{
			"user_id": 1,
			"answers": []entity.Answer{
				{QuestionID: 1, Answer: "Answer 1"},
				{QuestionID: 2, Answer: "Answer 2"},
			},
		}

		reqJSON, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/submissions", bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		// Setup the mock expectation
		mockService.EXPECT().CreateSubmission(gomock.Any(), gomock.Any()).Return(fmt.Errorf("some error")).Times(1)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "error create submission: some error")
	})

	t.Run("invalid request body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/submissions", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error create submission")
	})

	t.Run("error validasi", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{})
		req, _ := http.NewRequest(http.MethodPost, "/submissions", bytes.NewBuffer(reqBody))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "UserID Can not be empty")
	})
}

func TestSubHandler_GetSubmission(t *testing.T) {
	_, ctrl, mockService, subsHandler := setupTestSubs(t)
	defer ctrl.Finish()

	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.GET("/submissions/:id", subsHandler.GetSubmission)

	t.Run("get submission by id", func(t *testing.T) {
		userId := 1
		sub := entity.Submission{
			ID:     1,
			UserID: userId,
		}

		mockService.EXPECT().GetSubmissionByID(gomock.Any(), userId).Return(sub, nil).Times(1)

		req, _ := http.NewRequest(http.MethodGet, "/submissions/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("cant get sub by id", func(t *testing.T) {
		mockService.EXPECT().GetSubmissionByID(gomock.Any(), 1).Return(entity.Submission{}, errors.New("error get submission id")).Times(1)

		req, _ := http.NewRequest(http.MethodGet, "/submissions/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("get submission by user_id", func(t *testing.T) {
		userId := 1
		sub := entity.Submission{
			ID:     1,
			UserID: userId,
		}

		mockService.EXPECT().GetSubmissionByUserID(gomock.Any(), userId).Return(sub, nil).Times(1)

		req, _ := http.NewRequest(http.MethodGet, "/submissions/0?user_id=1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("cant get sub by user_id", func(t *testing.T) {
		mockService.EXPECT().GetSubmissionByUserID(gomock.Any(), 1).Return(entity.Submission{}, errors.New("error get submission id")).Times(1)

		req, _ := http.NewRequest(http.MethodGet, "/submissions/0?user_id=1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestSubHandler_DeleteSubmission(t *testing.T) {
	_, ctrl, mockService, subsHandler := setupTestSubs(t)
	defer ctrl.Finish()

	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.DELETE("/submissions/:id", subsHandler.DeleteSubmission)

	t.Run("success delete submission", func(t *testing.T) {
		mockService.EXPECT().DeleteSubmission(gomock.Any(), 1).Return(nil).Times(1)

		req, _ := http.NewRequest(http.MethodDelete, "/submissions/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("cant delete submission", func(t *testing.T) {
		mockService.EXPECT().DeleteSubmission(gomock.Any(), 1).Return(errors.New("error deleting submission")).Times(1)

		req, _ := http.NewRequest(http.MethodDelete, "/submissions/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestSubHandler_GetAllSubmissions(t *testing.T) {
	_, ctrl, mockService, subsHandler := setupTestSubs(t)
	defer ctrl.Finish()

	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.GET("/submissions", subsHandler.GetAllSubmissions)

	t.Run("success", func(t *testing.T) {
		page := 1
		limit := 2
		totalCount := 5
		expectedSubmissions := []entity.Submission{
			{ID: 1, UserID: 1},
			{ID: 2, UserID: 2},
		}

		// Setup the mock expectations
		mockService.EXPECT().GetTotalSubs(gomock.Any()).Return(totalCount, nil).Times(1)
		mockService.EXPECT().GetAllSubmissions(gomock.Any(), limit, (page-1)*limit).Return(expectedSubmissions, nil).Times(1)

		req, _ := http.NewRequest(http.MethodGet, "/submissions?page="+strconv.Itoa(page)+"&limit="+strconv.Itoa(limit), nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		// Add assertions for response body if needed
	})

	t.Run("get total count error", func(t *testing.T) {
		page := 1
		limit := 2

		// Setup the mock expectation
		mockService.EXPECT().GetTotalSubs(gomock.Any()).Return(0, fmt.Errorf("some error")).Times(1)

		req, _ := http.NewRequest(http.MethodGet, "/submissions?page="+strconv.Itoa(page)+"&limit="+strconv.Itoa(limit), nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Contains(t, resp.Body.String(), "error get user count")
	})

	t.Run("get all submissions error", func(t *testing.T) {
		page := 1
		limit := 2
		totalCount := 5

		// Setup the mock expectations
		mockService.EXPECT().GetTotalSubs(gomock.Any()).Return(totalCount, nil).Times(1)
		mockService.EXPECT().GetAllSubmissions(gomock.Any(), limit, (page-1)*limit).Return(nil, fmt.Errorf("some error")).Times(1)

		req, _ := http.NewRequest(http.MethodGet, "/submissions?page="+strconv.Itoa(page)+"&limit="+strconv.Itoa(limit), nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Contains(t, resp.Body.String(), "error get all submission")
	})
}
