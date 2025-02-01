package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api/v1") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Server could not understand the URL",
			})
			c.Abort()
			return
		}

		_, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Access token not found",
			})
			c.Abort()
			return
		}
	}
}

func NullifyTokenCookies(c *gin.Context) {

}

func SetAuthAndRefreshCookies(c *gin.Context) {

}

func GrabCsrfFromReq(c *gin.Context) {

}
