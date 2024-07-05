package service

import (
	"assignment_1/entity"
	"context"
	"encoding/json"
	"log"
)

// ISubmissionService mendefinisikan interface untuk layanan submission
type ISubmissionService interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) error
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
	GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context, limit, offset int) ([]entity.Submission, error)
	GetTotalSubs(ctx context.Context) (int, error)
}

// ISubmissionRepository mendefinisikan interface untuk repository submission
type ISubmissionRepository interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
	GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context, limit, offset int) ([]entity.Submission, error)
	GetTotalSubs(ctx context.Context) (int, error)
}

// submissionService adalah implementasi dari ISubmissionService yang menggunakan ISubmissionRepository
type submissionService struct {
	submissionRepo ISubmissionRepository
}

// NewSubmissionService membuat instance baru dari submissionService
func NewSubmissionService(submissionRepo ISubmissionRepository) ISubmissionService {
	return &submissionService{submissionRepo: submissionRepo}
}

func (s *submissionService) GetTotalSubs(ctx context.Context) (int, error) {
	c, err := s.submissionRepo.GetTotalSubs(ctx)
	if err != nil {
		log.Println("[service] error get total submission")
		return 0, err
	}
	return c, nil

}

func (s *submissionService) CreateSubmission(ctx context.Context, submission *entity.Submission) error {
	//TODO implement me
	// input validation
	var ans []entity.Answer
	json.Unmarshal(submission.Answers, &ans)

	// calculate risk profile
	score, category, definition := CalculateProfileRiskFromAnswers(ans)

	submission.RiskScore = score
	submission.RiskCategory = category
	submission.RiskDefinition = definition

	_, err := s.submissionRepo.CreateSubmission(ctx, submission)
	if err != nil {
		log.Println("error creating submission")
		return err
	}

	return nil
}

func (s *submissionService) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	subs, err := s.submissionRepo.GetSubmissionByID(ctx, id)
	if err != nil {
		log.Println("[service] error get submissoin by id")
		return entity.Submission{}, err
	}
	return subs, nil
}

func (s *submissionService) GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error) {
	subs, err := s.submissionRepo.GetSubmissionByUserID(ctx, id)
	if err != nil {
		log.Println("[service] error get submissoin by user_id")
		return entity.Submission{}, err
	}
	return subs, nil
}

func (s *submissionService) DeleteSubmission(ctx context.Context, id int) error {
	err := s.submissionRepo.DeleteSubmission(ctx, id)
	if err != nil {
		log.Println("[service] error delete use submission")
		return err
	}
	return nil
}

func (s *submissionService) GetAllSubmissions(ctx context.Context, limit, offset int) ([]entity.Submission, error) {
	listSubs, err := s.submissionRepo.GetAllSubmissions(ctx, limit, offset)
	if err != nil {
		log.Println("[service] error get all submission")
		return nil, err
	}
	return listSubs, nil
}

// TODO: implement logic for profile risk calculation based on score mapping from entity.RiskMapping
// calculateProfileRiskFromAnswers will be used on submission creation
func CalculateProfileRiskFromAnswers(answers []entity.Answer) (score int, category entity.ProfileRiskCategory, definition string) {
	// TODO: calculate total score from answers
	for _, answer := range answers {
		for _, quest := range entity.Questions {
			if quest.ID == answer.QuestionID {
				for _, option := range quest.Options {
					if option.Answer == answer.Answer {
						score += option.Weight
						break
					}
				}
				break
			}
		}
	}

	// TODO: get category and definition based on total score
	for _, v := range entity.RiskMapping {
		if score >= v.MinScore && score <= v.MaxScore {
			category = v.Category
			definition = v.Definition
			break
		}
	}
	return score, category, definition
}
