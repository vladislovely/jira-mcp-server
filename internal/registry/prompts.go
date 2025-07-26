package registry

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"vladislove-mcp/internal/handlers"
)

func RegisterPrompts(mcpServer *server.MCPServer) {
	mcpServer.AddPrompt(
		mcp.NewPrompt(
			string(PromptGetCurrentDate),
			mcp.WithPromptDescription("Получить текущие дату и время."),
		), handlers.HandleCurrentDatePrompt,
	)
}
