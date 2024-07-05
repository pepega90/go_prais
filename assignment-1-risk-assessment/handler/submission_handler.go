package handler

import (
	"assignment_1/entity"
	"assignment_1/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
)

// ISubmissionHandler mendefinisikan interface untuk handler submission
type ISubmissionHandler interface {
	CreateSubmission(c *gin.Context)
	GetSubmission(c *gin.Context)
	GetAllSubmissions(c *gin.Context)
	DeleteSubmission(c *gin.Context)
}

type PaginationResponseSubs struct {
	TotalPages  int                 `json:"total"`
	CurrentPage int                 `json:"page"`
	Limit       int                 `json:"limit"`
	Subs        []entity.Submission `json:"submissions"`
}

// NewSubmissionHandler membuat instance baru dari SubmissionHandler
func NewSubmissionHandler(submissionService service.ISubmissionService) ISubmissionHandler {
	return &SubmissionHandler{
		submissionService: submissionService,
	}
}

type SubmissionHandler struct {
	submissionService service.ISubmissionService
}

func (s *SubmissionHandler) CreateSubmission(c *gin.Context) {
	var req struct {
		UserID  int             `json:"user_id" valid:"Required"`
		Answers []entity.Answer `json:"answers" valid:"Required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error create submission"})
		return
	}

	valid := validation.Validation{}
	validasi, _ := valid.Valid(&req)

	if !validasi {
		errorMessages := getErrorMsg(valid.Errors, valid)
		c.JSON(http.StatusBadRequest, errorMessages)
		return
	}

	answersJSON, _ := json.Marshal(req.Answers)

	e := entity.Submission{
		UserID:  req.UserID,
		Answers: answersJSON,
	}

	err := s.submissionService.CreateSubmission(c, &e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("error create submission: %v", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("submission id %v created successfully", e.ID),
	})
}

func (s *SubmissionHandler) GetSubmission(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	idParam, _ := strconv.Atoi(c.Param("id"))

	if user_id != 0 {
		subs, err := s.submissionService.GetSubmissionByUserID(c, user_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error get submission by user id"})
			return
		}
		c.JSON(http.StatusOK, subs)
	} else if idParam != 0 {
		subs, err := s.submissionService.GetSubmissionByID(c, idParam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error get submission id"})
			return
		}
		c.JSON(http.StatusOK, subs)
	}
}

func (s *SubmissionHandler) DeleteSubmission(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	err := s.submissionService.DeleteSubmission(c, idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error deleting submission"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("deleting submission id %v successfully", idParam)})
}

func (s *SubmissionHandler) GetAllSubmissions(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "2")

	page, err := strconv.Atoi(pageStr)
	limit, err := strconv.Atoi(limitStr)

	// Get the total count of users
	totalCount, err := s.submissionService.GetTotalSubs(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error get user count"})
		return
	}

	totalPages := (totalCount + limit - 1) / limit
	offset := (page - 1) * limit

	listSubs, err := s.submissionService.GetAllSubmissions(c, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error get all submission"})
		return
	}

	res := PaginationResponseSubs{
		TotalPages:  totalPages,
		CurrentPage: page,
		Limit:       limit,
		Subs:        listSubs,
	}

	c.JSON(http.StatusOK, res)
}
