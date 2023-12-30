package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func saveImage(c *gin.Context, file *multipart.FileHeader) (string, error) {
	imageName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
	uploadPath := "../assets/"
	dst := filepath.Join(uploadPath, imageName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return "", err
	}
	return imageName, nil
}

func CreateImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageName, err := saveImage(c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_name": imageName})
}

func GetImageForUser(c *gin.Context) {
	imageName := c.Param("username")
	imagePath := "../assets/"
	listFilesInDirectory("../assets")

	if _, err := os.Stat(filepath.Join(imagePath, imageName)); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.File(filepath.Join(imagePath, imageName))
}

func listFilesInDirectory(directoryPath string) {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Files in", directoryPath)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
}
