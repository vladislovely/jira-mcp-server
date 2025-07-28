package registry

type ToolName string
type PromptName string
type ResourceName string

const (
	ToolGetProjects    ToolName = "project.list"
	ToolCreateProject  ToolName = "project.create"
	ToolArchiveProject ToolName = "project.archive"
	ToolDeleteProject  ToolName = "project.delete"
	ToolRestoreProject ToolName = "project.restore"

	ToolAvailableIssueFields ToolName = "issue.available-fields"
	ToolAvailableIssueTypes  ToolName = "issue.available-types"
)

const (
	PromptGetCurrentDate PromptName = "get_current_date"
)

const (
	ResourceCurrentDate ResourceName = "resource://current-date"
	ResourceVersion     ResourceName = "resource://version"
	ResourcePing        ResourceName = "resource://ping"
	ResourceUserInfo    ResourceName = "resource://user-info"
)
