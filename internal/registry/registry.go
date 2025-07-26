package registry

import (
	"github.com/mark3labs/mcp-go/server"
)

func RegisterAll(mcpServer *server.MCPServer) {
	RegisterTools(mcpServer)
	RegisterResources(mcpServer)
	RegisterPrompts(mcpServer)
	RegisterValidators()
}
