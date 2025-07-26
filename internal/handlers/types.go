package handlers

type ResourceName string

const (
	ResourceCurrentDate ResourceName = "resource://current-date"
	ResourceVersion     ResourceName = "resource://version"
	ResourcePing        ResourceName = "resource://ping"
	ResourceUserInfo    ResourceName = "resource://user-info"
)
