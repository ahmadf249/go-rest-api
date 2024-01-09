package controllers

import (
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
