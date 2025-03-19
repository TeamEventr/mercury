package routes

import (
	"github.com/gin-gonic/gin"
)

func CustomerPaymentRoutes(engine *gin.RouterGroup) {
	engine.POST("/payment/checkout")
	engine.POST("/payment/checkout/verify")
}
