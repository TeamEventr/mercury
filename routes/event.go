package routes

import (
	"github.com/gin-gonic/gin"
)

func EventRoutes(engine *gin.RouterGroup) {
	// User interactions
	engine.GET("/event/user")
	engine.GET("/event/:eventId/user")
	engine.GET("/event/user/:username/bookmark")
	engine.GET("/event/user/:username/purchase")
	engine.GET("/event/user/:username/visited")
	engine.POST("/event/user/register")

	// Host interactions
	engine.GET("/event/host/:username")
	engine.POST("/event/host/new")
	engine.POST("/event/host/edit")
	engine.POST("/event/host/publish")

	// Host-Multimedia interactions
}
