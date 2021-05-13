package models

//TagMaster .. mongo record for tag master
type TagMaster struct {
	Tag         string      `json:"tag,omitempty" bson:"tag,omitempty"`
	TagMetadata interface{} `json:"tagMetadata,omitempty" bson:"tagMetadata,omitempty"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	UreatedAt   Time.time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
