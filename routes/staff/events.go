package routes

import (
	c "github.com/IAmRiteshKoushik/mercury/controllers/staff"
	"github.com/gin-gonic/gin"
)

func StaffEventRoutes(engine *gin.RouterGroup) {
	engine.GET("/event", c.FetchStaffEvents)
	engine.GET("/event/:eventId", c.FetchStaffEventById)
}
