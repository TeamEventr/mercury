package routes

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(engine *gin.RouterGroup) {

	engine.POST("/check/user")
	engine.POST("/auth/user/login")
	engine.POST("/auth/user/register")
	engine.POST("/auth/user/register/otp/verify")
	engine.POST("/auth/user/register/otp/resend")

	engine.GET("/auth/google/")
	engine.GET("/auth/google/callback")

	engine.POST("/auth/host/register")
	engine.POST("/auth/host/register/otp/verify")
	engine.POST("/auth/host/register/otp/resend")
}
