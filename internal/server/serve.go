package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func Start(wg *sync.WaitGroup) {
	defer wg.Done()

	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":   "uuid",
			"time": time.Now(),
		})
	})
	r.Run("0.0.0.0:8080")
}
