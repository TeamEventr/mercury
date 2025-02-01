package routes

import (
	"github.com/IAmRiteshKoushik/mercury/controllers"
	"github.com/gin-gonic/gin"
)

func DashboardRoutes(engine *gin.RouterGroup) {
	engine.GET("/dashboard", controllers.GetHostDashboard)
}
