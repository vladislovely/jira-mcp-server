package handlers

type ResourceName string

const (
	ResourceCurrentDate ResourceName = "resource://current-date"
	ResourceVersion     ResourceName = "resource://version"
	ResourcePing        ResourceName = "resource://ping"
	ResourceUserInfo    ResourceName = "resource://user-info"
)

type GetProjectsInput struct{}

type CreateProjectInput struct {
	AccountID      string `json:"account_id"       validate:"required"`
	AssigneeType   string `json:"assignee_type"    validate:"required,oneof=PROJECT_LEAD UNASSIGNED"`
	Name           string `json:"name"             validate:"required"`
	Description    string `json:"description"      validate:"required"`
	TaskPrefixKey  string `json:"task_prefix_key"  validate:"required"`
	ProjectType    string `json:"project_type"     validate:"required_if=ProjectTypeKey software,omitempty,oneof=kanban scrum"`
	ProjectTypeKey string `json:"project_type_key" validate:"required,oneof=business software"`
}

type ArchiveProjectInput struct {
	ProjectIDOrKey string `json:"project_id_or_key" validate:"required"`
}

type DeleteProjectInput struct {
	ProjectIDOrKey string `json:"project_id_or_key" validate:"required"`
}

type RestoreProjectInput struct {
	ProjectIDOrKey string `json:"project_id_or_key" validate:"required"`
}
