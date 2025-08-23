package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config represents the main configuration structure
type Config struct {
	Providers ProvidersConfig
	Global    GlobalConfig
}

// ProvidersConfig contains configuration for different CI providers
type ProvidersConfig struct {
	GitLab    GitLabConfig
	GitHub    GitHubConfig
	CircleCI  CircleCIConfig
	Bitbucket BitbucketConfig
}

// GlobalConfig contains global configuration options
type GlobalConfig struct {
	DefaultProvider string
	OutputFormat    string
}

// LoadEnvFile loads .env file from current directory
func LoadEnvFile() {
	envViper := viper.New()
	envViper.SetConfigName(".env")
	envViper.SetConfigType("env")
	envViper.AddConfigPath(".")

	if err := envViper.ReadInConfig(); err == nil {
		fmt.Printf("Loaded .env file\n")
		// Set environment variables from .env file with correct case
		for _, key := range envViper.AllKeys() {
			if value := envViper.GetString(key); value != "" {
				// Convert to uppercase for environment variables
				envKey := strings.ToUpper(key)
				os.Setenv(envKey, value)
			}
		}
	}
}

// Load loads configuration from environment variables and .env file
func Load() *Config {
	// Load .env file from current directory
	LoadEnvFile()

	return &Config{
		Providers: ProvidersConfig{
			GitHub:    GitHubConfig{Token: os.Getenv("GITHUB_TOKEN"), URL: os.Getenv("GITHUB_URL")},
			GitLab:    GitLabConfig{Token: os.Getenv("GITLAB_TOKEN"), URL: os.Getenv("GITLAB_URL")},
			CircleCI:  CircleCIConfig{Token: os.Getenv("CIRCLECI_TOKEN")},
			Bitbucket: BitbucketConfig{Token: os.Getenv("BITBUCKET_TOKEN"), URL: os.Getenv("BITBUCKET_URL")},
		},
		Global: GlobalConfig{
			OutputFormat: getEnvOrDefault("CI_CLI_OUTPUT_FORMAT", "table"),
		},
	}
}

// getEnvOrDefault returns environment variable value or default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
