package models

type ParentSegments struct {
	Data []ParentSegment `json:"data"`
}

type ParentSegment struct {
	ID            string                     `json:"id"`
	Type          string                     `json:"type"`
	Attributes    ParentSegmentAttributes    `json:"attributes"`
	Relationships ParentSegmentRelationships `json:"relationships"`
}

type ParentSegmentAttributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ParentSegmentRelationships struct {
	ParentFolder Relationship `json:"parentFolder"`
}
