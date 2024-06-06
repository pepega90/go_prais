package handler

import (
	"fmt"
	"go_prais/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	db = model.New()
)

func GetAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, db.GetAll())
}

func CreateUser(c *gin.Context) {
	var req model.User
	req.Id = len(db.DB) + 1
	req.CreatedAt = time.Now().UTC()
	req.UpdatedAt = time.Now().UTC()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.CreateUser(&req)
	c.JSON(http.StatusOK, req)
}

func UpdateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.UpdateUser(&req)
	c.JSON(http.StatusOK, req)
}

func GetUser(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.Param("id"))
	user := db.GetUserByID(idUser)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.Param("id"))
	db.DeleteUser(idUser)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Berhasil hapus user dengan id %d", idUser)})
}
