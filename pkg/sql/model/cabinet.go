package model

type Cabinet struct {
	ID   int64   `bson:"_id,omitempty" json:"id"`
	Name string  `json:"name"`
	Type CabType `json:"type"`
}
