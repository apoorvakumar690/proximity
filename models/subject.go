package models

//SubjectMaster .. mongo document for subject
type SubjectMaster struct {
	Name      string      `json:"etCode,omitempty" bson:"etCode,omitempty"`
	Domain    string      `json:"domain,omitempty" bson:"domain,omitempty"`
	Metadata  interface{} `json:"metadata,omitempty" bson:"metadata,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	UreatedAt Time.time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
