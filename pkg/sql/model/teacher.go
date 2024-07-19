package model

type Teacher struct {
	ID               int64                 `bson:"_id,omitempty" json:"id"`
	Name             string                `json:"name"`
	Links            map[*Group][]*Subject `json:"type"`
	LinksID          int64                 `json:"links_id"`
	RecommendSchCap_ int                   `json:"capacity"`
}

// map[*Group.Id][]*Subject.Id

func (t *Teacher) ShorterLinks() map[int64][]int64 {
	response := make(map[int64][]int64)
	for group, subjects := range t.Links {
		response[group.ID] = make([]int64, 0)
		for _, subject := range subjects {
			response[group.ID] = append(response[group.ID], subject.ID)
		}
	}
	return response
}
