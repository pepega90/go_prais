package handler

import (
	"go_prais/model"
	"go_prais/services"
	"net/http"
	"strconv"
	"strings"

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
	users, _ := h.userService.GetAllUsers(c.Request.Context())
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := err.Error()
		errMsg = convertUserMandatoryFieldErrorString(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}
	u, _ := h.userService.CreateUser(c.Request.Context(), &req)
	c.JSON(http.StatusOK, u)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	updateUser, err := h.userService.UpdateUser(c.Request.Context(), idParam, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updateUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userService.GetUserByID(c.Request.Context(), idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	err := h.userService.DeleteUser(c.Request.Context(), idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfuly delete user"})
}

func convertUserMandatoryFieldErrorString(oldErrorMsg string) string {
	switch {
	case strings.Contains(oldErrorMsg, "'Name' failed on the 'required' tag"):
		return "Name tidak boleh kosong!"
	case strings.Contains(oldErrorMsg, "'Email' failed on the 'required' tag"):
		return "E-mail tidak boleh kosong!"
	}
	return oldErrorMsg
}
