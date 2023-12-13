package usercontroller

import (
	"api-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user models.User
	userID := c.Param("id")

	// Cari user berdasarkan ID
	err := models.DB.First(&user, userID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func CreateUser(c *gin.Context) {
	var user models.User

	// Bind data dari request ke struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
