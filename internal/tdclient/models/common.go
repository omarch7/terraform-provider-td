package models

type ParentFolder struct {
	Data FolderData `json:"data"`
}

type FolderData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
