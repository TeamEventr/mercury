package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/IAmRiteshKoushik/mercury/cmd"
	"github.com/IAmRiteshKoushik/mercury/middleware"
	"github.com/IAmRiteshKoushik/mercury/routes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartApp() {
	failMsg := "Could not initialize app\n%w"

	// Initialize global environment variables
	env, err := cmd.NewEnvConfig()
	if err != nil {
		panic(fmt.Errorf(failMsg, err))
	}
	cmd.EnvVars = env

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
	router.Use(middleware.PromMiddleweare)

	// Test endpoint
	router.GET("/api/v1/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is LIVE",
		})
		c.Done()
	})

	// Metrics endpoint
	// TODO: Should not be open to the world
	router.GET("/api/v1/metrics", gin.WrapH(promhttp.Handler()))

	// TODO: Handle in next release
	// routes.AccountRoutes(v1)

	v1 := router.Group("/api/v1")
	routes.AuthRoutes(v1)
	routes.ProfileRoutes(v1)
	routes.EventRoutes(v1)
	routes.PaymentRoutes(v1)
	routes.DashboardRoutes(v1)

	return router
}

func main() {
	StartApp()
}
