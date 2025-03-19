package routes

import (
	c "github.com/IAmRiteshKoushik/mercury/controllers/host"
	"github.com/gin-gonic/gin"
)

func HostEventRoutes(engine *gin.RouterGroup) {

	engine.GET("/event/create", c.CreateEventCsrf)
	engine.POST("/event/create", c.CreateEvent)
	engine.GET("/event/:eventId/edit", c.EditEventCsrf)
	engine.POST("/event/:eventId/edit", c.EditEvent)

	engine.GET("/event/:eventId", c.FetchEventById)

	engine.POST("/event/:eventId", c.PublishEvent)
	engine.DELETE("/event/:eventId", c.DeleteUnpublishedEvent)

	engine.POST("/event/:eventId/poster", c.UploadEventPoster)
	engine.DELETE("/event/:eventId/poster", c.DeleteEventPoster)
}
