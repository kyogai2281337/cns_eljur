package constructor

import "go.mongodb.org/mongo-driver/bson/primitive"

type CabType int

const (
	Normal CabType = iota
	Flowable
	Laboratory
	Computered
	Sport
)

func (c CabType) String() string {
	switch c {
	case Normal:
		return "Normal"
	case Flowable:
		return "Flowable"
	case Laboratory:
		return "Laboratory"
	case Computered:
		return "Computered"
	case Sport:
		return "Sport"
	}
	return "Unknown"
}

type Lecture struct {
	Cabinet *Cabinet
	Teacher *Teacher
	Group   *Group
	Subject *Subject
}

type Group struct {
	ID             primitive.ObjectID
	Specialization *Specialization
	Name           string
	MaxPairs       int // 18 с начала
	SpecPlan       map[*Subject]int
}

type GroupHeap []*Group

func (h GroupHeap) Len() int           { return len(h) }
func (h GroupHeap) Less(i, j int) bool { return h[i].Priority() < h[j].Priority() }
func (h GroupHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GroupHeap) Push(x interface{}) {
	*h = append(*h, x.(*Group))
}

func (h *GroupHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
func (h *GroupHeap) Peek() *Group {
	return (*h)[len(*h)-1]
}

func (h *GroupHeap) Find(g *Group) int {
	for i, v := range *h {
		if v == g {
			return i
		}
	}
	return -1
}

// Возвращение приоритета по количеству оставшихся пар
func (g *Group) Priority() int {
	return g.MaxPairs // Можно использовать другой критерий для приоритета
}

type Specialization struct {
	ID      primitive.ObjectID
	Name    string
	Course  int
	EduPlan map[*Subject]int
}

type Subject struct {
	ID               primitive.ObjectID
	Name             string
	RecommendCabType CabType
}

type Teacher struct {
	ID               primitive.ObjectID
	Name             string
	Links            map[*Group][]*Subject
	RecommendSchCap_ int //Нагрузка которую он хочет
}

type Cabinet struct {
	ID   primitive.ObjectID
	Name int
	Type CabType
}
