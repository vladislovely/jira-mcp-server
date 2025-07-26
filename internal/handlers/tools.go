package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleGetProjectsTool(
	ctx context.Context,
	_ mcp.CallToolRequest,
	input GetProjectsInput,
) (*mcp.CallToolResult, error) {
	client := GetJiraClient()

	logger := slog.Default().With("component", "HandleGetProjectsTool")

	logger.InfoContext(ctx, "📋 Получение списка проектов")

	projects, err := client.GetProjects(ctx)
	if err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"❌ Ошибка при получении списка проектов.\nОшибка: %v",
				err,
			),
		), nil
	}

	total := len(projects.Data)

	var b strings.Builder
	b.WriteString(
		fmt.Sprintf("✅ Найдено проектов: %d\n", total),
	)
	b.WriteString("📋 Список проектов:\n\n")

	for i, p := range projects.Data {
		b.WriteString(fmt.Sprintf("%d. %s\n", i+1, p.Name))
		b.WriteString(fmt.Sprintf("ID: %s\n", p.Id))
		b.WriteString(fmt.Sprintf("Key: %s\n", p.Key))
		b.WriteString(fmt.Sprintf("EntityID: %s\n", p.EntityId))
		b.WriteString(fmt.Sprintf("Name: %s\n\n", p.Name))
	}

	return mcp.NewToolResultText(b.String()), nil
}
