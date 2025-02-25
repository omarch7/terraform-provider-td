package tdclient

type Folders struct {
	Data []Folder `json:"data"`
}

type Folder struct {
	ID            string              `json:"id"`
	Type          string              `json:"type"`
	Attributes    FolderAttributes    `json:"attributes"`
	Relationships FolderRelationships `json:"relationships"`
}

type FolderAttributes struct {
	AudienceId  string `json:"audience_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type FolderRelationships struct {
	Parent FolderParent `json:"parent"`
}

type FolderParent struct {
	Data FolderParentData `json:"data"`
}

type FolderParentData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
