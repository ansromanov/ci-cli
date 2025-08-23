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

// GlobalConfig contains global configuration options
type GlobalConfig struct {
	DefaultProvider string `yaml:"default_provider"`
	OutputFormat    string `yaml:"output_format"`
}

// getEnvWithDefault returns the value of the environment variable with the given key, or the default value if the environment variable is not set
func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// LoadFromEnv loads configuration from envirment variables
func LoadFromEnv() *Config {
	config := &Config{
		Providers: ProvidersConfig{
			GitHub:    LoadGitHubFromEnv(),
			GitLab:    LoadGitLabFromEnv(),
			CircleCI:  LoadCircleCIFromEnv(),
			Bitbucket: LoadBitbucketFromEnv(),
		},
		Global: GlobalConfig{
			OutputFormat: getEnvWithDefault("CI_CLI_OUTPUT_FORMAT", "table"),
		},
	}

	return config
}

func LoadConfig(configPath string) (*Config, error) {
	config := LoadFromEnv()

	// If config file exists, merge with environment variables
	if configPath != "" {
		fileConfig, err := Load(configPath)
		if err != nil {
			// Log warning but don't fail if file doesn't exist
			fmt.Printf("Warning: Could not load config file: %v\n", err)
		} else {
			// Merge configurations (env vars take precedence)
			config = mergeConfigs(fileConfig, config)
		}
	}

	return config, nil
}

func mergeConfigs(fileConfig, envConfig *Config) *Config {
	// Merge all provider configurations (env vars take precedence)
	envConfig.Providers.GitHub = MergeGitHubConfig(fileConfig.Providers.GitHub, envConfig.Providers.GitHub)
	envConfig.Providers.GitLab = MergeGitLabConfig(fileConfig.Providers.GitLab, envConfig.Providers.GitLab)
	envConfig.Providers.CircleCI = MergeCircleCIConfig(fileConfig.Providers.CircleCI, envConfig.Providers.CircleCI)
	envConfig.Providers.Bitbucket = MergeBitbucketConfig(fileConfig.Providers.Bitbucket, envConfig.Providers.Bitbucket)

	return envConfig
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
