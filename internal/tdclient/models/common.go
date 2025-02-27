package models

type Relationship struct {
	Data RelationshipData `json:"data"`
}

type RelationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
