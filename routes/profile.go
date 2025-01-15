package routes

import (
	"github.com/gin-gonic/gin"
)

func ProfileRoutes(engine *gin.RouterGroup) {
	// Fetch profile
	engine.GET("/:username")
	engine.GET("/visit/:username")

	// Edit Profile Details
	engine.POST("/edit/:username")
	engine.POST("/edit/profile-picture/:username")
	engine.DELETE("/edit/profile-picture/:username")
}
