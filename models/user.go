package models

//User .. mongo record for Users
type User struct {
	Name      string      `json:"name,omitempty" bson:"name,omitempty"`
	Role      string      `json:"role,omitempty" bson:"role,omitempty"`
	UserName  string      `json:"userName,omitempty" bson:"userName,omitempty"`
	Access    []string    `json:"access,omitempty" bson:"access,omitempty"`
	Metadata  interface{} `json:"metadata,omitempty" bson:"metadata,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	UreatedAt Time.time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
