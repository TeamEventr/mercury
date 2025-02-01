package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/IAmRiteshKoushik/mercury/controllers"
)

func ProfileRoutes(engine *gin.RouterGroup) {
	// get user-profile (protected routes)
	engine.GET("/:username", controllers.GetUserProfile)

	// edti user-profile details (protected routes)
	engine.POST("/edit/:username", controllers.EditUserProfile)
	engine.POST("/add/profile-picture/:username", controllers.AddUserPfp)
	engine.POST("/edit/profile-picture/:username", controllers.EditUserPfp)
	engine.DELETE("/delete/profile-picture/:username", controllers.DeleteUserPfp)
}
