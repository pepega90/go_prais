package postgres_gorm

import (
	"assignment_1/entity"
	"assignment_1/service"
	"context"
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

type submissionRepository struct {
	db *gorm.DB
}

// NewSubmissionRepository membuat instance baru dari submissionRepository
func NewSubmissionRepository(db *gorm.DB) service.ISubmissionRepository {
	return &submissionRepository{db: db}
}

func (s *submissionRepository) CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error) {
	anwers_list, _ := json.Marshal(submission.Answers)
	submission.Answers = anwers_list
	err := s.db.Create(&submission).Error
	if err != nil {
		log.Println("error creating submission")
		return entity.Submission{}, err
	}
	return *submission, nil
}

func (s *submissionRepository) GetTotalSubs(ctx context.Context) (int, error) {
	var count int64
	err := s.db.Model(&entity.Submission{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (s *submissionRepository) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	var subs entity.Submission
	err := s.db.Find(&subs, id).Error
	if err != nil {
		log.Println("error get subsmission by id")
		return entity.Submission{}, err
	}
	return subs, nil
}

func (s *submissionRepository) GetSubmissionByUserID(ctx context.Context, id int) (entity.Submission, error) {
	var subs entity.Submission
	err := s.db.Find(&subs, "user_id = ?", id).Error
	if err != nil {
		log.Println("error get submission by id")
		return entity.Submission{}, nil
	}
	return subs, nil
}

func (s *submissionRepository) DeleteSubmission(ctx context.Context, id int) error {
	err := s.db.Delete(&entity.Submission{}, id).Error
	if err != nil {
		log.Println("error deleting submission")
		return err
	}
	return nil
}

func (s *submissionRepository) GetAllSubmissions(ctx context.Context, limit, offset int) ([]entity.Submission, error) {
	var listSubs []entity.Submission
	err := s.db.Limit(limit).Offset(offset).Find(&listSubs).Error
	if err != nil {
		log.Println("error get all submissions")
		return nil, err
	}
	return listSubs, nil
}
