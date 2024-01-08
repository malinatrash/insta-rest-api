package controllers

import (
	"errors"
	"fmt"
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

// GetUserByUsernameAndPassword возвращает пользователя по id сессии
// @Summary Получить пользователя по id сессии
// @Tags users
// @Accept json
// @Produce json
// @Param session query string true "id сессии"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Router /users/login [get]
func GetUserBySession(c *gin.Context) {
	session := c.Query("session")

	var needSession models.Session
	result := database.DB.Where("session = ?", session).First(&needSession)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Сессии не существует"})
		return
	}

	var user models.User
	result = database.DB.Where("ID = ?", needSession.UserRefer).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при поиске пользователя"})
		return
	}
	fmt.Println(user)

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
	session := CreateSession(newUser)
	database.DB.Create(&session)
	c.JSON(201, session.Session)
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
