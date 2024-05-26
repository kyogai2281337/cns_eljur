package constructor

type CabType int

const (
	Normal CabType = iota
	Flowable
	Laboratory
	Computered
	Sport
)

type Lecture struct {
	Cabinet *Cabinet
	Teacher *Teacher
	Group   *Group
	Subject *Subject
}

type Group struct {
	Specialization *Specialization
	Name           string
	MaxPairs       int // 18 с начала
}

type Specialization struct {
	Name    string
	Course  int
	EduPlan map[*Subject]int
}

type Subject struct {
	Name     string
	Flowable bool // возможность быть поточной
}

type Teacher struct {
	Name             string
	Links            map[*Group][]*Subject
	RecommendSchCap_ int //Нагрузка которую он хочет
}

type Cabinet struct {
	Name int
	Type CabType
}
