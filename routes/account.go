package routes

import (
	"github.com/gin-gonic/gin"
)

func AccountRoutes(engine *gin.RouterGroup) {
	// Disable account
	engine.POST("/user/disable")
	engine.POST("/user/disable/otp/verify")

	// Enable account
	engine.POST("/user/enable")
	engine.POST("/user/enable/otp/verify")

	// Permanently Delete account
	engine.DELETE("/user/delete")
	engine.POST("/user/delete/otp/verify")
	engine.POST("/user/delete/otp/resend")
}
