package controllers

import (
	"errors"
	"net/http"
	"test/go-rest-api/database"
	"test/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func InsertImage(c *gin.Context) {
	var image *models.Image
	err := c.ShouldBind(&image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(image)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create article",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"image": image,
	})
}

func ViewImage(c *gin.Context) {
	var image *models.Image
	publicId := c.Param("publicId")
	res := database.DB.Find(&image, publicId)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "image not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"image": image,
	})
}

func ViewImages(c *gin.Context) {
	var images []models.Image
	res := database.DB.Find(&images)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("images not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"images": images,
	})
}

func DeleteImages(c *gin.Context) {
	var image models.Image
	publicId := c.Param("publicId")
	res := database.DB.Find(&image, publicId)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "image not found",
		})
		return
	}
	database.DB.Delete(&image)
	c.JSON(http.StatusOK, gin.H{
		"message": "image deleted successfully",
	})
}
