package core

import (
	"context"
)

// CIProvider defines the interface that all CI providers must implement
type CIProvider interface {
	// GetBuilds retrieves builds for a project
	GetBuilds(ctx context.Context, project string, options BuildOptions) ([]Build, error)

	// GetBuild retrieves a specific build by ID
	GetBuild(ctx context.Context, project string, buildID string) (*Build, error)

	// GetPipelines retrieves pipelines for a project
	GetPipelines(ctx context.Context, project string, options PipelineOptions) ([]Pipeline, error)

	// GetPipeline retrieves a specific pipeline by ID
	GetPipeline(ctx context.Context, project string, pipelineID string) (*Pipeline, error)

	// TriggerBuild triggers a new build
	TriggerBuild(ctx context.Context, project string, request TriggerRequest) (*TriggerResponse, error)

	// CancelBuild cancels a running build
	CancelBuild(ctx context.Context, project string, buildID string) error

	// GetJobLogs retrieves logs for a specific job
	GetJobLogs(ctx context.Context, project string, buildID string, jobID string) (string, error)

	// GetProviderName returns the name of the provider
	GetProviderName() string
}

// BuildOptions contains options for retrieving builds
type BuildOptions struct {
	Branch string `json:"branch"`
	Status string `json:"status"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Since  string `json:"since"`
	Until  string `json:"until"`
}

// PipelineOptions contains options for retrieving pipelines
type PipelineOptions struct {
	Branch string `json:"branch"`
	Status string `json:"status"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Since  string `json:"since"`
	Until  string `json:"until"`
}

// ProviderFactory creates CI provider instances
type ProviderFactory interface {
	CreateProvider(providerType string, config map[string]interface{}) (CIProvider, error)
}

// DefaultProviderFactory implements ProviderFactory
type DefaultProviderFactory struct{}

// CreateProvider creates a new CI provider instance
func (f *DefaultProviderFactory) CreateProvider(providerType string, config map[string]interface{}) (CIProvider, error) {
	switch providerType {
	case "gitlab":
		return NewGitLabProvider(config)
	case "github":
		return NewGitHubProvider(config)
	case "circleci":
		return NewCircleCIProvider(config)
	default:
		return nil, &UnsupportedProviderError{Provider: providerType}
	}
}

// UnsupportedProviderError is returned when a provider is not supported
type UnsupportedProviderError struct {
	Provider string
}

func (e *UnsupportedProviderError) Error() string {
	return "unsupported provider: " + e.Provider
}

// NewGitLabProvider creates a new GitLab provider instance
func NewGitLabProvider(config map[string]interface{}) (CIProvider, error) {
	// TODO: Implement GitLab provider
	return nil, &UnsupportedProviderError{Provider: "gitlab"}
}

// NewGitHubProvider creates a new GitHub provider instance
func NewGitHubProvider(config map[string]interface{}) (CIProvider, error) {
	// TODO: Implement GitHub provider
	return nil, &UnsupportedProviderError{Provider: "github"}
}

// NewCircleCIProvider creates a new CircleCI provider instance
func NewCircleCIProvider(config map[string]interface{}) (CIProvider, error) {
	// TODO: Implement CircleCI provider
	return nil, &UnsupportedProviderError{Provider: "circleci"}
}
