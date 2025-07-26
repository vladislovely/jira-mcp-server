package handlers

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleCurrentDatePrompt(
	_ context.Context,
	_ mcp.GetPromptRequest,
) (*mcp.GetPromptResult, error) {
	messages := []mcp.PromptMessage{
		{
			Role: mcp.RoleUser,
			Content: mcp.TextContent{
				Type: "text",
				Text: "Дай мне сегодняшнюю дату",
			},
		},
	}

	return mcp.NewGetPromptResult("Получить текущие дату и время.", messages), nil
}
