package cmd

import (
	"log"
	"net/http"

	"github.com/IAmRiteshKoushik/mercury/routes"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.New()
	router.Use(gin.Logger())

	// Test route
	router.GET("/api/v1/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is LIVE",
		})
		c.Done()
	})

	v1 := router.Group("/api/v1")

	routes.AuthRoutes(v1)

	// TODO: Handle in next release
	// routes.AccountRoutes(v1)

	routes.ProfileRoutes(v1)
	routes.EventRoutes(v1)
	routes.PaymentRoutes(v1)
	routes.DashboardRoutes(v1)

	err := router.Run(":9000")
	if err != nil {
		log.Fatal("[ERROR]: Could not start server")
	}
}
