package routes

import (
	"github.com/IAmRiteshKoushik/mercury/controllers"
	"github.com/gin-gonic/gin"
)

func EventRoutes(engine *gin.RouterGroup) {
	// User interactions
	engine.GET("/event", controllers.GetEventCatalog)
	engine.GET("/event/:eventId", controllers.GetEventByEventId)

	// TODO: Event bookmarking not enabled for now
	// engine.GET("/event/user/:username/bookmark", )
	// engine.GET("/event/user/:username/purchase")
	// engine.GET("/event/user/:username/visited")

	engine.POST("/event/:eventId/register", controllers.RegisterForEvent)

	// Host interactions
	engine.GET("/event/host/:username", controllers.GetEventsByHostUsername)
	engine.POST("/event/host/new", controllers.CreateEvent)
	engine.POST("/event/host/edit", controllers.EditEvent)
	engine.POST("/event/host/publish", controllers.PublishEvent)
	engine.DELETE("/event/host/remove", controllers.DeleteUnpublishedEvent)

	// Host-Multimedia interactions
	engine.POST("/event/host/poster/upload", controllers.UploadEventPoster)
	engine.DELETE("/event/host/poster/delete", controllers.DeleteEventPoster)

	engine.POST("/event/host/thumbnail/upload", controllers.UploadEventThumbnail)
	engine.DELETE("/event/host/thumbnail/delete", controllers.DeleteEventThumbnail)
}
