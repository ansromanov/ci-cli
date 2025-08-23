package config

import (
	"fmt"
	"os"
)

const (
	// CircleCI environment variables
	CircleCITokenEnvVar = "CIRCLECI_TOKEN"
	CircleCIURLEnvVar   = "CIRCLECI_URL"
)

// CircleCIConfig contains CircleCI-specific configuration
type CircleCIConfig struct {
	Token string `yaml:"token"`
	URL   string `yaml:"url"`
}

// ValidateCircleCIConfig validates CircleCI configuration
func (c *Config) ValidateCircleCIConfig() error {
	if c.Providers.CircleCI.Token == "" {
		return fmt.Errorf("CircleCI token is required. Please set the %s environment variable", CircleCITokenEnvVar)
	}
	return nil
}

// LoadCircleCIFromEnv loads CircleCI configuration from environment variables
func LoadCircleCIFromEnv() CircleCIConfig {
	return CircleCIConfig{
		Token: os.Getenv(CircleCITokenEnvVar),
		URL:   os.Getenv(CircleCIURLEnvVar),
	}
}

// MergeCircleCIConfig merges CircleCI configuration with environment variables taking precedence
func MergeCircleCIConfig(fileConfig, envConfig CircleCIConfig) CircleCIConfig {
	if envConfig.Token == "" && fileConfig.Token != "" {
		envConfig.Token = fileConfig.Token
	}
	if envConfig.URL == "" && fileConfig.URL != "" {
		envConfig.URL = fileConfig.URL
	}
	return envConfig
}
