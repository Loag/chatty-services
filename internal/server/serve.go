package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":   "uuid",
			"time": time.Now(),
		})
	})
	r.Run("0.0.0.0:8080")
}
