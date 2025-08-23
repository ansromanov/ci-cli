package core

import (
	"time"
)

// BuildStatus represents the status of a build
type BuildStatus string

const (
	StatusPending   BuildStatus = "pending"
	StatusRunning   BuildStatus = "running"
	StatusSuccess   BuildStatus = "success"
	StatusFailed    BuildStatus = "failed"
	StatusCancelled BuildStatus = "cancelled"
	StatusSkipped   BuildStatus = "skipped"
)

// Build represents a CI build across different providers
type Build struct {
	ID          string      `json:"id"`
	Number      int         `json:"number"`
	Status      BuildStatus `json:"status"`
	Branch      string      `json:"branch"`
	Commit      string      `json:"commit"`
	CommitMsg   string      `json:"commit_message"`
	Author      string      `json:"author"`
	StartedAt   *time.Time  `json:"started_at"`
	FinishedAt  *time.Time  `json:"finished_at"`
	Duration    int64       `json:"duration_seconds"`
	Provider    string      `json:"provider"`
	Project     string      `json:"project"`
	Repository  string      `json:"repository"`
	TriggeredBy string      `json:"triggered_by"`
	Jobs        []Job       `json:"jobs"`
}

// Job represents a job within a build
type Job struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Status     BuildStatus `json:"status"`
	StartedAt  *time.Time  `json:"started_at"`
	FinishedAt *time.Time  `json:"finished_at"`
	Duration   int64       `json:"duration_seconds"`
	Stage      string      `json:"stage"`
	Runner     string      `json:"runner"`
	Logs       string      `json:"logs"`
}

// Pipeline represents a CI pipeline
type Pipeline struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Status     BuildStatus `json:"status"`
	Branch     string      `json:"branch"`
	Commit     string      `json:"commit"`
	CreatedAt  *time.Time  `json:"created_at"`
	UpdatedAt  *time.Time  `json:"updated_at"`
	Provider   string      `json:"provider"`
	Project    string      `json:"project"`
	Repository string      `json:"repository"`
	Builds     []Build     `json:"builds"`
}

// TriggerRequest represents a request to trigger a new build
type TriggerRequest struct {
	Branch     string            `json:"branch"`
	Commit     string            `json:"commit"`
	Variables  map[string]string `json:"variables"`
	Parameters map[string]string `json:"parameters"`
}

// TriggerResponse represents the response from triggering a build
type TriggerResponse struct {
	BuildID    string `json:"build_id"`
	PipelineID string `json:"pipeline_id"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}
