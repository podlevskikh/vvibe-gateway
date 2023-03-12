package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello world")
	})

	router.GET("/openapi.json", func(c *gin.Context) { http.ServeFile(c.Writer, c.Request, "./openapi.json") })

	router.Run(":" + port)
}
