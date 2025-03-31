package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"slices"

	"github.com/joho/godotenv"
)

var EnvVars *EnvConfig

type EnvConfig struct {
	Environment       string
	Port              int
	DBUrl             string
	RzpKey            string
	RzpSecret         string
	OAuthClient       string
	OAuthSecret       string
	OAuthCallback     string
	CfApiKey          string // Needed for REST API for KV (cache)
	CfApiEmail        string
	CfBucketName      string // Needed for R2 Bucket
	CfAccountId       string
	CfBucketAccessKey string
	CfBucketSecretKey string
}

func NewEnvConfig() (*EnvConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf(".env file not found")
	}

	cfg := &EnvConfig{}
	validEnvs := []string{"development", "testing", "production"}

	environment := os.Getenv("ENVIRONMENT")
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DATABASE_URL")
	rzpKey := os.Getenv("RAZORPAY_CLIENT_ID")
	rzpSecret := os.Getenv("RAZORPAY_CLIENT_SECRET")
	oauthClientId := os.Getenv("OAUTH_CLIENT_ID")
	oauthClientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	oauthCallback := os.Getenv("OAUTH_CALLBACK_URL")
	cfApiKey := os.Getenv("CF_API_KEY")
	cfApiEmail := os.Getenv("CF_API_EMAIL")
	cfAccountId := os.Getenv("CF_ACCOUNT_ID")
	cfBucketName := os.Getenv("CF_R2_BUCKET_NAME")
	cfBucketAccessKey := os.Getenv("CF_R2_BUCKET_ACCESS_KEY")
	cfBucketSecretKey := os.Getenv("CF_R2_BUCKET_SECRET_KEY")

	environment = strings.ToLower(environment)
	isValid := slices.Contains(validEnvs, environment)
	if !isValid {
		return nil, fmt.Errorf("Invalid ENVIRONMENT value: %s", environment)
	}
	cfg.Environment = environment
	if port == "" {
		return nil, fmt.Errorf("PORT environment variable is missing.")
	}
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("Invalid PORT value: %w", err)
	}
	cfg.Port = portNum
	if dbUrl == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is missing.")
	}
	cfg.DBUrl = dbUrl
	if rzpKey == "" {
		return nil, fmt.Errorf("RAZORPAY_CLIENT_ID environment variable is missing.")
	}
	cfg.RzpKey = rzpKey
	if rzpSecret == "" {
		return nil, fmt.Errorf("RAZORPAY_CLIENT_SECRET environment variable is missing.")
	}
	cfg.RzpSecret = rzpSecret
	if oauthClientId == "" {
		return nil, fmt.Errorf("OAUTH_CLIENT_ID environment variable is missing.")
	}
	cfg.OAuthClient = oauthClientId
	if oauthClientSecret == "" {
		return nil, fmt.Errorf("OAUTH_CLIENT_SECRET environment variable is missing.")
	}
	cfg.OAuthSecret = oauthClientSecret
	if oauthCallback == "" {
		return nil, fmt.Errorf("OAUTH_CALLBACK_URL environment variable is missing.")
	}
	cfg.OAuthCallback = oauthCallback
	if cfApiKey == "" {
		return nil, fmt.Errorf("CF_API_KEY variable is missing.")
	}
	cfg.CfApiKey = cfApiKey
	if cfApiEmail == "" {
		return nil, fmt.Errorf("CF_API_EMAIL environment variable is missing.")
	}
	cfg.CfApiEmail = cfApiEmail
	if cfAccountId == "" {
		return nil, fmt.Errorf("CF_ACCOUNT_ID environment variable is missing.")
	}
	cfg.CfAccountId = cfAccountId
	if cfBucketName == "" {
		return nil, fmt.Errorf("CF_R2_BUCKET_NAME environment variable is missing.")
	}
	cfg.CfBucketName = cfBucketName
	if cfBucketAccessKey == "" {
		return nil, fmt.Errorf("CF_R2_BUCKET_ACCESS_KEY environment variable is missing.")
	}
	cfg.CfBucketAccessKey = cfBucketAccessKey
	if cfBucketSecretKey == "" {
		return nil, fmt.Errorf("CF_R2_BUCKET_SECRET_KEY environment variable is missing.")
	}
	cfg.CfBucketSecretKey = cfBucketSecretKey

	return cfg, nil
}
