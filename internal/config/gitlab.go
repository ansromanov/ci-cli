package config

import (
	"fmt"
	"os"
)

const (
	// GitLab environment variables
	GitLabTokenEnvVar = "GITLAB_TOKEN"
	GitLabURLEnvVar   = "GITLAB_URL"
)

// GitLabConfig contains GitLab-specific configuration
type GitLabConfig struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

// ValidateGitLabConfig validates GitLab configuration
func (c *Config) ValidateGitLabConfig() error {
	if c.Providers.GitLab.Token == "" {
		return fmt.Errorf("GitLab token is required. Please set the %s environment variable", GitLabTokenEnvVar)
	}
	return nil
}

// LoadGitLabFromEnv loads GitLab configuration from environment variables
func LoadGitLabFromEnv() GitLabConfig {
	return GitLabConfig{
		Token: os.Getenv(GitLabTokenEnvVar),
		URL:   os.Getenv(GitLabURLEnvVar),
	}
}

// MergeGitLabConfig merges GitLab configuration with environment variables taking precedence
func MergeGitLabConfig(fileConfig, envConfig GitLabConfig) GitLabConfig {
	if envConfig.Token == "" && fileConfig.Token != "" {
		envConfig.Token = fileConfig.Token
	}
	if envConfig.URL == "" && fileConfig.URL != "" {
		envConfig.URL = fileConfig.URL
	}
	return envConfig
}
