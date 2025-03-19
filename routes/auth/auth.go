package routes

import (
	c "github.com/IAmRiteshKoushik/mercury/controllers/auth"
	mw "github.com/IAmRiteshKoushik/mercury/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(engine *gin.RouterGroup) {

	engine.POST("/user/check", c.CheckUsername)

	engine.POST("/user/login", c.LoginUser)
	engine.POST("/user/register", c.RegisterUserAccount)
	engine.POST("/user/register/otp/verify", mw.UserAuth, c.VerifyUserOtp)
	engine.POST("/user/register/otp/resend", mw.UserAuth, c.ResendUserOtp)

	engine.GET("/google/", c.OAuthGoogleLogin)
	engine.GET("/google/callback", c.OAuthGoogleCallback)

	engine.POST("/host/register", c.RegisterHostAccount)
	engine.POST("/host/register/otp/verify", mw.UserAuth, c.VerifyHostOtp)
	engine.POST("/host/register/otp/resend", mw.UserAuth, c.ResendHostOtp)
}
