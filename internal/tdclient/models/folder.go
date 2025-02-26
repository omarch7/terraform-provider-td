package models

type Folders struct {
	Data []Folder `json:"data"`
}

type FolderReponse struct {
	Data Folder `json:"data"`
}

type Folder struct {
	ID            string              `json:"id,omitempty"`
	Type          string              `json:"type"`
	Attributes    FolderAttributes    `json:"attributes"`
	Relationships FolderRelationships `json:"relationships"`
}

type FolderAttributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AudienceId  string `json:"audienceId,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type FolderRelationships struct {
	ParentFolder Relationship `json:"parentFolder"`
}
