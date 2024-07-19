package model

type Group struct {
	ID             int64            `bson:"_id,omitempty" json:"id"`
	Specialization *Specialization  `json:"specialization"`
	Name           string           `json:"name"`
	MaxPairs       int              `json:"max_pairs"`
	SpecPlan       map[*Subject]int `json:"-"`
}

func (g *Group) Priority() int {
	return g.MaxPairs // Можно использовать другой критерий для приоритета
}
