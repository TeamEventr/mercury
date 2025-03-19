package routes

import (
	c "github.com/IAmRiteshKoushik/mercury/controllers/host"
	mw "github.com/IAmRiteshKoushik/mercury/middleware"
	"github.com/gin-gonic/gin"
)

func HostDashboardRoutes(engine *gin.RouterGroup) {
	engine.GET("/dashboard", mw.UserAuth, c.GetHostDashboard)
}
