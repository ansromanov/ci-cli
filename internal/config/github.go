package config

import (
	"fmt"
	"os"
)

const (
	// GitHub environment variables
	GitHubTokenEnvVar = "GITHUB_TOKEN"
	GitHubURLEnvVar   = "GITHUB_URL"
)

// GitHubConfig contains GitHub-specific configuration
type GitHubConfig struct {
	Token string `yaml:"token"`
	URL   string `yaml:"url"`
}

// ValidateGitHubConfig validates GitHub configuration
func (c *Config) ValidateGitHubConfig() error {
	if c.Providers.GitHub.Token == "" {
		return fmt.Errorf("GitHub token is required. Please set the %s environment variable", GitHubTokenEnvVar)
	}
	return nil
}

// LoadGitHubFromEnv loads GitHub configuration from environment variables
func LoadGitHubFromEnv() GitHubConfig {
	return GitHubConfig{
		Token: os.Getenv(GitHubTokenEnvVar),
		URL:   os.Getenv(GitHubURLEnvVar),
	}
}

// MergeGitHubConfig merges GitHub configuration with environment variables taking precedence
func MergeGitHubConfig(fileConfig, envConfig GitHubConfig) GitHubConfig {
	if envConfig.Token == "" && fileConfig.Token != "" {
		envConfig.Token = fileConfig.Token
	}
	if envConfig.URL == "" && fileConfig.URL != "" {
		envConfig.URL = fileConfig.URL
	}
	return envConfig
}
