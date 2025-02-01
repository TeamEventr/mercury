package routes

import (
	"github.com/IAmRiteshKoushik/mercury/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(engine *gin.RouterGroup) {

	engine.POST("/auth/user/check")
	engine.POST("/auth/user/login")
	engine.POST("/auth/user/register")
	engine.POST("/auth/user/register/otp/verify", middleware.AuthMiddleware())
	engine.POST("/auth/user/register/otp/resend", middleware.AuthMiddleware())

	engine.GET("/auth/google/")
	engine.GET("/auth/google/callback")

	engine.POST("/auth/host/register")
	engine.POST("/auth/host/register/otp/verify", middleware.AuthMiddleware())
	engine.POST("/auth/host/register/otp/resend", middleware.AuthMiddleware())
}
