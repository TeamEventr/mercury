package routes

import (
	c "github.com/IAmRiteshKoushik/mercury/controllers/customer"
	"github.com/gin-gonic/gin"
)

func CustomerProfileRoutes(engine *gin.RouterGroup) {
	engine.GET("/:username", c.GetMyProfile)
	engine.POST("/edit/:username", c.EditMyProfile)
	engine.POST("/add/profile-picture/:username", c.AddPfp)
	engine.DELETE("/delete/profile-picture/:username", c.DeletePfp)
}
