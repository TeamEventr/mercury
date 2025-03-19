package routes

import (
	c "github.com/IAmRiteshKoushik/mercury/controllers/host"
	mw "github.com/IAmRiteshKoushik/mercury/middleware"
	"github.com/gin-gonic/gin"
)

func HostArtistRoutes(engine *gin.RouterGroup) {
	engine.GET("/event/:eventId/artist/list", mw.UserAuth, c.FetchEventArtists)
	engine.GET("/event/:eventId/artist", mw.UserAuth, c.CreateArtistCsrf)
	engine.POST("/event/:eventId/artist", mw.UserAuth, c.CreateArtist)
	engine.DELETE("/event/:eventId/artist", mw.UserAuth, c.DeleteEventArtist)
}
