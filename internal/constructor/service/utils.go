package service

import (
	"context"
	"fmt"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/kyogai2281337/cns_eljur/internal/constructor/structures"
	mongoDB "github.com/kyogai2281337/cns_eljur/internal/mongo"
	mongostructures "github.com/kyogai2281337/cns_eljur/internal/mongo/structs"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

func CreateMongoSchedule(schedule *constructor.Schedule) error {
	client, dbCtx, cancel := mongoDB.ConnectMongoDB("")
	_, err := schedule.Make(dbCtx)
	if err != nil {
		return err
	}
	fmt.Println(schedule)
	defer client.Disconnect(dbCtx)
	defer cancel()
	// schedule = dbCtx.Value(constructor.Done{}).(*constructor.Schedule)
	if err = schedule.MakeReview(); err != nil {
		return err
	}
	mongoSchedule := mongostructures.ToMongoSchedule(schedule)

	schedulesCollection := client.Database("eljur").Collection("schedules")
	_, err = schedulesCollection.InsertOne(dbCtx, mongoSchedule)
	if err != nil {
		return err
	}

	return nil
}

func TestSch() error {
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

	schedule := constructor.MakeSchedule("", 6, 6, groupArr, teachArr, cabArr, []*model.Specialization{&speca}, 4, 18)
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

	client, dbCtx, cancel := mongoDB.ConnectMongoDB("")
	defer client.Disconnect(dbCtx)
	defer cancel()

	mongoSchedule := mongostructures.ToMongoSchedule(schedule)

	schedulesCollection := client.Database("eljur").Collection("schedules")
	_, err = schedulesCollection.InsertOne(dbCtx, mongoSchedule)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConstructorController) makeLists(request *structures.CreateRequest) ([]*model.Group, []*model.Cabinet, []*model.Teacher, []*model.Specialization, error) {
	groups := make([]*model.Group, 0)
	cabs := make([]*model.Cabinet, 0)
	teachers := make([]*model.Teacher, 0)
	plans := make([]*model.Specialization, 0)

	for _, groupID := range request.Groups {
		group, err := c.Server.Store.Group().Find(groupID)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		groups = append(groups, group)
	}

	for _, cabinetID := range request.Cabinets {
		cabinet, err := c.Server.Store.Cabinet().Find(cabinetID)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		cabs = append(cabs, cabinet)
	}

	for _, teacherID := range request.Teachers {
		teacher, err := c.Server.Store.Teacher().Find(teacherID)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		teachers = append(teachers, teacher)
	}

	for _, planID := range request.Plans {
		plan, err := c.Server.Store.Specialization().Find(planID)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		plans = append(plans, plan)
	}
	return groups, cabs, teachers, plans, nil
}
