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
		b.WriteString(fmt.Sprintf("Key: %s\n", p.Key))
		b.WriteString(fmt.Sprintf("EntityID: %s\n", p.EntityID))
		b.WriteString(fmt.Sprintf("Name: %s\n", p.Name))
		b.WriteString(fmt.Sprintf("Project Type: %s\n\n", p.ProjectTypeKey))
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
				"❌ Ошибка при создании проектоа.\nОшибка: %v",
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
