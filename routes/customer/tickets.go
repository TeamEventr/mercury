package routes

import (
	c "github.com/IAmRiteshKoushik/mercury/controllers/customer"
	"github.com/gin-gonic/gin"
)

func CustomerTicketRoutes(engine *gin.RouterGroup) {
	engine.GET("/tickets", c.FetchAllTickets)
	engine.GET("/tickets/:ticketId", c.FetchTicketById)
}
