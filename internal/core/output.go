package core

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// OutputFormat represents the format for output
type OutputFormat string

const (
	FormatText OutputFormat = "text"
	FormatJSON OutputFormat = "json"
	FormatYAML OutputFormat = "yaml"
)

// Formatter formats CI data for output
type Formatter struct {
	format OutputFormat
}

// NewFormatter creates a new formatter with the specified format
func NewFormatter(format OutputFormat) *Formatter {
	return &Formatter{format: format}
}

// FormatBuilds formats a list of builds for output
func (f *Formatter) FormatBuilds(builds []Build) (string, error) {
	switch f.format {
	case FormatJSON:
		return f.formatBuildsJSON(builds)
	case FormatYAML:
		return f.formatBuildsYAML(builds)
	case FormatText:
		return f.formatBuildsText(builds), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", f.format)
	}
}

// FormatBuild formats a single build for output
func (f *Formatter) FormatBuild(build *Build) (string, error) {
	switch f.format {
	case FormatJSON:
		return f.formatBuildJSON(build)
	case FormatYAML:
		return f.formatBuildYAML(build)
	case FormatText:
		return f.formatBuildText(build), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", f.format)
	}
}

// FormatPipelines formats a list of pipelines for output
func (f *Formatter) FormatPipelines(pipelines []Pipeline) (string, error) {
	switch f.format {
	case FormatJSON:
		return f.formatPipelinesJSON(pipelines)
	case FormatYAML:
		return f.formatPipelinesYAML(pipelines)
	case FormatText:
		return f.formatPipelinesText(pipelines), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", f.format)
	}
}

// formatBuildsText formats builds as text
func (f *Formatter) formatBuildsText(builds []Build) string {
	if len(builds) == 0 {
		return "No builds found"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Found %d builds:\n\n", len(builds)))

	for _, build := range builds {
		sb.WriteString(f.formatBuildText(&build))
		sb.WriteString("\n")
	}

	return sb.String()
}

// formatBuildText formats a single build as text
func (f *Formatter) formatBuildText(build *Build) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Build #%d (%s)\n", build.Number, build.ID))
	sb.WriteString(fmt.Sprintf("  Status: %s\n", build.Status))
	sb.WriteString(fmt.Sprintf("  Branch: %s\n", build.Branch))
	sb.WriteString(fmt.Sprintf("  Commit: %s\n", build.Commit))
	sb.WriteString(fmt.Sprintf("  Author: %s\n", build.Author))
	sb.WriteString(fmt.Sprintf("  Provider: %s\n", build.Provider))
	sb.WriteString(fmt.Sprintf("  Project: %s\n", build.Project))

	if build.StartedAt != nil {
		sb.WriteString(fmt.Sprintf("  Started: %s\n", build.StartedAt.Format(time.RFC3339)))
	}

	if build.FinishedAt != nil {
		sb.WriteString(fmt.Sprintf("  Finished: %s\n", build.FinishedAt.Format(time.RFC3339)))
	}

	if build.Duration > 0 {
		sb.WriteString(fmt.Sprintf("  Duration: %ds\n", build.Duration))
	}

	if len(build.Jobs) > 0 {
		sb.WriteString("  Jobs:\n")
		for _, job := range build.Jobs {
			sb.WriteString(fmt.Sprintf("    - %s: %s\n", job.Name, job.Status))
		}
	}

	return sb.String()
}

// formatPipelinesText formats pipelines as text
func (f *Formatter) formatPipelinesText(pipelines []Pipeline) string {
	if len(pipelines) == 0 {
		return "No pipelines found"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Found %d pipelines:\n\n", len(pipelines)))

	for _, pipeline := range pipelines {
		sb.WriteString(f.formatPipelineText(&pipeline))
		sb.WriteString("\n")
	}

	return sb.String()
}

// formatPipelineText formats a single pipeline as text
func (f *Formatter) formatPipelineText(pipeline *Pipeline) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Pipeline %s (%s)\n", pipeline.Name, pipeline.ID))
	sb.WriteString(fmt.Sprintf("  Status: %s\n", pipeline.Status))
	sb.WriteString(fmt.Sprintf("  Branch: %s\n", pipeline.Branch))
	sb.WriteString(fmt.Sprintf("  Commit: %s\n", pipeline.Commit))
	sb.WriteString(fmt.Sprintf("  Provider: %s\n", pipeline.Provider))
	sb.WriteString(fmt.Sprintf("  Project: %s\n", pipeline.Project))

	if pipeline.CreatedAt != nil {
		sb.WriteString(fmt.Sprintf("  Created: %s\n", pipeline.CreatedAt.Format(time.RFC3339)))
	}

	if pipeline.UpdatedAt != nil {
		sb.WriteString(fmt.Sprintf("  Updated: %s\n", pipeline.UpdatedAt.Format(time.RFC3339)))
	}

	if len(pipeline.Builds) > 0 {
		sb.WriteString(fmt.Sprintf("  Builds: %d\n", len(pipeline.Builds)))
	}

	return sb.String()
}

// formatBuildsJSON formats builds as JSON
func (f *Formatter) formatBuildsJSON(builds []Build) (string, error) {
	data, err := json.MarshalIndent(builds, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// formatBuildJSON formats a single build as JSON
func (f *Formatter) formatBuildJSON(build *Build) (string, error) {
	data, err := json.MarshalIndent(build, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// formatPipelinesJSON formats pipelines as JSON
func (f *Formatter) formatPipelinesJSON(pipelines []Pipeline) (string, error) {
	data, err := json.MarshalIndent(pipelines, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// formatBuildsYAML formats builds as YAML
func (f *Formatter) formatBuildsYAML(builds []Build) (string, error) {
	// TODO: Implement YAML formatting
	return "", fmt.Errorf("YAML formatting not implemented yet")
}

// formatBuildYAML formats a single build as YAML
func (f *Formatter) formatBuildYAML(build *Build) (string, error) {
	// TODO: Implement YAML formatting
	return "", fmt.Errorf("YAML formatting not implemented yet")
}

// formatPipelinesYAML formats pipelines as YAML
func (f *Formatter) formatPipelinesYAML(pipelines []Pipeline) (string, error) {
	// TODO: Implement YAML formatting
	return "", fmt.Errorf("YAML formatting not implemented yet")
}
