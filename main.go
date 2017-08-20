package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	SetUpDB()

	router := gin.New()
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	})

	router.POST("/company/new", Create)
	router.GET("/company/:profile_id", Fetch)

	router.Run(":" + port)
}
