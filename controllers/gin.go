package controllers

import (
	"net/http"
	"test/go-rest-api/database"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article *database.Article
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
	return
}
