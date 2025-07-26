package registry

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"vladislove-mcp/internal/handlers"
)

func RegisterResources(mcpServer *server.MCPServer) {
	mcpServer.AddResource(
		mcp.NewResource(
			string(ResourceCurrentDate),
			"Сегодня",
			mcp.WithMIMEType("text/plain"),
		),
		handlers.HandleCurrentDateResource,
	)

	mcpServer.AddResource(
		mcp.NewResource(
			string(ResourcePing),
			"Проверка соединения (ping)",
			mcp.WithMIMEType("text/plain"),
		),
		handlers.HandlePingResource,
	)

	mcpServer.AddResource(
		mcp.NewResource(
			string(ResourceVersion),
			"Информация о версии сервера",
			mcp.WithMIMEType("text/plain"),
		),
		handlers.HandleVersionResource,
	)

	mcpServer.AddResource(
		mcp.NewResource(
			string(ResourceUserInfo),
			"Информация о пользователе",
			mcp.WithMIMEType("text/plain"),
		),
		handlers.HandleGetUserInfoResource,
	)
}
