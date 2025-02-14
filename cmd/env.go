package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var EnvVars EnvConfig

type EnvConfig struct {
	Environment   string `default:"development" required:"true"`
	Port          int    `default:"9000" required:"true"`
	DBUrl         string `default:"postgresql://postgres:1234@localhost:5432/" required:"true"`
	RzpKey        string
	RzpSecret     string
	OAuthClient   string
	OAuthSecret   string
	OAuthCallback string
	CacheAddr     string
	CachePwd      string
}

func (e *EnvConfig) NewEnvConfig() (*EnvConfig, error) {
	EnvVars := &EnvConfig{}
	validEnvs := []string{"development", "testing", "production"}

	environment := os.Getenv("ENVIRONMENT")
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DATABASE_URL")
	rzpKey := os.Getenv("RAZORPAY_CLIENT_ID")
	rzpSecret := os.Getenv("RAZORPAY_CLIENT_SECRET")
	oauthClientId := os.Getenv("OAUTH_CLIENT_ID")
	oauthClientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	oauthCallback := os.Getenv("OAUTH_CALLBACK_URL")
	cacheAddr := os.Getenv("REDIS_URL")
	cachePwd := os.Getenv("REDIS_PASSWORD")

	environment = strings.ToLower(environment)
	isValid := false
	for _, validEnv := range validEnvs {
		if environment == validEnv {
			isValid = true
			break
		}
	}
	if !isValid {
		return nil, fmt.Errorf("Invalid ENVIRONMENT value: %s", environment)
	}
	EnvVars.Environment = environment
	if port == "" {
		return nil, fmt.Errorf("PORT environment variable is missing.")
	}
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("Invalid PORT value: %w", err)
	}
	EnvVars.Port = portNum
	if dbUrl == "" {
		return nil, fmt.Errorf("PORT environment variable is missing.")
	}
	EnvVars.DBUrl = dbUrl
	if rzpKey == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is missing.")
	}
	EnvVars.RzpKey = rzpKey
	if rzpSecret == "" {
		return nil, fmt.Errorf("RAZORPAY_SECRET environment variable is missing.")
	}
	EnvVars.RzpSecret = rzpSecret
	if oauthClientId == "" {
		return nil, fmt.Errorf("RAZORPAY_CLIENT_ID environment variable is missing.")
	}
	EnvVars.OAuthClient = oauthClientId
	if oauthClientSecret == "" {
		return nil, fmt.Errorf("OAUTH_CLIENT_SECRET environment variable is missing.")
	}
	EnvVars.OAuthSecret = oauthClientSecret
	if oauthCallback == "" {
		return nil, fmt.Errorf("OAUTH_CALLBACK_URL environment variable is missing.")
	}
	EnvVars.OAuthCallback = oauthCallback
	if cacheAddr == "" {
		return nil, fmt.Errorf("REDIS_URL environment variable is missing.")
	}
	EnvVars.CacheAddr = cacheAddr
	if cachePwd == "" {
		return nil, fmt.Errorf("REDIS_PASSWORD environment variable is missing.")
	}
	EnvVars.CachePwd = cachePwd

	return EnvVars, nil
}
