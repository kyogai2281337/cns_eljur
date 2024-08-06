package model

type Teacher struct {
	ID               int64                 `json:"id"`
	Name             string                `json:"name"`
	Links            map[*Group][]*Subject `json:"full_links"`
	LinksID          string                `bson:"_id" json:"links_id"`
	RecommendSchCap_ int                   `json:"capacity"`
	SL               map[int64][]int64     `bson:"links" json:"links"`
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
