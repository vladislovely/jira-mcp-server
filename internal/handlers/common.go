package handlers

import (
	"vladislove-mcp/internal/clients"
	"vladislove-mcp/internal/config"
)

func GetJiraClient() *clients.Client {
	apiURL := config.GetEnv("JIRA_API_URL", "")
	token := config.GetEnv("JIRA_TOKEN", "")

	return clients.NewClient(apiURL, token)
}
