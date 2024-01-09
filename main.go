package main

import (
	"test/go-rest-api/controllers"
	"test/go-rest-api/database"

	"github.com/gin-gonic/gin"
)

// var images = []image{
// 	{PublicID: "wl4go0mhaxymy7g", Tags: "Blue Dragon Awards", Source: "/images/", Filename: "eunsoo-1.jpg"},
// 	{PublicID: "y0ld2emktilga35", Tags: "Twinkling Watermelon", Source: "/images/", Filename: "eunsoo-2.jpg"},
// 	{PublicID: "unf96lz4r6og9zn", Tags: "APAN Awards", Source: "/images/", Filename: "eunsoo-3.jpg"},
// }

func main() {
	database.DatabaseConnection()
	router := gin.Default()
	// router.GET("/images", getImages)
	// router.GET("/images/:publicId", getImagesByID)

	router.POST("/article", controllers.CreateArticle)
	router.POST("/image", controllers.InsertImage)

	router.Run("localhost:5000")
}
