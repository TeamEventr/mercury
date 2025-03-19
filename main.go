package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/IAmRiteshKoushik/mercury/cmd"
	"github.com/IAmRiteshKoushik/mercury/helpers"
	mw "github.com/IAmRiteshKoushik/mercury/middleware"
	rAuth "github.com/IAmRiteshKoushik/mercury/routes/auth"
	rUser "github.com/IAmRiteshKoushik/mercury/routes/customer"
	rHost "github.com/IAmRiteshKoushik/mercury/routes/host"
	rStaff "github.com/IAmRiteshKoushik/mercury/routes/staff"
)

func StartApp() {
	failMsg := "Could not initialize app\n%w"

	// Initialize global environment variables
	env, err := cmd.NewEnvConfig()
	if err != nil {
		panic(fmt.Errorf(failMsg, err))
	}
	cmd.EnvVars = env

	// Setup RSA + Initialize PASETO
	err = helpers.GenerateRSAKeyPair(2048)
	if err != nil {
		panic(fmt.Errorf(failMsg, err))
	}
	err = helpers.InitPaseto()
	if err != nil {
		panic(fmt.Errorf(failMsg, err))
	}

	// Initialize database connection pool
	cmd.DBPool, err = cmd.InitDB()
	if err != nil {
		panic(fmt.Errorf(failMsg, err))
	}

	// Initialize redis cache
	ctx := context.Background()
	cfg := cmd.RedisConfig{
		Addr:     cmd.EnvVars.CacheAddr,
		Password: cmd.EnvVars.CachePwd,
	}
	cmd.Cache, err = cmd.NewRedisClient(ctx, cfg)
	if err != nil {
		panic(fmt.Errorf(failMsg, err))
	}

	// Initialize s3 (object store)

	// Initialize SQS (mailer service)

	// Initialize razorpay
	cmd.Pay, err = cmd.NewRzpConfig(cmd.EnvVars.RzpKey, cmd.EnvVars.RzpSecret)
	if err != nil {
		panic(fmt.Errorf(failMsg, err))
	}

	// Initialize logger
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf(failMsg, err))
	}
	defer f.Close()
	cmd.Log = cmd.NewLoggerService(cmd.EnvVars.Environment, f)

	// Initialize Server
	server := InitServer()
	err = server.Run(":" + string(cmd.EnvVars.Port))
	if err != nil {
		panic(fmt.Errorf(failMsg, err))
	}
}

func InitServer() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(mw.PromMiddleweare)

	// Test endpoint
	router.GET("/api/v1/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is LIVE",
		})
		c.Done()
	})

	// Metrics endpoint
	// Setup prometheus metrics for this endpoint

	// TODO: Handle in next release
	// routes.AccountRoutes(v1)

	v1 := router.Group("/api/v1")

	authRouter := v1.Group("/auth")
	customerRouter := v1.Group("/user")
	hostRouter := v1.Group("/host")
	staffRouter := v1.Group("/staff")

	rAuth.AuthRoutes(authRouter)

	rUser.CustomerEventRoutes(customerRouter)
	rUser.CustomerProfileRoutes(customerRouter)
	rUser.CustomerBookmarkRoutes(customerRouter)
	rUser.CustomerPaymentRoutes(customerRouter)
	rUser.CustomerTicketRoutes(customerRouter)

	rHost.HostDashboardRoutes(hostRouter)
	rHost.HostEventRoutes(hostRouter)
	rHost.HostArtistRoutes(hostRouter)
	rHost.HostPricingRoutes(hostRouter)
	rHost.HostSearchRoutes(hostRouter)
	rHost.HostStaffRoutes(hostRouter)

	rStaff.StaffEventRoutes(staffRouter)
	rStaff.StaffTicketRoutes(staffRouter)

	return router
}

func main() {
	StartApp()
}
