package constructor

import (
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupHeap []*model.Group

type Lecture struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Cabinet *model.Cabinet     `json:"cabinet"`
	Teacher *model.Teacher     `json:"teacher"`
	Group   *model.Group       `json:"group"`
	Subject *model.Subject     `json:"subject"`
}

func (h GroupHeap) Len() int           { return len(h) }
func (h GroupHeap) Less(i, j int) bool { return h[i].Priority() < h[j].Priority() }
func (h GroupHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GroupHeap) Push(x interface{}) {
	*h = append(*h, x.(*model.Group))
}

func (h *GroupHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
func (h *GroupHeap) Peek() *model.Group {
	return (*h)[len(*h)-1]
}

func (h *GroupHeap) Find(g *model.Group) int {
	for i, v := range *h {
		if v == g {
			return i
		}
	}
	return -1
}

// Возвращение приоритета по количеству оставшихся пар
