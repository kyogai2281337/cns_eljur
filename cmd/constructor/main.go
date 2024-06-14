package main

import (
	sch "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/kyogai2281337/cns_eljur/internal/constructor/xlsx"
	"github.com/kyogai2281337/cns_eljur/pkg/set"
)

func main() {
	// subjects

	sa := sch.Subject{
		Name:             "Go",
		RecommendCabType: sch.Computered,
	}
	sb := sch.Subject{
		Name:             "C++",
		RecommendCabType: sch.Normal,
	}
	sc := sch.Subject{
		Name:             "Java",
		RecommendCabType: sch.Laboratory,
	}

	subSet := set.Set{}
	subSet.Push(&sa)
	subSet.Push(&sb)
	subSet.Push(&sc)
	// specializations

	speca := sch.Specialization{
		Name:    "IT",
		Course:  1,
		EduPlan: map[*sch.Subject]int{&sa: 5, &sb: 6, &sc: 7},
	}

	// cabinets
	a := sch.Cabinet{
		Name: 207,
		Type: sch.Normal,
	}
	b := sch.Cabinet{
		Name: 208,
		Type: sch.Normal,
	}
	c := sch.Cabinet{
		Name: 209,
		Type: sch.Computered,
	}
	g := sch.Cabinet{
		Name: 210,
		Type: sch.Normal,
	}
	e := sch.Cabinet{
		Name: 211,
		Type: sch.Normal,
	}
	f := sch.Cabinet{
		Name: 212,
		Type: sch.Laboratory,
	}

	cabSet := &set.Set{}
	cabSet.Push(&a)
	cabSet.Push(&b)
	cabSet.Push(&c)
	cabSet.Push(&g)
	cabSet.Push(&e)
	cabSet.Push(&f)

	// groups
	g1 := sch.Group{
		Specialization: &speca,
		Name:           "201IT",
		MaxPairs:       18,
	}

	g2 := sch.Group{
		Specialization: &speca,
		Name:           "202IT",
		MaxPairs:       18,
	}

	g3 := sch.Group{
		Specialization: &speca,
		Name:           "203IT",
		MaxPairs:       18,
	}

	grSet := &set.Set{}
	grSet.Push(&g3)
	grSet.Push(&g1)
	grSet.Push(&g2)

	// teachers
	t1 := sch.Teacher{
		Name: "Ivan Ivanov",
		Links: map[*sch.Group][]*sch.Subject{
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

	t2 := sch.Teacher{
		Name: "Petr Petrov",
		Links: map[*sch.Group][]*sch.Subject{
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
	t3 := sch.Teacher{
		Name: "Sidor Sidorov",
		Links: map[*sch.Group][]*sch.Subject{

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
	d := sch.NewSchCab(6, 6, len(cabSet.Set))

	if err := d.AssignLecturesViaCabinet(grSet, teachSet, cabSet); err != nil {
		panic(err)
	}
	v := d.FindVulnerabilities(grSet, teachSet)
	d.CheckTeacherLoad(teachSet)

	d.Out()
	v.Out()
	// if err := xlsx.LoadDump(*d, "test"); err != nil {
	// 	panic(err)
	// }
	xlsx.LoadFile(*d, "schedule")

}
