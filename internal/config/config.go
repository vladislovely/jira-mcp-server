package config

import (
	"context"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JiraUser      string
	JiraToken     string
	JiraApiUrl    string
	McpOutputPort string
}

func LoadConfig() Config {
	var cfg Config
	logger := slog.Default().With("component", "config")

	err := godotenv.Load()
	if err != nil {
		logger.InfoContext(context.Background(), "Error loading .env file")
	}

	cfg.JiraToken = GetEnv("JIRA_TOKEN", "")
	cfg.JiraApiUrl = GetEnv("JIRA_API_URL", "")
	cfg.JiraUser = GetEnv("JIRA_USER", "")
	cfg.McpOutputPort = GetEnv("MCP_PORT", "")

	return cfg
}

func GetEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
