package routes

import (
	"github.com/gin-gonic/gin"
)

func DashboardRoutes(engine *gin.RouterGroup) {
	engine.GET("/dashboard")
}
