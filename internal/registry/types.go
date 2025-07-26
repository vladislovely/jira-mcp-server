package registry

type PromptName string
type ResourceName string

const (
	PromptGetCurrentDate PromptName = "get_current_date"
)

const (
	ResourceCurrentDate ResourceName = "resource://current-date"
	ResourceVersion     ResourceName = "resource://version"
	ResourcePing        ResourceName = "resource://ping"
	ResourceUserInfo    ResourceName = "resource://user-info"
)
