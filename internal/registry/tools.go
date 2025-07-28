package registry

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"vladislove-mcp/internal/handlers"
)

func RegisterTools(mcpServer *server.MCPServer) {
	const minTaskKeyLength = 2
	const maxTaskKeyLength = 10

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
				mcp.MinLength(minTaskKeyLength),
				mcp.MaxLength(maxTaskKeyLength),
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

	mcpServer.AddTool(
		mcp.NewTool(
			string(ToolArchiveProject),
			mcp.WithDescription(
				"Архивировать проект. Используй этот инструмент только для архивации реальных проектов.",
			),
			mcp.WithString(
				"project_id_or_key",
				mcp.Required(),
				mcp.Description(
					"ID или ключ проекта, который нужно архивировать. Бери только из ответа инструмента projects.list.",
				),
			),
		), mcp.NewTypedToolHandler(handlers.HandleArchiveProjectTool),
	)

	mcpServer.AddTool(
		mcp.NewTool(
			string(ToolDeleteProject),
			mcp.WithDescription(
				"Удалить проект. Проект удаляется не полностью его можно будет восстановить. Используй этот инструмент только для удаления реальных проектов.",
			),
			mcp.WithString(
				"project_id_or_key",
				mcp.Required(),
				mcp.Description(
					"ID или ключ проекта, который нужно удалить. Бери только из ответа инструмента projects.list.",
				),
			),
		), mcp.NewTypedToolHandler(handlers.HandleDeleteProjectTool),
	)

	mcpServer.AddTool(
		mcp.NewTool(
			string(ToolRestoreProject),
			mcp.WithDescription(
				"Восстановить архивированный проект. Используй этот инструмент только для восстановления реальных проектов.",
			),
			mcp.WithString(
				"project_id_or_key",
				mcp.Required(),
				mcp.Description(
					"ID или ключ проекта, который нужно восстановить. Бери только из ответа инструмента projects.list.",
				),
			),
		), mcp.NewTypedToolHandler(handlers.HandleRestoreProjectTool),
	)

	mcpServer.AddTool(
		mcp.NewTool(
			string(ToolAvailableIssueFields),
			mcp.WithDescription(
				"Получить список доступных полей для создания задачи. Используй этот инструмент только для получения реальных полей для создания задачи.",
			),
			mcp.WithString(
				"project_key",
				mcp.Required(),
				mcp.Description(
					"Ключ проекта, по которому нужно найти список доступных полей.",
				),
			),
		), mcp.NewTypedToolHandler(handlers.HandleIssueFieldsTool),
	)

	mcpServer.AddTool(
		mcp.NewTool(
			string(ToolAvailableIssueTypes),
			mcp.WithDescription(
				"Получить список доступных типов для создания задачи. Используй этот инструмент только для получения реальных типов для создания задачи.",
			),
			mcp.WithString(
				"project_id_or_key",
				mcp.Required(),
				mcp.Description(
					"ID или ключ проекта, по которому нужно получить список доступных типов задач.",
				),
			),
		), mcp.NewTypedToolHandler(handlers.HandleIssueTypesTool),
	)
}
