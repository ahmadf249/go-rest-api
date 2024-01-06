package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type image struct {
	PublicID string `json:"publicId"`
	Tags     string `json:"tags"`
	Source   string `json:"src"`
	Filename string `json:"filename"`
}

var images = []image{
	{PublicID: "wl4go0mhaxymy7g", Tags: "Blue Dragon Awards", Source: "/images/", Filename: "eunsoo-1.jpg"},
	{PublicID: "y0ld2emktilga35", Tags: "Twinkling Watermelon", Source: "/images/", Filename: "eunsoo-2.jpg"},
	{PublicID: "unf96lz4r6og9zn", Tags: "APAN Awards", Source: "/images/", Filename: "eunsoo-3.jpg"},
}

func main() {
	router := gin.Default()
	router.GET("/images", getImages)
	router.GET("/images/:publicId", getImagesByID)

	router.Run("localhost:8080")
}

func getImages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, images)
}

func getImagesByID(c *gin.Context) {
	publicId := c.Param("publicId")

	for _, a := range images {
		if a.PublicID == publicId {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "images not found"})
}
