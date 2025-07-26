package handlers

import (
	"context"
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
	_ context.Context,
	_ mcp.ReadResourceRequest,
) ([]mcp.ResourceContents, error) {
	var b strings.Builder
	b.WriteString("Имя - Владислав Зворыгин\n")
	b.WriteString("Логин - vladislove\n")
	b.WriteString("Email - vladislav.zvorygin147@gmail.com\n")
	b.WriteString("Role - Software Engineer\n")

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      string(ResourceUserInfo),
			MIMEType: "text/plain",
			Text:     b.String(),
		},
	}, nil
}
