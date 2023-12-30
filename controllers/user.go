package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/database"
	"github.com/malinatrash/insta-rest-api/models"
	"gorm.io/gorm"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(200, users)
}

func GetUserByUsernameAndPassword(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user models.User
	result := database.DB.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильное имя пользователя или пароль"})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	c.BindJSON(&newUser)

	var existingUser models.User
	if err := database.DB.Where("username = ?", newUser.Username).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{"error": "Username already exists"})
		return
	}

	database.DB.Create(&newUser)
	c.JSON(201, newUser)
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var existingUser models.User
	result := database.DB.First(&existingUser, userID)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}

	database.DB.Model(&existingUser).Updates(updatedUser)
	c.JSON(200, existingUser)
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := database.DB.Delete(&user, userID).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
