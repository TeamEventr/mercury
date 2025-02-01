package routes

import (
	"github.com/IAmRiteshKoushik/mercury/controllers"
	"github.com/IAmRiteshKoushik/mercury/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(engine *gin.RouterGroup) {

	engine.POST("/auth/user/check", controllers.CheckUsername)
	engine.POST("/auth/user/login", controllers.LoginUser)
	engine.POST("/auth/user/register", controllers.RegisterUserAccount)
	engine.POST("/auth/user/register/otp/verify", middleware.AuthMiddleware(), controllers.VerifyUserOtp)
	engine.POST("/auth/user/register/otp/resend", middleware.AuthMiddleware(), controllers.ResendUserOtp)

	engine.GET("/auth/google/")
	engine.GET("/auth/google/callback")

	engine.POST("/auth/host/register", controllers.RegisterHostAccount)
	engine.POST("/auth/host/register/otp/verify", middleware.AuthMiddleware(), controllers.VerifyHostOtp)
	engine.POST("/auth/host/register/otp/resend", middleware.AuthMiddleware(), controllers.ResendHostOtp)
}
