package constructor

import (
	"testing"

	"github.com/kyogai2281337/cns_eljur/pkg/set"
)

func BenchmarkScheduleAtomic(b *testing.B) {
	// subjects

	sa := Subject{
		Name:             "Go",
		RecommendCabType: Computered,
	}
	sb := Subject{
		Name:             "C++",
		RecommendCabType: Normal,
	}
	sc := Subject{
		Name:             "Java",
		RecommendCabType: Laboratory,
	}

	subSet := set.Set{}
	subSet.Push(&sa)
	subSet.Push(&sb)
	subSet.Push(&sc)
	// specializations

	speca := Specialization{
		Name:    "IT",
		Course:  1,
		EduPlan: map[*Subject]int{&sa: 5, &sb: 6, &sc: 7},
	}

	// cabinets
	a := Cabinet{
		Name: 207,
		Type: Laboratory,
	}
	cb := Cabinet{
		Name: 208,
		Type: Normal,
	}
	c := Cabinet{
		Name: 209,
		Type: Computered,
	}
	g := Cabinet{
		Name: 210,
		Type: Normal,
	}
	e := Cabinet{
		Name: 211,
		Type: Computered,
	}
	f := Cabinet{
		Name: 212,
		Type: Laboratory,
	}

	cabSet := &set.Set{}
	cabSet.Push(&a)
	cabSet.Push(&cb)
	cabSet.Push(&c)
	cabSet.Push(&g)
	cabSet.Push(&e)
	cabSet.Push(&f)

	// groups
	g1 := Group{
		Specialization: &speca,
		Name:           "201IT",
		MaxPairs:       18,
	}

	g2 := Group{
		Specialization: &speca,
		Name:           "202IT",
		MaxPairs:       18,
	}

	g3 := Group{
		Specialization: &speca,
		Name:           "203IT",
		MaxPairs:       18,
	}

	grSet := &set.Set{}
	grSet.Push(&g3)
	grSet.Push(&g1)
	grSet.Push(&g2)

	// teachers
	t1 := Teacher{
		Name: "Ivan Ivanov",
		Links: map[*Group][]*Subject{
			&g1: {
				&sa,
				&sb,
			},
			&g2: {
				&sc,
			},
		},
		RecommendSchCap_: 18,
	}

	t2 := Teacher{
		Name: "Petr Petrov",
		Links: map[*Group][]*Subject{
			&g1: {
				&sc,
			},
			&g2: {
				&sa,
				&sb,
			},
		},
		RecommendSchCap_: 18,
	}
	t3 := Teacher{
		Name: "Sidor Sidorov",
		Links: map[*Group][]*Subject{

			&g3: {
				&sa,
				&sb,
				&sc,
			},
		},
		RecommendSchCap_: 18,
	}
	teachSet := &set.Set{}
	teachSet.Push(&t1)
	teachSet.Push(&t2)
	teachSet.Push(&t3)
	// realization
	d := NewSchCab(6, 6, len(cabSet.Set))

	if err := d.AssignLecturesViaCabinet(grSet, teachSet, cabSet); err != nil {
		panic(err)
	}
	d.CheckAndFixGaps()
	d.CheckTeacherLoad(teachSet)
}
