package routes

import (
	c "github.com/IAmRiteshKoushik/mercury/controllers/customer"
	"github.com/gin-gonic/gin"
)

func CustomerEventRoutes(engine *gin.RouterGroup) {
	engine.GET("/event", c.GetEventCatalog)
	engine.GET("/event/:eventId", c.GetEventByEventId)
	engine.GET("/event/:eventId/register", c.RegisterForEventCsrf)
	engine.POST("/event/:eventId/register", c.RegisterForEvent)
}
