package config

import (
	"fmt"
	"os"
)

const (
	// Bitbucket environment variables
	BitbucketTokenEnvVar = "BITBUCKET_TOKEN"
	BitbucketURLEnvVar   = "BITBUCKET_URL"
)

// BitbucketConfig contains Bitbucket-specific configuration
type BitbucketConfig struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

// ValidateBitbucketConfig validates Bitbucket configuration
func (c *Config) ValidateBitbucketConfig() error {
	if c.Providers.Bitbucket.Token == "" {
		return fmt.Errorf("Bitbucket token is required. Please set the %s environment variable", BitbucketTokenEnvVar)
	}
	return nil
}

// LoadBitbucketFromEnv loads Bitbucket configuration from environment variables
func LoadBitbucketFromEnv() BitbucketConfig {
	return BitbucketConfig{
		Token: os.Getenv(BitbucketTokenEnvVar),
		URL:   os.Getenv(BitbucketURLEnvVar),
	}
}

// MergeBitbucketConfig merges Bitbucket configuration with environment variables taking precedence
func MergeBitbucketConfig(fileConfig, envConfig BitbucketConfig) BitbucketConfig {
	if envConfig.Token == "" && fileConfig.Token != "" {
		envConfig.Token = fileConfig.Token
	}
	if envConfig.URL == "" && fileConfig.URL != "" {
		envConfig.URL = fileConfig.URL
	}
	return envConfig
}
