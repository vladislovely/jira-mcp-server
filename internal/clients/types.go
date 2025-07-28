package clients

type MeAPIResponse struct {
	AccountID        string `json:"accountId"`
	AccountType      string `json:"accountType"`
	Active           bool   `json:"active"`
	ApplicationRoles struct {
		Items []interface{} `json:"items"`
		Size  int           `json:"size"`
	} `json:"applicationRoles"`
	AvatarURLs struct {
		X16 string `json:"16x16"`
		X24 string `json:"24x24"`
		X32 string `json:"32x32"`
		X48 string `json:"48x48"`
	} `json:"avatarUrls"`
	DisplayName  string `json:"displayName"`
	EmailAddress string `json:"emailAddress"`
	Groups       struct {
		Items []interface{} `json:"items"`
		Size  int           `json:"size"`
	} `json:"groups"`
	Self     string `json:"self"`
	TimeZone string `json:"timeZone"`
	Locale   string `json:"locale"`
}

type Project struct {
	Expand      string `json:"expand"`
	Self        string `json:"self"`
	ID          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarURLs  struct {
		X48 string `json:"48x48"`
		X24 string `json:"24x24"`
		X16 string `json:"16x16"`
		X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	Lead struct {
		Self        string `json:"self"`
		AccountID   string `json:"accountId"`
		AccountType string `json:"accountType"`
		AvatarURLs  struct {
			X48 string `json:"48x48"`
			X24 string `json:"24x24"`
			X16 string `json:"16x16"`
			X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"lead"`
	ProjectTypeKey string   `json:"projectTypeKey"`
	Simplified     bool     `json:"simplified"`
	Style          string   `json:"style"`
	IsPrivate      bool     `json:"isPrivate"`
	Properties     struct{} `json:"properties"`
	EntityID       string   `json:"entityId"`
	UUID           string   `json:"uuid"`
	Archived       bool     `json:"archived"`
	ArchivedDate   string   `json:"archivedDate"`
	ArchivedBy     struct {
		Self       string `json:"self"`
		AccountID  string `json:"accountId"`
		AvatarURLs struct {
			X48 string `json:"48x48"`
			X24 string `json:"24x24"`
			X16 string `json:"16x16"`
			X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"archivedBy"`
}

type GetProjectsAPIResponse struct {
	Data []Project
}

type CreateProject struct {
	AssigneeType       string `json:"assigneeType"`
	Description        string `json:"description"`
	Key                string `json:"key"`
	LeadAccountID      string `json:"leadAccountId"`
	Name               string `json:"name"`
	PermissionScheme   int    `json:"permissionScheme"`
	ProjectTemplateKey string `json:"projectTemplateKey"`
	ProjectTypeKey     string `json:"projectTypeKey"`
}

type CreateProjectAPIResponse struct {
	Self string `json:"self"`
	ID   int    `json:"id"`
	Key  string `json:"key"`
}

type CreateProjectInput struct {
	AccountID          string
	AssigneeType       string
	Name               string
	Description        string
	TaskPrefixKey      string
	ProjectTypeKey     string
	ProjectTemplateKey string
}

type ArchiveProjectInput struct {
	ProjectIDOrKey string
}

type DeleteProjectInput struct {
	ProjectIDOrKey string
}

type RestoreProjectInput struct {
	ProjectIDOrKey string
}

type ArchiveProjectAPIResponse struct{}

type RestoreProjectAPIResponse struct {
	Project
}

type DeleteProjectAPIResponse struct{}
