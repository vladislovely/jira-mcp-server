package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mark3labs/mcp-go/mcp"

	"vladislove-mcp/internal/clients"
	customValidator "vladislove-mcp/internal/validator"
)

func HandleGetProjectsTool(
	ctx context.Context,
	_ mcp.CallToolRequest,
	_ GetProjectsInput,
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
		b.WriteString(fmt.Sprintf("ID: %s\n", p.ID))
		b.WriteString(fmt.Sprintf("Ключ: %s\n", p.Key))
		b.WriteString(fmt.Sprintf("Название: %s\n", p.Name))
		if p.Description != "" {
			b.WriteString(fmt.Sprintf("Описание: %s\n", p.Description))
		}

		b.WriteString(fmt.Sprintf("Тип проекта: %s\n", p.ProjectTypeKey))
		b.WriteString(
			fmt.Sprintf("Руководитель: %s (ID: %s)\n", p.Lead.DisplayName, p.Lead.AccountID),
		)
		b.WriteString(fmt.Sprintf("Ссылка: %s\n", p.Self))
		b.WriteString(fmt.Sprintf("Аватар: %s\n", p.AvatarURLs.X48))

		if p.Archived {
			b.WriteString("Статус: Архивирован\n")
			b.WriteString(fmt.Sprintf("Дата архивации: %s\n", p.ArchivedDate))
			b.WriteString(fmt.Sprintf("Кто архивировал: %s\n", p.ArchivedBy.DisplayName))
		} else {
			b.WriteString("Статус: Активен\n")
		}

		b.WriteString(fmt.Sprintf("Приватный: %t\n", p.IsPrivate))
		b.WriteString(fmt.Sprintf("Стиль: %s\n", p.Style))
		if p.EntityID != "" {
			b.WriteString(fmt.Sprintf("EntityID: %s\n", p.EntityID))
		}
		if p.UUID != "" {
			b.WriteString(fmt.Sprintf("UUID: %s\n", p.UUID))
		}
		b.WriteString("\n")
	}

	return mcp.NewToolResultText(b.String()), nil
}

func HandleCreateProjectTool(
	ctx context.Context,
	_ mcp.CallToolRequest,
	input CreateProjectInput,
) (*mcp.CallToolResult, error) {
	var validate = validator.New()

	if result := customValidator.ValidateInput(validate, input); result != nil {
		return result, nil
	}

	client := GetJiraClient()

	logger := slog.Default().With("component", "HandleCreateProjectTool")

	logger.InfoContext(ctx, "Создание проекта")

	var projectTemplateKey string
	switch input.ProjectTypeKey {
	case "business":
		projectTemplateKey = "com.pyxis.greenhopper.jira:gh-simplified-scrum"
	case "software":
		if input.ProjectType == "kanban" {
			projectTemplateKey = "com.pyxis.greenhopper.jira:gh-kanban-template"
		}

		if input.ProjectType == "scrum" {
			projectTemplateKey = "com.pyxis.greenhopper.jira:gh-scrum-template"
		}
	default:
		return mcp.NewToolResultError("❌ Неизвестный тип проекта"), nil
	}

	data := clients.CreateProjectInput{
		AccountID:          input.AccountID,
		AssigneeType:       input.AssigneeType,
		Name:               input.Name,
		Description:        input.Description,
		TaskPrefixKey:      input.TaskPrefixKey,
		ProjectTypeKey:     input.ProjectTypeKey,
		ProjectTemplateKey: projectTemplateKey,
	}

	project, err := client.CreateProject(ctx, data)
	if err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"❌ Ошибка при создании проекта.\nОшибка: %v",
				err,
			),
		), nil
	}

	var b strings.Builder
	b.WriteString("✅ Проект успешно создан")
	b.WriteString(fmt.Sprintf("Project ID - %d\n", project.ID))
	b.WriteString(fmt.Sprintf("Task Prefix Key: %s\n", project.Key))
	b.WriteString(fmt.Sprintf("Link: %s\n", project.Self))

	return mcp.NewToolResultText(b.String()), nil
}

func HandleArchiveProjectTool(
	ctx context.Context,
	_ mcp.CallToolRequest,
	input ArchiveProjectInput,
) (*mcp.CallToolResult, error) {
	var validate = validator.New()

	if result := customValidator.ValidateInput(validate, input); result != nil {
		return result, nil
	}

	client := GetJiraClient()

	logger := slog.Default().With("component", "HandleArchiveProjectTool")

	logger.InfoContext(ctx, "Архивация проекта")

	project, err := client.ArchiveProject(ctx, clients.ArchiveProjectInput(input))
	if err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"❌ Ошибка при архивации проекта.\nОшибка: %v",
				err,
			),
		), nil
	}

	if project != nil {
		return mcp.NewToolResultText("✅ Проект успешно архивирован"), nil
	}

	return mcp.NewToolResultError(
		"❌ Ошибка при архивации проекта.",
	), nil
}

func HandleDeleteProjectTool(
	ctx context.Context,
	_ mcp.CallToolRequest,
	input DeleteProjectInput,
) (*mcp.CallToolResult, error) {
	var validate = validator.New()

	if result := customValidator.ValidateInput(validate, input); result != nil {
		return result, nil
	}

	client := GetJiraClient()

	logger := slog.Default().With("component", "HandleDeleteProjectTool")

	logger.InfoContext(ctx, "Удаления проекта")

	project, err := client.DeleteProject(ctx, clients.DeleteProjectInput(input))
	if err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"❌ Ошибка при архивации проекта.\nОшибка: %v",
				err,
			),
		), nil
	}

	if project != nil {
		return mcp.NewToolResultText("✅ Проект успешно удален"), nil
	}

	return mcp.NewToolResultError(
		"❌ Ошибка при удалении проекта.",
	), nil
}

func HandleRestoreProjectTool(
	ctx context.Context,
	_ mcp.CallToolRequest,
	input RestoreProjectInput,
) (*mcp.CallToolResult, error) {
	var validate = validator.New()

	if result := customValidator.ValidateInput(validate, input); result != nil {
		return result, nil
	}

	client := GetJiraClient()

	logger := slog.Default().With("component", "HandleRestoreProjectTool")

	logger.InfoContext(ctx, "Восстановление проекта")

	project, err := client.RestoreProject(ctx, clients.RestoreProjectInput(input))
	if err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"❌ Ошибка при восстановлении проекта.\nОшибка: %v",
				err,
			),
		), nil
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("ID: %s\n", project.ID))
	b.WriteString(fmt.Sprintf("Ключ: %s\n", project.Key))
	b.WriteString(fmt.Sprintf("Название: %s\n", project.Name))
	if project.Description != "" {
		b.WriteString(fmt.Sprintf("Описание: %s\n", project.Description))
	}

	b.WriteString(fmt.Sprintf("Тип проекта: %s\n", project.ProjectTypeKey))
	b.WriteString(
		fmt.Sprintf(
			"Руководитель: %s (ID: %s)\n",
			project.Lead.DisplayName,
			project.Lead.AccountID,
		),
	)
	b.WriteString(fmt.Sprintf("Ссылка на проект: %s\n", project.Self))

	if project.Archived {
		b.WriteString("Статус: Архивирован\n")
		b.WriteString(fmt.Sprintf("Дата архивации: %s\n", project.ArchivedDate))
		b.WriteString(fmt.Sprintf("Кто архивировал: %s\n", project.ArchivedBy.DisplayName))
	} else {
		b.WriteString("Статус: Активен\n")
	}

	b.WriteString(fmt.Sprintf("Приватный: %t\n", project.IsPrivate))
	b.WriteString(fmt.Sprintf("Стиль: %s\n", project.Style))
	if project.EntityID != "" {
		b.WriteString(fmt.Sprintf("EntityID: %s\n", project.EntityID))
	}
	if project.UUID != "" {
		b.WriteString(fmt.Sprintf("UUID: %s\n\n", project.UUID))
	}

	return mcp.NewToolResultText(b.String()), nil
}

func HandleIssueFieldsTool(
	ctx context.Context,
	_ mcp.CallToolRequest,
	input IssueFieldsInput,
) (*mcp.CallToolResult, error) {
	var validate = validator.New()

	if result := customValidator.ValidateInput(validate, input); result != nil {
		return result, nil
	}

	client := GetJiraClient()

	logger := slog.Default().With("component", "HandleIssueFieldsTool")

	logger.InfoContext(ctx, "Получение список доступных полей для создания задачи")

	response, err := client.IssueFields(ctx, clients.IssueFieldsInput(input))
	if err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"❌ Ошибка при получении доступных полей.\nОшибка: %v",
				err,
			),
		), nil
	}

	p := response.Projects[0]

	var b strings.Builder
	b.WriteString(fmt.Sprintf("📌 Проект: %s (%s)\n\n", p.Name, p.Key))

	for _, it := range p.IssueTypes {
		b.WriteString(fmt.Sprintf("Тип задачи: %s\n", it.Name))
		for key, f := range it.Fields {
			b.WriteString(fmt.Sprintf("  - %s (key: %s)\n", f.Name, key))
			b.WriteString(fmt.Sprintf("    Обязательное: %t\n", f.Required))
			b.WriteString(fmt.Sprintf("    По умолчанию: %t\n", f.HasDefaultValue))

			if len(f.AllowedValues) > 0 {
				var values []string
				for _, v := range f.AllowedValues {
					if m, ok := v.(map[string]interface{}); ok {
						var item string
						if name, exists := m["name"].(string); exists {
							item = name
						}
						if id, exists := m["id"].(string); exists {
							if item != "" {
								item = fmt.Sprintf("%s (ID: %s)", item, id)
							} else {
								item = fmt.Sprintf("ID: %s", id)
							}
						}
						if item != "" {
							values = append(values, item)
						}
					}
				}
				if len(values) > 0 {
					b.WriteString(fmt.Sprintf("    Допустимые значения: %s\n", strings.Join(values, ", ")))
				}
			}

			if len(f.Operations) > 0 {
				b.WriteString(fmt.Sprintf("    Операции: %s\n", strings.Join(f.Operations, ", ")))
			}

			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	return mcp.NewToolResultText(b.String()), nil
}

func HandleIssueTypesTool(
	ctx context.Context,
	_ mcp.CallToolRequest,
	input IssueTypesInput,
) (*mcp.CallToolResult, error) {
	var validate = validator.New()

	if result := customValidator.ValidateInput(validate, input); result != nil {
		return result, nil
	}

	client := GetJiraClient()

	logger := slog.Default().With("component", "HandleIssueTypesTool")

	logger.InfoContext(ctx, "Получение список доступных типов задач")

	response, err := client.IssueTypes(ctx, clients.IssueTypesInput(input))
	if err != nil {
		return mcp.NewToolResultError(
			fmt.Sprintf(
				"❌ Ошибка при получении доступных типов задач.\nОшибка: %v",
				err,
			),
		), nil
	}

	var b strings.Builder
	for _, t := range response.IssueTypes {
		b.WriteString(fmt.Sprintf("Название: %s\n", t.Name))
		b.WriteString(fmt.Sprintf("ID: %s\n", t.ID))
		if t.Description != "" {
			b.WriteString(fmt.Sprintf("Описание: %s\n", t.Description))
		}
		b.WriteString(fmt.Sprintf("Sub-task: %t\n", t.Subtask))
		b.WriteString("\n")
	}

	return mcp.NewToolResultText(b.String()), nil
}
