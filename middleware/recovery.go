package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Panic("Recovered!")
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}
