package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleCurrentDateResource(
	_ context.Context,
	_ mcp.ReadResourceRequest,
) ([]mcp.ResourceContents, error) {
	now := time.Now().Format(time.DateTime)
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      string(ResourceCurrentDate),
			MIMEType: "text/plain",
			Text:     "📅 Current date/time: " + now,
		},
	}, nil
}

func HandlePingResource(
	_ context.Context,
	_ mcp.ReadResourceRequest,
) ([]mcp.ResourceContents, error) {
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      string(ResourcePing),
			MIMEType: "text/plain",
			Text:     "✅ Pong",
		},
	}, nil
}

func HandleVersionResource(
	_ context.Context,
	_ mcp.ReadResourceRequest,
) ([]mcp.ResourceContents, error) {
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      string(ResourceVersion),
			MIMEType: "text/plain",
			Text:     "🛠 MCP Server v1.0.0",
		},
	}, nil
}

func HandleGetUserInfoResource(
	ctx context.Context,
	_ mcp.ReadResourceRequest,
) ([]mcp.ResourceContents, error) {
	client := GetJiraClient()

	userInfo, err := client.Me(ctx)

	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("AccountID: %s\n", userInfo.AccountID))
	b.WriteString(fmt.Sprintf("Пользователь: %s\n", userInfo.DisplayName))
	b.WriteString(fmt.Sprintf("Email: %s\n", userInfo.EmailAddress))
	b.WriteString(fmt.Sprintf("Timezone: %s\n", userInfo.TimeZone))
	b.WriteString(fmt.Sprintf("Locale: %s\n", userInfo.Locale))

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      string(ResourceUserInfo),
			MIMEType: "text/plain",
			Text:     b.String(),
		},
	}, nil
}
