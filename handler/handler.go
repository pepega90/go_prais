package handler

import (
	"go_prais/model"
	"go_prais/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	GetAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(service services.IUserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.userService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	u := h.userService.CreateUser(req)
	c.JSON(http.StatusOK, u)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	updateUser, err := h.userService.UpdateUser(idParam, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, updateUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userService.GetUser(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	err := h.userService.DeleteUser(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfuly delete user"})
}
