package constructor

import (
	"sync"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

type Windows struct {
	Groups   map[*model.Group][]int
	Teachers map[*model.Teacher][]int
}
type Metrics struct {
	Plans        map[*model.Group]map[*model.Subject]int
	Wins         *Windows
	TeacherLoads map[*model.Teacher]int
}

func MakeWindows() *Windows {
	return &Windows{
		Groups:   make(map[*model.Group][]int),
		Teachers: make(map[*model.Teacher][]int),
	}
}

func MakeMetrics() *Metrics {
	return &Metrics{
		Plans:        make(map[*model.Group]map[*model.Subject]int),
		Wins:         MakeWindows(),
		TeacherLoads: make(map[*model.Teacher]int),
	}
}

func MakeLecture(subject *model.Subject, cabinet *model.Cabinet, teacher *model.Teacher, group *model.Group) *Lecture {
	grs := make([]*model.Group, 0)
	return &Lecture{
		Cabinet: cabinet,
		Teacher: teacher,
		Groups:  append(grs, group),
		Subject: subject,
	}
}

func MakeFlowableLecture(subject *model.Subject, cabinet *model.Cabinet, teacher *model.Teacher, groups []*model.Group) *Lecture {
	return &Lecture{
		Cabinet: cabinet,
		Teacher: teacher,
		Groups:  groups,
		Subject: subject,
	}
}

type Schedule struct {
	Name                      string
	Groups                    []*model.Group
	Teachers                  []*model.Teacher
	Cabinets                  []*model.Cabinet
	Plans                     []*model.Specialization
	Days                      int
	Pairs                     int
	Metrics                   *Metrics
	Main                      [][][]*Lecture
	MaxGroupLecturesFor2Weeks int
	MaxGroupLecturesForDay    int
	// METADATA, DO NOT PARSE TO MONGOLDB
	_metaGroupDay    map[string]int
	_metaTeachDay    map[string]int
	_metaCabinetPair map[*model.Cabinet]int
	_metaTeachPair   []*model.Teacher
	_metaGroupPair   []*model.Group
	WM               *sync.Mutex
}

func MakeSchedule(name string, days, pairs int, groups []*model.Group, teachers []*model.Teacher, cabinets []*model.Cabinet, plans []*model.Specialization, maxDay, maxWeeks int) *Schedule {
	arr := make([][][]*Lecture, days)
	for i := range arr {
		arr[i] = make([][]*Lecture, pairs)
		for j := range arr[i] {
			arr[i][j] = make([]*Lecture, 0)
		}
	}
	metrics := MakeMetrics()
	for _, group := range groups {
		metrics.Plans[group] = make(map[*model.Subject]int)
		for subject, pairsCount := range group.Specialization.EduPlan {
			metrics.Plans[group][subject] = pairsCount
		}
	}
	for _, teacher := range teachers {
		metrics.TeacherLoads[teacher] = teacher.RecommendSchCap_
	}
	s := &Schedule{
		Name:                      name,
		Groups:                    groups,
		Teachers:                  teachers,
		Cabinets:                  cabinets,
		Plans:                     plans,
		Days:                      days,
		Pairs:                     pairs,
		Metrics:                   metrics,
		Main:                      arr,
		MaxGroupLecturesFor2Weeks: maxWeeks,
		MaxGroupLecturesForDay:    maxDay,
		_metaGroupDay:             make(map[string]int),
		_metaTeachDay:             make(map[string]int),
		_metaCabinetPair:          make(map[*model.Cabinet]int),
		_metaTeachPair:            make([]*model.Teacher, 0),
		_metaGroupPair:            make([]*model.Group, 0),
		WM:                        new(sync.Mutex),
	}
	s.Normalize()
	return s
}
