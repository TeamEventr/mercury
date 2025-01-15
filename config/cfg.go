package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EnvConfig struct {
	Environment string `default:"development" required:"true"`
	Port        int    `default:"9000" required:"true"`
	// Host        string `default:"localhost" required:"true"`
	Db string `default:"postgresql://postgres:1234@localhost:5432/" required:"true"`
}

func (e *EnvConfig) NewEnv() (*EnvConfig, error) {
	cfg := &EnvConfig{}
	validEnvs := []string{"development", "testing", "production"}

	envStr := os.Getenv("ENVIRONMENT")
	portStr := os.Getenv("PORT")
	database := os.Getenv("DATABASE_URL")

	// Validating and setting ENVIRONMENT
	if envStr == "" {
		return nil, fmt.Errorf("ENVIRONMENT environment variable is missing.")
	}
	envStr = strings.ToLower(envStr)
	isValid := false
	for _, validEnv := range validEnvs {
		if envStr == validEnv {
			isValid = true
			break
		}
	}
	if !isValid {
		return nil, fmt.Errorf("Invalid ENVIRONMENT value: %s", envStr)
	}
	cfg.Environment = envStr

	// Validating and setting PORT
	if portStr == "" {
		return nil, fmt.Errorf("PORT environment variable is missing.")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("Invalid PORT value: %w", err)
	}
	cfg.Port = port

	// Validating and setting DATABASE_URL
	if database == "" {
		return nil, fmt.Errorf("DATABASE environment variable is missing.")
	}
	cfg.Db = database

	return cfg, nil
}
