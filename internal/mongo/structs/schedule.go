package mongostructures

import (
	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/logic"
)

type CabType int

const (
	Normal CabType = iota
	Flowable
	Laboratory
	Computered
	Sport
)

type MongoWindows struct {
	Groups   map[string][]int `bson:"groups" json:"groups"`
	Teachers map[string][]int `bson:"teachers" json:"teachers"`
}
type MongoMetrics struct {
	Plans        map[string]map[string]int `bson:"plans" json:"plans"`
	Wins         *MongoWindows             `bson:"windows" json:"windows"`
	TeacherLoads map[string]int            `bson:"teacher_loads" json:"teacher_loads"`
}

type MongoLecture struct {
	Cabinet string   `bson:"cabinet" json:"cabinet"`
	Teacher string   `bson:"teacher" json:"teacher"`
	Groups  []string `bson:"group" json:"group"`
	Subject string   `bson:"subject" json:"subject"`
}

type MongoSchedule struct {
	Name string `bson:"name" json:"name"`
	//ID                        primitive.ObjectID  `bson:"_id" json:"-"`
	Groups                    []string            `bson:"groups" json:"groups"`
	Teachers                  []string            `bson:"teachers" json:"teachers"`
	Cabinets                  []string            `bson:"cabinets" json:"cabinets"`
	Plans                     []string            `bson:"plans" json:"plans"`
	Days                      int                 `bson:"days" json:"days"`
	Pairs                     int                 `bson:"pairs" json:"pairs"`
	Metrics                   *MongoMetrics       `bson:"metrics" json:"metrics"`
	Main                      [][][]*MongoLecture `bson:"schedule" json:"schedule"`
	MaxGroupLecturesFor2Weeks int                 `bson:"weeklimit" json:"weeklimit"`
	MaxGroupLecturesForDay    int                 `bson:"daylimit" json:"daylimit"`
}

func ToMongoLecture(l *constructor.Lecture) *MongoLecture {
	grs := make([]string, 0)
	for _, group := range l.Groups {
		grs = append(grs, group.Name)
	}

	return &MongoLecture{
		Cabinet: l.Cabinet.Name,
		Teacher: l.Teacher.Name,
		Groups:  grs,
		Subject: l.Subject.Name,
	}
}

func ToMongoSchedule(s *constructor.Schedule) *MongoSchedule {
	m := &MongoSchedule{}
	if s == nil {
		return m
	}
	m.Groups = make([]string, 0)
	if s.Groups != nil {
		for _, group := range s.Groups {
			if group == nil {
				continue
			}
			m.Groups = append(m.Groups, group.Name)
		}
	}
	m.Teachers = make([]string, 0)
	if s.Teachers != nil {
		for _, teacher := range s.Teachers {
			if teacher == nil {
				continue
			}
			m.Teachers = append(m.Teachers, teacher.Name)
		}
	}
	m.Cabinets = make([]string, 0)
	if s.Cabinets != nil {
		for _, cabinet := range s.Cabinets {
			if cabinet == nil {
				continue
			}
			m.Cabinets = append(m.Cabinets, cabinet.Name)
		}
	}
	m.Plans = make([]string, 0)
	if s.Plans != nil {
		for _, plan := range s.Plans {
			if plan == nil {
				continue
			}
			m.Plans = append(m.Plans, plan.Name)
		}
	}
	m.Days = s.Days
	m.Pairs = s.Pairs
	mm := &MongoMetrics{
		Plans: make(map[string]map[string]int),
		Wins: &MongoWindows{
			Groups:   make(map[string][]int),
			Teachers: make(map[string][]int),
		},
		TeacherLoads: make(map[string]int),
	}
	if s.Metrics != nil {
		for group, groupMap := range s.Metrics.Plans {
			if group == nil || groupMap == nil {
				continue
			}
			mm.Plans[group.Name] = make(map[string]int)
			for subject, count := range groupMap {
				if subject == nil {
					continue
				}
				mm.Plans[group.Name][subject.Name] = count
			}
		}
		for teacher, count := range s.Metrics.TeacherLoads {
			if teacher == nil {
				continue
			}
			mm.TeacherLoads[teacher.Name] = count
		}
		for group, groupMap := range s.Metrics.Wins.Groups {
			if group == nil || groupMap == nil {
				continue
			}
			mm.Wins.Groups[group.Name] = make([]int, s.Days)
			copy(mm.Wins.Groups[group.Name], groupMap)
		}
		for teacher, teacherMap := range s.Metrics.Wins.Teachers {
			if teacher == nil || teacherMap == nil {
				continue
			}
			mm.Wins.Teachers[teacher.Name] = make([]int, s.Days)
			copy(mm.Wins.Teachers[teacher.Name], teacherMap)
		}
	}
	m.Metrics = mm

	if s.Main == nil {
		return m
	}
	m.Main = make([][][]*MongoLecture, 0)
	for di, day := range s.Main {
		if day == nil {
			continue // Skip nil days to avoid nil pointer dereference
		}
		m.Main = append(m.Main, make([][]*MongoLecture, 0))
		for pi, pair := range day {
			if pair == nil {
				continue // Skip nil pairs to avoid nil pointer dereference
			}
			m.Main[di] = append(m.Main[di], make([]*MongoLecture, 0))
			for _, l := range pair {
				if l == nil {
					continue // Skip nil lectures to avoid nil pointer dereference
				}
				mongoLecture := ToMongoLecture(l)
				if mongoLecture == nil {
					continue // Check if ToMongoLecture returns nil, to avoid adding it
				}
				m.Main[di][pi] = append(m.Main[di][pi], mongoLecture)
			}
		}
	}
	m.MaxGroupLecturesFor2Weeks = s.MaxGroupLecturesFor2Weeks
	m.MaxGroupLecturesForDay = s.MaxGroupLecturesForDay
	m.Name = s.Name
	return m
}
