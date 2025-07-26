package handlers

import (
	"encoding/base64"

	"vladislove-mcp/internal/clients"
	"vladislove-mcp/internal/config"
)

func GetJiraClient() *clients.Client {
	apiURL := config.GetEnv("JIRA_API_URL", "")
	user := config.GetEnv("JIRA_USER", "")
	token := config.GetEnv("JIRA_TOKEN", "")

	auth := user + ":" + token
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	return clients.NewClient(apiURL, encodedAuth)
}
