package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func findFileByNameWithoutExtension(directoryPath, fileName string) (string, error) {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return "", err
	}

	fmt.Println(fileName)
	for _, file := range files {
		fmt.Println(file.Name())
		nameWithoutExt := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		if nameWithoutExt == fileName {
			return filepath.Join(directoryPath, file.Name()), nil
		}
	}

	return "", fmt.Errorf("File not found: %s", fileName)
}

func GetImageForUser(c *gin.Context) {
	imageName := c.Query("username")
	imagePath := "./assets/"
	if imageName == "" {
		c.JSON(404, gin.H{"error": "username required"})
		return
	}

	foundImagePath, err := findFileByNameWithoutExtension(imagePath, imageName)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.File(foundImagePath)
}
func UploadImageForUser(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image not found in request"})
		return
	}

	username := c.PostForm("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not provided"})
		return
	}

	imagePath := "./assets/" + username + filepath.Ext(file.Filename)

	if err := deleteFileByNameWithoutExtension("./assets/", username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete existing image"})
		return
	}

	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}

func deleteFileByNameWithoutExtension(directoryPath, fileName string) error {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		nameWithoutExt := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		if nameWithoutExt == fileName {
			err := os.Remove(filepath.Join(directoryPath, file.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
