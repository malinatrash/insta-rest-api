package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/database"
	"github.com/malinatrash/insta-rest-api/models"
	"gorm.io/gorm"
)

// GetAllUsers возвращает список всех пользователей
// @Summary Получить всех пользователей
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users/ [get]
func GetAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(200, users)
}

// GetUserByUsernameAndPassword возвращает пользователя по имени пользователя и паролю
// @Summary Получить пользователя по имени пользователя и паролю
// @Tags users
// @Accept json
// @Produce json
// @Param username query string true "Имя пользователя"
// @Param password query string true "Пароль"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Router /users/login [get]
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

// GetUserByID возвращает пользователя по ID
// @Summary Получить пользователя по ID
// @Description Получить пользователя по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 200 {object} models.User
// @Failure 404 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /users/{id} [get]
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

// CreateUser создает нового пользователя
// @Summary Создать нового пользователя
// @Description Создает нового пользователя с предоставленными данными
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "Объект пользователя"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorMessage
// @Router /users/ [post]
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

// UpdateUser обновляет данные пользователя по ID
// @Summary Обновить данные пользователя
// @Description Обновляет данные существующего пользователя по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Param user body models.User true "Обновленный объект пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Router /users/{id} [put]
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

// DeleteUser удаляет пользователя по ID
// @Summary Удалить пользователя по ID
// @Description Удаляет пользователя по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 200 {object} models.SuccessMessage
// @Failure 404 {object} models.ErrorMessage
// @Failure 500 {object} models.ErrorMessage
// @Router /users/{id} [delete]
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
