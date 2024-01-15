package main

import (
	"test/go-rest-api/controllers"
	"test/go-rest-api/database"
	"test/go-rest-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	database.DatabaseConnection()
}

func main() {
	router := gin.Default()
	// router.GET("/images", getImages)
	// router.GET("/images/:publicId", getImagesByID)

	router.POST("api/article", controllers.CreateArticle)
	router.POST("api/image", controllers.InsertImage)
	router.GET("api/image/:PublicId", controllers.ViewImage)
	router.GET("api/image", controllers.ViewImages)
	router.DELETE("api/image/:PublicId", controllers.DeleteImages)

	router.Run()
}
