package handler

import (
	"assignment_1/entity"
	"assignment_1/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
)

type PaginationResponse struct {
	Users       []entity.User `json:"users"`
	TotalPages  int           `json:"total_pages"`
	CurrentPage int           `json:"current_page"`
}

// IUserHandler mendefinisikan interface untuk handler user
type IUserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	userService service.IUserService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Name  string `json:"name" valid:"Required"`
		Email string `json:"email" valid:"Required;Email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error creating user"})
		return
	}

	valid := validation.Validation{}
	validasi, err := valid.Valid(&req)

	if !validasi {
		errorMessages := getErrorMsg(valid.Errors, valid)
		c.JSON(http.StatusBadRequest, errorMessages)
		return
	}

	userReq := &entity.User{
		Name:  req.Name,
		Email: req.Email,
	}
	createdUser, err := u.userService.CreateUser(c, userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error creating user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("user id %v created successfully", createdUser.ID)})
}

func (u *UserHandler) GetUser(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	getUser, err := u.userService.GetUserByID(c, idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("error get user with id = %v", idParam)})
		return
	}
	c.JSON(http.StatusOK, getUser)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Name  string `json:"name" valid:"Required"`
		Email string `json:"email" valid:"Required;Email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error updating user"})
		return
	}

	valid := validation.Validation{}
	validasi, err := valid.Valid(&req)

	if !validasi {
		errorMessages := getErrorMsg(valid.Errors, valid)
		c.JSON(http.StatusBadRequest, errorMessages)
		return
	}

	userReq := entity.User{
		Name:  req.Name,
		Email: req.Email,
	}
	updatedUser, err := u.userService.UpdateUser(c, idParam, userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error updating user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user id %v updated successfully", updatedUser.ID)})
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	err := u.userService.DeleteUser(c, idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error deleting user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user id %v deleted successfully", idParam)})
}

func (u *UserHandler) GetAllUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "2")

	page, err := strconv.Atoi(pageStr)
	limit, err := strconv.Atoi(limitStr)

	// Get the total count of users
	totalCount, err := u.userService.GetUserCount(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error get user count"})
		return
	}

	totalPages := (totalCount + limit - 1) / limit
	offset := (page - 1) * limit

	listUsers, err := u.userService.GetAllUsers(c, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error get all users"})
		return
	}

	response := PaginationResponse{
		Users:       listUsers,
		TotalPages:  totalPages,
		CurrentPage: page,
	}

	c.JSON(http.StatusOK, response)
}

func getErrorMsg(validErr []*validation.Error, valid validation.Validation) []map[string]string {
	errorMessages := []map[string]string{}
	for _, err := range valid.Errors {
		errorMessages = append(errorMessages, map[string]string{err.Field: err.Message})
	}
	return errorMessages
}
