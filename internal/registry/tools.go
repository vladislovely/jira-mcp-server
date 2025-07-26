package registry

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"vladislove-mcp/internal/handlers"
)

func RegisterTools(mcpServer *server.MCPServer) {
	mcpServer.AddTool(
		mcp.NewTool(
			string(ToolGetProjects),
			mcp.WithDescription(
				"Получить полный список всех проектов. Используй этот инструмент только для получения реальных проектов. Не придумывай и не выдумывай проекты.",
			),
		), mcp.NewTypedToolHandler(handlers.HandleGetProjectsTool),
	)
}
