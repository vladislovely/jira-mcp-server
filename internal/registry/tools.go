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

	mcpServer.AddTool(
		mcp.NewTool(
			string(ToolCreateProject),
			mcp.WithDescription(
				"Создать новый проект. Используй этот инструмент только для создания реальных проектов.",
			),
			mcp.WithString(
				"account_id",
				mcp.Required(),
				mcp.Description("ID пользователя-руководителя проекта. Бери из get-user-info."),
				mcp.DefaultString("617644e8b9c549006fc08a11"),
			),
			mcp.WithString(
				"assignee_type",
				mcp.Required(),
				mcp.Enum("PROJECT_LEAD", "UNASSIGNED"),
				mcp.Description(
					"Кто будет назначаться исполнителем задач по умолчанию. Один из списка - PROJECT_LEAD, UNASSIGNED",
				),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Название проекта. Используй реальный ввод пользователя."),
			),
			mcp.WithString(
				"description",
				mcp.Required(),
				mcp.Description(
					"Краткое описание проекта. Бери из реального пользовательского ввода.",
				),
			),
			mcp.WithString(
				"task_prefix_key",
				mcp.Required(),
				mcp.Description(
					"Ключ проекта (2–10 латинских букв). Будет префиксом задач, например EX-1.",
				),
			),
			mcp.WithString(
				"project_type",
				mcp.Enum("kanban", "scrum"),
				mcp.Description(
					"Методология проекта: kanban, scrum. Обязательный если тип проекта software.",
				),
			),
			mcp.WithString(
				"project_type_key",
				mcp.Required(),
				mcp.Enum("business", "service_desk", "software"),
				mcp.Description("Тип проекта: business, service_desk или software."),
			),
		), mcp.NewTypedToolHandler(handlers.HandleCreateProjectTool),
	)
}
