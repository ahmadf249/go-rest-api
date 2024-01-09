package controllers

import (
	"net/http"
	"test/go-rest-api/database"
	"test/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article *models.Article
	err := c.ShouldBind(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(article)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create article",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

/* func getImages(c *gin.Context) {
c.IndentedJSON(http.StatusOK, images)
}
*/

/* func getImagesByID(c *gin.Context) {
	publicId := c.Param("publicId")

	for _, a := range images {
		if a.PublicID == publicId {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "images not found"})
}
*/
