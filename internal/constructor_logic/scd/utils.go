package constructor_logic_entrypoint

import (
	"fmt"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/logic"
	"github.com/kyogai2281337/cns_eljur/internal/mongo/primitives"
	mongostructures "github.com/kyogai2281337/cns_eljur/internal/mongo/structs"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

func (w *LogicWorker) RecoverToFull(mongoSchedule *mongostructures.MongoSchedule) (*constructor.Schedule, error) {

	groups := make([]*model.Group, 0)
	teachers := make([]*model.Teacher, 0)
	cabs := make([]*model.Cabinet, 0)
	plans := make([]*model.Specialization, 0)

	for _, group := range mongoSchedule.Groups {
		g, err := w.store.Group().FindByName(group)
		if err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}

	for _, teacher := range mongoSchedule.Teachers {
		t, err := w.store.Teacher().FindByName(teacher)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, t)
	}

	for _, cabinet := range mongoSchedule.Cabinets {
		c, err := w.store.Cabinet().FindByName(cabinet)
		if err != nil {
			return nil, err
		}
		cabs = append(cabs, c)
	}

	for _, plan := range mongoSchedule.Plans {
		p, err := w.store.Specialization().FindByName(plan)
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

func (w *LogicWorker) GetSchedule(dir Directive) (*CacheItem, error) {
	var schedule *CacheItem
	schedule, ok := w.schedBuf[dir.ScheduleID]
	if !ok {
		// Finding, recovering and parsing schedule
		mongoschedule, err := primitives.NewMongoConn().Schedule().Find(dir.ScheduleID)
		if err != nil {
			return nil, err
		}
		sch, err := w.RecoverToFull(mongoschedule)
		if err != nil {
			return nil, err
		}
		schedule = NewCacheItem(sch)

		w.schedBuf[dir.ScheduleID] = schedule
	}
	return schedule, nil
}
