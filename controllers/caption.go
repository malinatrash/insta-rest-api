package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/database"
	"github.com/malinatrash/insta-rest-api/models"
)

func AddCaption(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var captionData struct {
		Caption string `json:"caption"`
	}
	if err := c.BindJSON(&captionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	user.Caption = captionData.Caption

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Caption updated successfully", "user": user})
}
