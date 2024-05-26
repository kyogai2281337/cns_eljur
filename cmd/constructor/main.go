package main

import (
	sch "github.com/kyogai2281337/cns_eljur/internal/constructor"
	"github.com/kyogai2281337/cns_eljur/pkg/set"
)

func main() {
	// subjects

	sa := sch.Subject{
		Name: "Go",
	}
	sb := sch.Subject{
		Name: "C++",
	}
	sc := sch.Subject{
		Name: "Java",
	}

	subSet := set.Set{}
	subSet.Push(&sa)
	subSet.Push(&sb)
	subSet.Push(&sc)
	subSet.Out()

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
		Type: sch.Flowable,
	}
	c := sch.Cabinet{
		Name: 209,
		Type: sch.Laboratory,
	}

	cabSet := &set.Set{}
	cabSet.Push(&a)
	cabSet.Push(&b)
	cabSet.Push(&c)
	cabSet.Out()

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

	grSet := &set.Set{}
	grSet.Push(&g1)
	grSet.Push(&g2)
	grSet.Out()

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
	teachSet := &set.Set{}
	teachSet.Push(&t1)
	teachSet.Push(&t2)
	teachSet.Out()

	// realization
	d := sch.NewSchedule()

	if err := d.AssignLectures(grSet, teachSet, cabSet); err != nil {
		panic(err)
	}

	d.Out()

	// Проверка нагрузки преподавателей
	d.CheckTeacherLoad(teachSet)

	// Проверка окон шкебеде и кабинетов
	d.CheckAndFixGaps()

}
