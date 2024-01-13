package controllers

import (
	"github.com/malinatrash/insta-rest-api/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/malinatrash/insta-rest-api/docs"
	"github.com/malinatrash/insta-rest-api/internal/models"
)

// AddCaption добавляет или обновляет описание пользователя.
// @Summary Добавляет или обновляет описание пользователя.
// @Description Этот эндпоинт позволяет добавлять или обновлять описание пользователя.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Param models.captionData body models.CaptionData true "Данные для обновления подписи"
// @Success 200 {object} models.SuccessMessage "OK"
// @Failure 400 {object} models.ErrorMessage "Invalid JSON data"
// @Failure 404 {object} models.ErrorMessage "User not found"
// @Failure 500 {object} models.ErrorMessage "Failed to update user"
// @Router  /users/caption/{id} [post]
func AddCaption(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var captionData models.CaptionData
	if err := c.BindJSON(&captionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	user.Caption = captionData.Caption

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Caption updated successfully"})
}
