package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the main configuration structure
type Config struct {
	Providers ProvidersConfig `yaml:"providers"`
	Global    GlobalConfig    `yaml:"global"`
}

// ProvidersConfig contains configuration for different CI providers
type ProvidersConfig struct {
	GitLab    GitLabConfig    `yaml:"gitlab"`
	GitHub    GitHubConfig    `yaml:"github"`
	CircleCI  CircleCIConfig  `yaml:"circleci"`
	Bitbucket BitbucketConfig `yaml:"bitbucket"`
}

// GitLabConfig contains GitLab-specific configuration
type GitLabConfig struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

// GitHubConfig contains GitHub-specific configuration
type GitHubConfig struct {
	Token string `yaml:"token"`
}

// CircleCIConfig contains CircleCI-specific configuration
type CircleCIConfig struct {
	Token string `yaml:"token"`
}

// BitbucketConfig contains Bitbucket-specific configuration
type BitbucketConfig struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

// GlobalConfig contains global configuration options
type GlobalConfig struct {
	DefaultProvider string `yaml:"default_provider"`
	OutputFormat    string `yaml:"output_format"`
}

// Load loads configuration from a file
func Load(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, fmt.Errorf("config path is required")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// Save saves configuration to a file
func (c *Config) Save(configPath string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
