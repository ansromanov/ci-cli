package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Common utility functions for the CI CLI

// GetCurrentProject attempts to determine the current project from the working directory
func GetCurrentProject() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}

	// Try to find git repository
	gitDir := findGitDirectory(wd)
	if gitDir != "" {
		// Extract project name from git remote or directory name
		return extractProjectFromGit(gitDir)
	}

	// Fallback to directory name
	return filepath.Base(wd), nil
}

// findGitDirectory searches for a .git directory in the current path
func findGitDirectory(path string) string {
	for {
		gitPath := filepath.Join(path, ".git")
		if _, err := os.Stat(gitPath); err == nil {
			return gitPath
		}

		parent := filepath.Dir(path)
		if parent == path {
			break // Reached root
		}
		path = parent
	}
	return ""
}

// extractProjectFromGit extracts project information from git repository
func extractProjectFromGit(gitDir string) (string, error) {
	// Read git config to get remote origin URL
	configPath := filepath.Join(gitDir, "config")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to read git config: %w", err)
	}

	// Simple parsing to find remote origin URL
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "[remote \"origin\"]" {
			// Look for URL in next few lines
			for j := i + 1; j < len(lines) && j < i+5; j++ {
				if strings.HasPrefix(strings.TrimSpace(lines[j]), "url = ") {
					url := strings.TrimSpace(strings.TrimPrefix(lines[j], "url = "))
					return extractProjectFromURL(url), nil
				}
			}
		}
	}

	// Fallback to directory name
	return filepath.Base(filepath.Dir(gitDir)), nil
}

// extractProjectFromURL extracts project name from git URL
func extractProjectFromURL(url string) string {
	// Remove .git suffix
	url = strings.TrimSuffix(url, ".git")

	// Handle different URL formats
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		parts := strings.Split(url, "/")
		if len(parts) >= 2 {
			// Extract last two parts for owner/repo format
			if len(parts) >= 3 {
				return fmt.Sprintf("%s/%s", parts[len(parts)-2], parts[len(parts)-1])
			}
			return parts[len(parts)-1]
		}
	} else if strings.HasPrefix(url, "git@") {
		// SSH format: git@github.com:owner/repo.git
		parts := strings.Split(url, ":")
		if len(parts) == 2 {
			repoParts := strings.Split(parts[1], "/")
			if len(repoParts) >= 2 {
				return fmt.Sprintf("%s/%s", repoParts[len(repoParts)-2], repoParts[len(repoParts)-1])
			}
			return repoParts[len(repoParts)-1]
		}
	}

	return url
}

// ValidateProject validates that a project string is in the correct format
func ValidateProject(project string) error {
	if project == "" {
		return fmt.Errorf("project cannot be empty")
	}

	// Basic validation - project should contain at least one character
	if len(strings.TrimSpace(project)) == 0 {
		return fmt.Errorf("project cannot be whitespace only")
	}

	return nil
}

// ParseProviderAndProject parses a string that might contain provider:project format
func ParseProviderAndProject(input string) (provider, project string) {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return "", parts[0]
}

// FormatDuration formats a duration in seconds to a human-readable string
func FormatDuration(seconds int64) string {
	if seconds < 60 {
		return fmt.Sprintf("%ds", seconds)
	} else if seconds < 3600 {
		minutes := seconds / 60
		remainingSeconds := seconds % 60
		return fmt.Sprintf("%dm %ds", minutes, remainingSeconds)
	} else {
		hours := seconds / 3600
		remainingMinutes := (seconds % 3600) / 60
		return fmt.Sprintf("%dh %dm", hours, remainingMinutes)
	}
}

// IsValidStatus checks if a status string is valid
func IsValidStatus(status string) bool {
	validStatuses := []string{
		string(StatusPending),
		string(StatusRunning),
		string(StatusSuccess),
		string(StatusFailed),
		string(StatusCancelled),
		string(StatusSkipped),
	}

	for _, valid := range validStatuses {
		if status == valid {
			return true
		}
	}
	return false
}

// NormalizeStatus normalizes a status string to a standard format
func NormalizeStatus(status string) BuildStatus {
	status = strings.ToLower(strings.TrimSpace(status))

	switch status {
	case "pending", "waiting":
		return StatusPending
	case "running", "in_progress", "in-progress":
		return StatusRunning
	case "success", "passed", "succeeded":
		return StatusSuccess
	case "failed", "failure", "error":
		return StatusFailed
	case "cancelled", "canceled":
		return StatusCancelled
	case "skipped":
		return StatusSkipped
	default:
		return BuildStatus(status)
	}
}
