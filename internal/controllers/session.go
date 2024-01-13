package controllers

import (
	"github.com/malinatrash/insta-rest-api/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/internal/models"
	"github.com/oklog/ulid/v2"
)

func CreateSession(user models.User) models.Session {
	curr := time.Now().Add(time.Hour)
	println(curr.String())
	session := models.Session{
		UserRefer:      int(user.ID),
		Session:        ulid.Make().String(),
		ExpireDateTime: curr,
	}
	return session
}

// GetSessionByUsernameAndPassword возвращает id Сессии по имени пользователя и паролю
// @Summary Получить пользователя по имени пользователя и паролю
// @Tags sessions
// @Accept json
// @Produce json
// @Param username query string true "Имя пользователя"
// @Param password query string true "Пароль"
// @Success 200 {object} models.SessionResponse
// @Failure 400 {object} models.ErrorMessage
// @Failure 404 {object} models.ErrorMessage
// @Router /session [get]
func GetSessionByUsernameAndPassword(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user models.User
	var session models.Session
	result := database.DB.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильное имя пользователя или пароль"})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	result = database.DB.Where("user_refer = ?", user.ID).First(&session)
	if result.Error != nil {
		newSession := CreateSession(user)
		database.DB.Create(&newSession)
		c.JSON(http.StatusOK, newSession.Session)
		return
	}
	if time.Now().Sub(session.ExpireDateTime) < 0 {
		// Сессия не истекла
		c.JSON(http.StatusOK, session.Session)
	} else {
		// Сессия истекла
		database.DB.Delete(&session)
		newSession := CreateSession(user)
		database.DB.Create(&newSession)
		c.JSON(http.StatusOK, gin.H{"session": newSession.Session})
	}

}
