package model

type Specialization struct {
	ID      int64            `bson:"_id,omitempty" json:"id"`
	Name    string           `json:"name"`
	Course  int              `json:"course"`
	EduPlan map[*Subject]int `json:"plan"`
	PlanId  int64            `json:"plan_id"`
}

func (s *Specialization) FindNeedableSubject(t CabType) *Subject {
	for sub := range s.EduPlan {
		if sub.RecommendCabType == t {
			return sub
		}
	}
	return nil
}
