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
	Expand     string `json:"expand"`
	Self       string `json:"self"`
	Id         string `json:"id"`
	Key        string `json:"key"`
	Name       string `json:"name"`
	AvatarUrls struct {
		X48 string `json:"48x48"`
		X24 string `json:"24x24"`
		X16 string `json:"16x16"`
		X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	ProjectTypeKey string   `json:"projectTypeKey"`
	Simplified     bool     `json:"simplified"`
	Style          string   `json:"style"`
	IsPrivate      bool     `json:"isPrivate"`
	Properties     struct{} `json:"properties"`
	EntityId       string   `json:"entityId"`
	Uuid           string   `json:"uuid"`
}

type GetProjectsAPIResponse struct {
	Data []Project
}
