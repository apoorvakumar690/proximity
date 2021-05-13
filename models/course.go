package models

//CourseMaster .. mongo document for the course details
type CourseMaster struct {
	Name      string      `json:"name,omitempty" bson:"name,omitempty"`
	Link      string      `json:"link,omitempty" bson:"link,omitempty"`
	Metadata  interface{} `json:"metadata,omitempty" bson:"metadata,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	UreatedAt Time.time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
