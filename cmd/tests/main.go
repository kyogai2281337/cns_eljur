package main

import (
	"context"
	"fmt"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

func main() {
	// subjects

	sa := model.Subject{
		Name:             "Go",
		RecommendCabType: model.Computered,
	}
	sb := model.Subject{
		Name:             "C++",
		RecommendCabType: model.Normal,
	}
	sc := model.Subject{
		Name:             "Java",
		RecommendCabType: model.Laboratory,
	}
	// specializations

	speca := model.Specialization{
		Name:    "IT",
		Course:  1,
		EduPlan: map[*model.Subject]int{&sa: 5, &sb: 6, &sc: 7},
	}

	// cabinets
	a := model.Cabinet{
		Name: "207",
		Type: model.Laboratory,
	}
	be := model.Cabinet{
		Name: "208",
		Type: model.Normal,
	}
	c := model.Cabinet{
		Name: "209",
		Type: model.Computered,
	}
	g := model.Cabinet{
		Name: "210",
		Type: model.Normal,
	}
	e := model.Cabinet{
		Name: "211",
		Type: model.Computered,
	}
	f := model.Cabinet{
		Name: "212",
		Type: model.Laboratory,
	}
	cabArr := []*model.Cabinet{&a, &be, &c, &g, &e, &f}

	// groups
	g1 := model.Group{
		Specialization: &speca,
		Name:           "201IT",
		MaxPairs:       18,
	}

	g2 := model.Group{
		Specialization: &speca,
		Name:           "202IT",
		MaxPairs:       18,
	}

	g3 := model.Group{
		Specialization: &speca,
		Name:           "203IT",
		MaxPairs:       18,
	}
	groupArr := []*model.Group{&g1, &g2, &g3}
	// teacher
	t1 := model.Teacher{
		Name: "Ivan Ivanov",
		Links: map[*model.Group][]*model.Subject{
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

	t2 := model.Teacher{
		Name: "Petr Petrov",
		Links: map[*model.Group][]*model.Subject{
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
	t3 := model.Teacher{
		Name: "Sidor Sidorov",
		Links: map[*model.Group][]*model.Subject{

			&g3: {
				&sa,
				&sb,
				&sc,
			},
		},
		RecommendSchCap_: 18,
	}
	teachArr := []*model.Teacher{&t1, &t2, &t3}

	schedule := constructor.MakeSchedule(6, 6, groupArr, teachArr, cabArr, []*model.Specialization{&speca}, 4, 18)
	ctx, err := schedule.Make(context.Background())
	if err != nil {
		panic(err)
	}

	schedule, ok := ctx.Value(constructor.Done{}).(*constructor.Schedule)
	if !ok {
		panic("not ok")
	}
	schedule.MakeReview()
	fmt.Println(schedule)
}
