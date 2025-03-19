package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/IAmRiteshKoushik/mercury/helpers"
	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	if !strings.HasPrefix(c.Request.RequestURI, "/api/v1") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Server could not understand the URL",
		})
		return
	}

	// FIX: Setup AuthMiddleware properly
	// _, err := c.Cookie("access_token")
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Access token not found",
	// 	})
	// 	return
	// }

	c.Next()
}

func HostAuth(c *gin.Context) {

}

func NullifyTokenCookies(c *gin.Context) {

	// Delete cookies after revoking
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refesh_token", "", -1, "/", "", false, true)

	// If there is an error saying that there is no cookie
	_, err := c.Cookie("refresh_token")
	if err == http.ErrNoCookie {
		return
	} else {
		// TODO: Setup logging to DEBUG
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "",
		})
	}
	// err = helpers.RevokeRefreshToken(refreshToken)
	// if err != nil {
	// 	// TODO: Setup loggin for DEBUG
	// }
}

func SetAuthAndRefreshCookies(c *gin.Context, auth, refresh string) {
	// FIX: domain
	domain := "." + "domain.com" // allows domains and subdomains
	c.SetCookie("access_token", auth, int(helpers.AuthTokenValidTime), "", domain, true, true)
	c.SetCookie("refresh_token", refresh, int(helpers.RefreshTokenValidTime), "", domain, true, true)
	// TODO: Log out the cookie setting
	return
}

func SetCsrf(c *gin.Context, csrf string) {
	// FIX: domain
	domain := "." + "domain.com" // allows domains and subdomains
	c.SetCookie("X-CSRF-TOKEN", csrf, int(time.Minute*3), "", domain, true, true)
	return
}
