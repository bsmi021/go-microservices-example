package config

import (
	"os"
)

const (
	apiGitHubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	// LogLevel is the type of logging for the application
	LogLevel = "info"
	goEnvironment = "GO_ENVIRONMENT"
	production = "production"
)

var (
	githubAccessToken = os.Getenv(apiGitHubAccessToken)
)

// GetGitHubAccessToken returns the access token for github from the system environment variables
func GetGitHubAccessToken () string {
	return githubAccessToken
}

func isProduction() bool {
	return os.Getenv(goEnvironment) == production
}