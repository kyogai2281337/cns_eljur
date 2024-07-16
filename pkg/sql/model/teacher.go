package model

type Teacher struct {
	ID               int64                 `bson:"_id,omitempty" json:"id"`
	Name             string                `json:"name"`
	Links            map[*Group][]*Subject `json:"links"`
	LinksID          int64                 `json:"links_id"`
	RecommendSchCap_ int                   `json:"capacity"`
}
