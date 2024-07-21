package openai

import (
	"os"

	"github.com/pkg/errors"
)

const (
	apiKeyEnv = "OPENAI_API_KEY"
)

type Config struct {
	APIKey string
}

func NewConfig() (*Config, error) {
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		return nil, errors.Errorf("Environment variable %s is required", apiKeyEnv)
	}

	return &Config{
		APIKey: apiKey,
	}, nil
}