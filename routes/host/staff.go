package routes

import (
	c "github.com/IAmRiteshKoushik/mercury/controllers/host"
	"github.com/gin-gonic/gin"
)

func HostStaffRoutes(engine *gin.RouterGroup) {
	engine.GET("/staff", c.FetchAllStaff)
	engine.POST("/staff/new", c.CreateNewStaff)
	engine.GET("/event/:eventId/staff", c.FetchAllEventStaff)
	engine.POST("/event/:eventId/staff", c.AddEventStaff)
	engine.DELETE("/event/:eventId/staff/:staffId", c.RemoveEventStaff)
}
