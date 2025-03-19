package routes

import (
	c "github.com/IAmRiteshKoushik/mercury/controllers/host"
	"github.com/gin-gonic/gin"
)

func HostPricingRoutes(engine *gin.RouterGroup) {
	engine.GET("/event/:eventId/price-tier/new", c.CreatePriceTierCsrf)
	engine.POST("/event/:eventId/price-tier/new", c.CreatePriceTier)
	engine.GET("/event/:eventId/price-tier", c.FetchAllPriceTier)
	engine.DELETE("/event/:eventId/staff/:staffId", c.RemovePriceTier)
}
