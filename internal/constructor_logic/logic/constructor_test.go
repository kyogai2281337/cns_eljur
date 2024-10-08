package constructor_test

import (
	"fmt"
	"testing"

	// constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	// "github.com/kyogai2281337/cns_eljur/internalconstructor_logic/xlsx"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/logic"
	"github.com/kyogai2281337/cns_eljur/internal/constructor_logic/xlsx"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

func BenchmarkScheduleAtomic(b *testing.B) {
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
		Name:     "207",
		Type:     model.Laboratory,
		Capacity: 1,
	}
	be := model.Cabinet{
		Name:     "208",
		Type:     model.Normal,
		Capacity: 1,
	}
	c := model.Cabinet{
		Name:     "209",
		Type:     model.Computered,
		Capacity: 2,
	}
	g := model.Cabinet{
		Name:     "210",
		Type:     model.Normal,
		Capacity: 1,
	}
	e := model.Cabinet{
		Name:     "211",
		Type:     model.Computered,
		Capacity: 2,
	}
	f := model.Cabinet{
		Name:     "212",
		Type:     model.Laboratory,
		Capacity: 2,
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

	schedule := constructor.MakeSchedule("", 6, 6, groupArr, teachArr, cabArr, []*model.Specialization{&speca}, 4, 18)
	err := schedule.Make()
	if err != nil {
		panic(err)
	}

	schedule.MakeReview()
	fmt.Println(schedule)

}

func TestFlowSchedule(t *testing.T) {
	// subjects

	sa := model.Subject{
		Name:             "Go",
		RecommendCabType: model.Flowable,
	}

	//specializations

	speca := model.Specialization{
		Name:    "IT",
		Course:  1,
		EduPlan: map[*model.Subject]int{&sa: 18},
	}

	//cabinets

	ca := model.Cabinet{
		Name:     "207",
		Type:     model.Flowable,
		Capacity: 2,
	}

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

	// teachers

	t1 := model.Teacher{
		Name: "Ivan Ivanov",
		Links: map[*model.Group][]*model.Subject{
			&g1: {
				&sa,
			},
			&g2: {
				&sa,
			},
		},
		RecommendSchCap_: 18,
	}

	schedule := constructor.MakeSchedule("", 6, 7, []*model.Group{&g1, &g2}, []*model.Teacher{&t1}, []*model.Cabinet{&ca}, []*model.Specialization{&speca}, 4, 18)
	err := schedule.Make()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(schedule)
	xlsx.LoadFile(schedule, "schedule.xlsx")

}

func TestFlowNSportSchedule(t *testing.T) {
	// subjects

	sa := model.Subject{
		Name:             "Go",
		RecommendCabType: model.Flowable,
	}

	sb := model.Subject{
		Name:             "Sport",
		RecommendCabType: model.Sport,
	}
	//specializations

	speca := model.Specialization{
		Name:    "IT",
		Course:  1,
		EduPlan: map[*model.Subject]int{&sa: 9, &sb: 9},
	}

	//cabinets

	ca := model.Cabinet{
		Name:     "207",
		Type:     model.Flowable,
		Capacity: 2,
	}
	cb := model.Cabinet{
		Name:     "Training complex",
		Type:     model.Sport,
		Capacity: 2,
	}

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

	// teachers

	t1 := model.Teacher{
		Name: "Ivan Ivanov",
		Links: map[*model.Group][]*model.Subject{
			&g1: {
				&sa,
			},
			&g2: {
				&sa,
			},
		},
		RecommendSchCap_: 9,
	}

	t2 := model.Teacher{
		Name: "Space Connector",
		Links: map[*model.Group][]*model.Subject{
			&g1: {
				&sb,
			},
			&g2: {
				&sb,
			},
		},
		RecommendSchCap_: 9,
	}

	schedule := constructor.MakeSchedule("", 6, 7, []*model.Group{&g1, &g2}, []*model.Teacher{&t1, &t2}, []*model.Cabinet{&ca, &cb}, []*model.Specialization{&speca}, 4, 18)
	err := schedule.Make()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(schedule)
	xlsx.LoadFile(schedule, "schedule.xlsx")

}
