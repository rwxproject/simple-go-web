package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	eng := gin.Default()

	eng.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "simple-go-web")
	})

	eng.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "simple-go-web"})
	})

	eng.GET("/_status/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "healthz")
	})

	eng.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
