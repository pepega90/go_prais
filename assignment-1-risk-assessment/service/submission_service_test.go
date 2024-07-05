package service_test

import (
	"assignment_1/entity"
	"assignment_1/service"
	mock_service "assignment_1/test/mock/services"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setupTestSubs(t *testing.T) (context.Context, *gomock.Controller, *mock_service.MockISubmissionRepository, service.ISubmissionService) {
	ctrl := gomock.NewController(t)
	mockSubRepo := mock_service.NewMockISubmissionRepository(ctrl)
	subsHandler := service.NewSubmissionService(mockSubRepo)
	ctx := context.Background()

	return ctx, ctrl, mockSubRepo, subsHandler
}

func TestService_CreateSubmission(t *testing.T) {
	ctx, ctrl, mockSubsRepo, subsService := setupTestSubs(t)
	defer ctrl.Finish()

	subs := &entity.Submission{
		ID:     1,
		UserID: 1,
	}

	t.Run("succesfully create submission", func(t *testing.T) {
		mockSubsRepo.EXPECT().CreateSubmission(ctx, subs).Return(*subs, nil)

		err := subsService.CreateSubmission(ctx, subs)
		assert.NoError(t, err)
	})

	t.Run("gagal create submission", func(t *testing.T) {
		mockSubsRepo.EXPECT().CreateSubmission(ctx, subs).Return(entity.Submission{}, errors.New("error creating submission"))
		err := subsService.CreateSubmission(ctx, subs)
		assert.Contains(t, err.Error(), "error creating submission")
	})
}

func TestCalculateProfileRiskFromAnswers(t *testing.T) {
	entity.Questions = []entity.Question{
		{
			ID: 1,
			Options: []entity.Option{
				{Answer: "A", Weight: 10},
				{Answer: "B", Weight: 20},
			},
		},
		{
			ID: 2,
			Options: []entity.Option{
				{Answer: "C", Weight: 30},
				{Answer: "D", Weight: 40},
			},
		},
	}

	entity.RiskMapping = []entity.ProfileRisk{
		{MinScore: 0, MaxScore: 20, Category: "Low", Definition: "Low risk"},
		{MinScore: 21, MaxScore: 50, Category: "Medium", Definition: "Medium risk"},
		{MinScore: 51, MaxScore: 100, Category: "High", Definition: "High risk"},
	}

	answers := []entity.Answer{
		{QuestionID: 1, Answer: "A"},
		{QuestionID: 2, Answer: "D"},
	}

	score, category, definition := service.CalculateProfileRiskFromAnswers(answers)

	assert.Equal(t, 50, score)
	assert.Equal(t, entity.ProfileRiskCategory("Medium"), category)
	assert.Equal(t, "Medium risk", definition)
}

func TestService_GetTotalSubs(t *testing.T) {
	ctx, ctrl, mockSubsRepo, subsService := setupTestSubs(t)
	defer ctrl.Finish()

	t.Run("get total subsmission", func(t *testing.T) {
		mockSubsRepo.EXPECT().GetTotalSubs(ctx).Return(1, nil)
		c, err := subsService.GetTotalSubs(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 1, c)
	})

	t.Run("cant get total submission", func(t *testing.T) {
		mockSubsRepo.EXPECT().GetTotalSubs(ctx).Return(0, errors.New("error get total submission"))
		c, err := subsService.GetTotalSubs(ctx)
		assert.Equal(t, 0, c)
		assert.Contains(t, err.Error(), "error get total submission")
	})
}

func TestService_GetSubmissionByID(t *testing.T) {
	ctx, ctrl, mockSubsRepo, subsService := setupTestSubs(t)
	defer ctrl.Finish()

	t.Run("sucessfuly get subs by id", func(t *testing.T) {
		mockSubsRepo.EXPECT().GetSubmissionByID(ctx, 1).Return(entity.Submission{}, nil)

		subs, err := subsService.GetSubmissionByID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, entity.Submission{}, subs)
	})

	t.Run("cant get subs by id", func(t *testing.T) {
		mockSubsRepo.EXPECT().GetSubmissionByID(ctx, 0).Return(entity.Submission{}, errors.New("[service] error get submissoin by id"))

		_, err := subsService.GetSubmissionByID(ctx, 0)
		assert.Contains(t, err.Error(), "[service] error get submissoin by id")
	})
}

func TestService_GetSubmissionByUserID(t *testing.T) {
	ctx, ctrl, mockSubsRepo, subsService := setupTestSubs(t)
	defer ctrl.Finish()

	t.Run("sucessfuly get subs by user_id", func(t *testing.T) {
		mockSubsRepo.EXPECT().GetSubmissionByUserID(ctx, 1).Return(entity.Submission{}, nil)

		subs, err := subsService.GetSubmissionByUserID(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, entity.Submission{}, subs)
	})

	t.Run("cant get subs by user_id", func(t *testing.T) {
		mockSubsRepo.EXPECT().GetSubmissionByUserID(ctx, 0).Return(entity.Submission{}, errors.New("[service] error get submissoin by user_id"))

		_, err := subsService.GetSubmissionByUserID(ctx, 0)
		assert.Contains(t, err.Error(), "[service] error get submissoin by user_id")
	})
}

func TestService_DeleteSubmission(t *testing.T) {
	ctx, ctrl, mockSubsRepo, subsService := setupTestSubs(t)
	defer ctrl.Finish()

	t.Run("successfully delete submission", func(t *testing.T) {
		mockSubsRepo.EXPECT().DeleteSubmission(ctx, 1).Return(nil)
		err := subsService.DeleteSubmission(ctx, 1)
		assert.NoError(t, err)
	})

	t.Run("cant delete submission", func(t *testing.T) {
		mockSubsRepo.EXPECT().DeleteSubmission(ctx, 0).Return(errors.New("[service] error delete use submission"))
		err := subsService.DeleteSubmission(ctx, 0)
		assert.Contains(t, err.Error(), "[service] error delete use submission")
	})
}

func TestService_GetAllSubmissions(t *testing.T) {
	// [service] error get all submission
	ctx, ctrl, mockSubsRepo, subsService := setupTestSubs(t)
	defer ctrl.Finish()

	listSubs := []entity.Submission{
		{
			ID:     1,
			UserID: 1,
		},
	}

	t.Run("succesfully get all submissions", func(t *testing.T) {
		mockSubsRepo.EXPECT().GetAllSubmissions(ctx, 1, 1).Return(listSubs, nil)
		subs, err := subsService.GetAllSubmissions(ctx, 1, 1)
		assert.NoError(t, err)
		assert.Equal(t, listSubs, subs)
	})

	t.Run("cant get all submissions", func(t *testing.T) {
		mockSubsRepo.EXPECT().GetAllSubmissions(ctx, 0, 0).Return([]entity.Submission(nil), errors.New("[service] error get all submission"))
		subs, err := subsService.GetAllSubmissions(ctx, 0, 0)
		assert.Equal(t, []entity.Submission(nil), subs)
		assert.Contains(t, err.Error(), "[service] error get all submission")
	})
}
