package model

type Subject struct {
	ID               int64   `bson:"_id,omitempty" json:"id"`
	Name             string  `json:"name"`
	RecommendCabType CabType `json:"type"`
}
