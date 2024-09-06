package service

import (
	"fmt"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/kyogai2281337/cns_eljur/internal/constructor/structures"
	mongoDB "github.com/kyogai2281337/cns_eljur/internal/mongo"
	mongostructures "github.com/kyogai2281337/cns_eljur/internal/mongo/structs"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

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

func (c *ConstructorController) RecoverToFull(mongoSchedule *mongostructures.MongoSchedule) (*constructor.Schedule, error) {

	groups := make([]*model.Group, 0)
	teachers := make([]*model.Teacher, 0)
	cabs := make([]*model.Cabinet, 0)
	plans := make([]*model.Specialization, 0)

	for _, group := range mongoSchedule.Groups {
		g, err := c.Server.Store.Group().FindByName(group)
		if err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}

	for _, teacher := range mongoSchedule.Teachers {
		t, err := c.Server.Store.Teacher().FindByName(teacher)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, t)
	}

	for _, cabinet := range mongoSchedule.Cabinets {
		c, err := c.Server.Store.Cabinet().FindByName(cabinet)
		if err != nil {
			return nil, err
		}
		cabs = append(cabs, c)
	}

	for _, plan := range mongoSchedule.Plans {
		p, err := c.Server.Store.Specialization().FindByName(plan)
		if err != nil {
			return nil, err
		}
		plans = append(plans, p)
	}

	schedule := constructor.MakeSchedule(mongoSchedule.Name, mongoSchedule.Days, mongoSchedule.Pairs, groups, teachers, cabs, plans, mongoSchedule.MaxGroupLecturesForDay, mongoSchedule.MaxGroupLecturesFor2Weeks)

	//  TODO: Изменение лекции UPD: Готово

	schedule.Main = make([][][]*constructor.Lecture, 0)
	for _, day := range mongoSchedule.Main {
		nDay := make([][]*constructor.Lecture, 0)
		for _, pair := range day {
			nPair := make([]*constructor.Lecture, 0)
			for _, lecture := range pair {
				if lecture == nil {
					continue
				}
				g := make([]*model.Group, 0)
				var t *model.Teacher
				var c *model.Cabinet
				var s *model.Subject
				for _, group := range groups {
					for _, mGrs := range lecture.Groups {
						if group.Name == mGrs {
							g = append(g, group)
							break
						}
					}
				}
				for _, teacher := range teachers {
					if teacher.Name == lecture.Teacher {
						t = teacher
						break
					}
				}
				for _, cabinet := range cabs {
					if cabinet.Name == lecture.Cabinet {
						c = cabinet
						break
					}
				}
				for _, plan := range plans {
					for sub := range plan.EduPlan {
						if sub.Name == lecture.Subject {
							s = sub
							break
						}
					}
				}

				l := constructor.MakeFlowableLecture(s, c, t, g)
				nPair = append(nPair, l)
			}
			nDay = append(nDay, nPair)
		}
		schedule.Main = append(schedule.Main, nDay)
	}
	schedule.Metrics = constructor.MakeMetrics()

	for mGroup, mSub := range mongoSchedule.Metrics.Plans {
		var g *model.Group
		for _, nGroup := range groups {
			if nGroup.Name == mGroup {
				g = nGroup
				break
			}
		}
		if g == nil {
			return nil, fmt.Errorf("group not found: %s", mGroup)
		}

		if schedule.Metrics.Plans[g] == nil {
			schedule.Metrics.Plans[g] = make(map[*model.Subject]int)
		}

		for sub, val := range mSub {
			var s *model.Subject
			for _, plan := range plans {
				for _sub := range plan.EduPlan {
					if _sub.Name == sub {
						s = _sub
						break
					}
				}
				if s != nil {
					break
				}
			}
			if s == nil {
				return nil, fmt.Errorf("subject not found: %s", sub)
			}

			schedule.Metrics.Plans[g][s] = val
		}
	}

	for mTeacher, val := range mongoSchedule.Metrics.TeacherLoads {
		var t *model.Teacher
		for _, nTeacher := range teachers {
			if nTeacher.Name == mTeacher {
				t = nTeacher
			}
		}
		schedule.Metrics.TeacherLoads[t] = val
	}

	for mTeacher, val := range mongoSchedule.Metrics.Wins.Teachers {
		var t *model.Teacher
		for _, nTeacher := range teachers {
			if nTeacher.Name == mTeacher {
				t = nTeacher
			}
		}
		schedule.Metrics.Wins.Teachers[t] = val
	}

	for mGroup, val := range mongoSchedule.Metrics.Wins.Groups {
		var g *model.Group
		for _, nGroup := range groups {
			if nGroup.Name == mGroup {
				g = nGroup
			}
		}
		schedule.Metrics.Wins.Groups[g] = val
	}

	schedule.Normalize()
	return schedule, nil
}
