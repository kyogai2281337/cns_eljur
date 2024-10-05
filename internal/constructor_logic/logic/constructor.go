package constructor

import (
	"errors"
	"fmt"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

type Done struct{}

// Перемешивает массивы вводных, типо чтобы было красиво
func (s *Schedule) _shuffleArrs() {
	s.Groups = _ShuffleArr(s.Groups)
	s.Cabinets = _ShuffleArr(s.Cabinets)
	s.Plans = _ShuffleArr(s.Plans)
	s.Teachers = _ShuffleArr(s.Teachers)
}

// Чистит буфера для суток
func (s *Schedule) _cleanUpMetaDay() {
	s._cleanUpMetaPair()
	s._metaGroupDay = make(map[string]int)
	s._metaTeachDay = make(map[string]int)
}

// Чистит буфера для пар
func (s *Schedule) _cleanUpMetaPair() {
	s._metaCabinetPair = make(map[*model.Cabinet]int)
	s._metaTeachPair = make([]*model.Teacher, 0)
	s._metaGroupPair = make([]*model.Group, 0)
}

// Проверяет свободен ли препод на момент контекста вызова
func (s *Schedule) _isAvailableTeacher(teacher *model.Teacher) bool {
	if s.Metrics.TeacherLoads[teacher] <= 0 {
		return false
	}
	for _, _metaPairTeacher := range s._metaTeachPair {
		if _metaPairTeacher.Name == teacher.Name {
			return false
		}
	}
	return true
}

// Проверяет свободен ли кабинет на момент контекста вызова
func (s *Schedule) _isAvailableCabinet(cabinet *model.Cabinet) bool {

	for cab, idx := range s._metaCabinetPair {
		if cab == cabinet && idx <= cabinet.Capacity {
			return false
		}
	}
	return true
}

// Проверяет свободна ли группа на момент контекста вызова
func (s *Schedule) _isAvailableGroup(group *model.Group) bool {
	if s._metaGroupDay[group.Name] >= s.MaxGroupLecturesForDay {
		return false
	}
	for _, _metaPairGroup := range s._metaGroupPair {
		if group == _metaPairGroup {
			return false
		}
	}

	return true
}

// Поиск лекции, вся душня тут. Надо бы рассинхронить, но мне так влом если честно,
// да и плюс в производительности скорее потеряем
func (s *Schedule) _findLecDataFlow(cabinet *model.Cabinet, teacher *model.Teacher) *Lecture {
	if cabinet == nil || teacher == nil {
		return nil
	}
	switch cabinet.Type {
	case model.Flowable:
		{
			connMap := make(map[*model.Subject][]*model.Group)
			for group, subArr := range teacher.Links {
				if !s._isAvailableGroup(group) {
					continue
				}
				for _, sub := range subArr {
					if sub.RecommendCabType == model.Flowable {
						if connMap[sub] != nil {
							connMap[sub] = append(connMap[sub], group)
						} else {
							connMap[sub] = []*model.Group{group}
						}
					}
				}
			}

			for sub, grs := range connMap {
				if len(connMap[sub]) == cabinet.Capacity {
					s._metaCabinetPair[cabinet]++
					s._metaTeachPair = append(s._metaTeachPair, teacher)
					s._metaTeachDay[teacher.Name]++

					s.Metrics.TeacherLoads[teacher]--
					for _, group := range grs {
						s._metaGroupPair = append(s._metaGroupPair, group)
						s.Metrics.Plans[group][sub]--
						s._metaGroupDay[group.Name]++
					}
					return MakeFlowableLecture(sub, cabinet, teacher, grs)

				} else {
					continue
				}
			}
		}
	case model.Sport:
		{
			connMap := make(map[*model.Subject][]*model.Group)
			for group, subArr := range teacher.Links {
				if !s._isAvailableGroup(group) {
					continue
				}
				for _, sub := range subArr {
					if sub.RecommendCabType == model.Sport {
						if connMap[sub] != nil {
							connMap[sub] = append(connMap[sub], group)
						} else {
							connMap[sub] = []*model.Group{group}
						}
					}
				}
			}

			for sub, grs := range connMap {
				if len(connMap[sub]) == cabinet.Capacity {
					s._metaCabinetPair[cabinet]++
					s._metaTeachPair = append(s._metaTeachPair, teacher)
					s._metaTeachDay[teacher.Name]++

					s.Metrics.TeacherLoads[teacher]--
					for _, group := range grs {
						s._metaGroupPair = append(s._metaGroupPair, group)
						s.Metrics.Plans[group][sub]--
						s._metaGroupDay[group.Name]++
					}
					return MakeFlowableLecture(sub, cabinet, teacher, grs)
				} else {
					continue
				}
			}
		}
	default:
		{
			for group, subArr := range teacher.Links {
				if s._isAvailableGroup(group) {
					for _, sub := range subArr {
						if sub.RecommendCabType == cabinet.Type {
							s._metaCabinetPair[cabinet]++
							s._metaTeachPair = append(s._metaTeachPair, teacher)
							s._metaTeachDay[teacher.Name]++

							s.Metrics.TeacherLoads[teacher]--
							s.Metrics.Plans[group][sub]--
							s._metaGroupPair = append(s._metaGroupPair, group)
							s._metaGroupDay[group.Name]++

							return MakeLecture(sub, cabinet, teacher, group)
						}
					}
				}
			}
		}
	}
	return nil
}

func (s *Schedule) Make() error {

	for day := range s.Main {
		for pair := range s.Main[day] {
			for _, cabinet := range s.Cabinets {
				if cabinet == nil {
					continue
				}
				if s._isAvailableCabinet(cabinet) {
					for _, teacher := range s.Teachers {
						if teacher == nil {
							continue
						}
						if s._isAvailableTeacher(teacher) {
							lecture := s._findLecDataFlow(cabinet, teacher)
							if lecture != nil {
								s.Main[day][pair] = append(s.Main[day][pair], lecture)
							}
						}
					}
				}
			}
			s._cleanUpMetaPair()
			s._shuffleArrs()
		}
		s._cleanUpMetaDay()
	}
	return nil
}

func (s *Schedule) String() string {

	response := "_______________________\n\tHIGHTIER_STRUCTURES REWIEW:\n_______________________\n"

	for _, group := range s.Groups {
		response += group.Name + " " + group.Specialization.Name + " " + fmt.Sprintf("%v ", group.Specialization.Course) + "\n"
	}
	for _, teacher := range s.Teachers {
		response += teacher.Name + " " + fmt.Sprintf("%v ", teacher.RecommendSchCap_) + " " + "Links:\n"
		for group, link := range teacher.Links {
			response += "\t" + group.Name + "\n"
			for _, sub := range link {
				response += "\t\t" + sub.Name + "\n"
			}
		}
	}
	for _, cabinet := range s.Cabinets {
		response += cabinet.Name + " " + fmt.Sprintf("%v ", cabinet.Type) + "\n"
	}
	for _, spec := range s.Plans {
		response += spec.Name + "\n"
		for sub, paircount := range spec.EduPlan {
			response += "\t" + sub.Name + " " + fmt.Sprintf("%v ", paircount) + "\n"
		}
	}

	response += "_______________________\n\tSCHEDULE_REVIEW:\n_______________________\n"
	for d, day := range s.Main {
		response += "Day: " + fmt.Sprintf("%v\n", d)
		for p, pair := range day {
			response += "\tPair: " + fmt.Sprintf("%v\n", p)
			for _, lecture := range pair {
				response += "\t\t" + lecture.String() + "\n"
			}
		}
	}

	response += "_______________________\n\tPLAN_REVIEW(leftToFill):\n_______________________\n"

	for _, group := range s.Groups {
		response += group.Name + " " + group.Specialization.Name + " "
		for subject, i := range s.Metrics.Plans[group] {
			response += subject.Name + " " + fmt.Sprintf("%v ", i)
		}
		response += "\n"
	}

	response += "_______________________\n\tTEACHERLOAD_REVIEW:\n_______________________\n"

	for _, teacher := range s.Teachers {
		response += teacher.Name + " " + fmt.Sprintf("Total: %v, Undone: %v", teacher.RecommendSchCap_, s.Metrics.TeacherLoads[teacher]) + "\n"
	}

	response += "_______________________\n\tWINDOWS_REVIEW:\n_______________________\n"

	response += "\t\tTeachers:\n"
	for teacher, teacherWins := range s.Metrics.Wins.Teachers {
		response += fmt.Sprintf("%s:\t", teacher.Name)
		for _, i := range teacherWins {
			response += fmt.Sprintf("%v ", i)
		}
		response += "\n"
	}

	response += "\t\tGroups:\n"
	for group, groupWins := range s.Metrics.Wins.Groups {
		response += fmt.Sprintf("%s:\t", group.Name)
		for _, i := range groupWins {
			response += fmt.Sprintf("%v ", i)
		}
		response += "\n"
	}

	return response
}

// MakeReview - create review of schedule
//
// Function generates two maps:
// - _PairGroups - map of groups and their first pair number
// - _PairTeachers - map of teachers and their first pair number
//
// Then it fills s.Metrics.Wins.Groups and s.Metrics.Wins.Teachers
// with the number of windows for each group and teacher
//
// At the end it returns nil if everything is ok, otherwise - error
// !!!ВАЖНО!!!
// Не стоит вызывать этот метод, когда нужно просто вывести ревью атомарного изменения.
// Жрет очень много ресов, в частности времени.
func (s *Schedule) MakeReview() error {
	// Definition of META structs
	_PairGroups := make(map[*model.Group]int)
	_PairTeachers := make(map[*model.Teacher]int)
	for _, group := range s.Groups {
		_PairGroups[group] = 0
	}
	for _, teacher := range s.Teachers {
		_PairTeachers[teacher] = 0
	}

	for _, group := range s.Groups {
		s.Metrics.Wins.Groups[group] = make([]int, s.Days)
	}

	for _, teacher := range s.Teachers {
		s.Metrics.Wins.Teachers[teacher] = make([]int, s.Days)
	}

	// Filling of META structs
	for _, dayLectures := range s.Main {
		for pair, pairLectures := range dayLectures {
			for _, lecture := range pairLectures {
				for _, group := range lecture.Groups {
					if _PairGroups[group] == 0 {
						_PairGroups[group] = pair
					}
				}
				if _PairTeachers[lecture.Teacher] == 0 {
					_PairTeachers[lecture.Teacher] = pair
				}
			}
		}
	}

	for currentDay, dayLectures := range s.Main {
		for currentPair, pairLectures := range dayLectures {
			for _, lecture := range pairLectures {
				grs := make([]*model.Group, 0)
				for _, group := range lecture.Groups {
					if _PairGroups[group]+1 < currentPair {
						s.Metrics.Wins.Groups[group][currentDay] += (_PairGroups[group] + 1 - currentPair)
						grs = append(grs, group)
					}
				}
				if _PairTeachers[lecture.Teacher]+1 < currentPair {
					s.Metrics.Wins.Teachers[lecture.Teacher][currentDay] += (_PairTeachers[lecture.Teacher] + 1 - currentPair)
				}
				for _, group := range grs {
					_PairGroups[group] = currentPair
				}
				_PairTeachers[lecture.Teacher] = currentPair
			}
		}
	}

	return nil
}

func (s *Schedule) _incrementObjectMetrics(l *Lecture) error {
	if l == nil {
		return errors.New("nil pointer dereference: Lecture is nil")
	}
	if l.Groups == nil {
		return errors.New("nil pointer dereference: Group is nil")
	}
	if l.Subject == nil {
		return errors.New("nil pointer dereference: Subject is nil")
	}
	if l.Teacher == nil {
		return errors.New("nil pointer dereference: Teacher is nil")
	}

	// Check for nil pointers in groups
	for _, g := range l.Groups {
		if g == nil {
			return errors.New("nil pointer dereference: Group is nil")
		}
	}

	// Check for nil pointers in subject
	if l.Subject == nil {
		return errors.New("nil pointer dereference: Subject is nil")
	}

	// Check for nil pointers in teacher
	if l.Teacher == nil {
		return errors.New("nil pointer dereference: Teacher is nil")
	}

	// Check for nil maps
	if s.Metrics.Plans == nil {
		return errors.New("nil pointer dereference: Metrics.Plans is nil")
	}
	if s.Metrics.TeacherLoads == nil {
		return errors.New("nil pointer dereference: Metrics.TeacherLoads is nil")
	}

	// Check for nil maps for group
	for _, g := range l.Groups {
		if s.Metrics.Plans[g] == nil {
			s.Metrics.Plans[g] = make(map[*model.Subject]int)
		}
	}

	// Check for nil maps for teacher
	if _, ok := s.Metrics.TeacherLoads[l.Teacher]; !ok {
		s.Metrics.TeacherLoads[l.Teacher] = 0
	}

	// Increment values
	for _, g := range l.Groups {
		s.Metrics.Plans[g][l.Subject]++
	}
	s.Metrics.TeacherLoads[l.Teacher]++

	return nil
}

func (s *Schedule) _decrementObjectMetrics(l *Lecture) error {
	if l == nil {
		return errors.New("nil pointer dereference: Lecture is nil")
	}
	if l.Groups == nil {
		return errors.New("nil pointer dereference: Group is nil")
	}
	if l.Subject == nil {
		return errors.New("nil pointer dereference: Subject is nil")
	}
	if l.Teacher == nil {
		return errors.New("nil pointer dereference: Teacher is nil")
	}

	for _, g := range l.Groups {
		if s.Metrics.Plans[g] == nil {
			return fmt.Errorf("key %v not found in map", g)
		}
		if _, ok := s.Metrics.Plans[g][l.Subject]; !ok {
			return fmt.Errorf("key %v not found in map", l.Subject)
		}

		s.Metrics.Plans[g][l.Subject]--
	}

	if _, ok := s.Metrics.TeacherLoads[l.Teacher]; !ok {
		return fmt.Errorf("key %v not found in map", l.Teacher)
	}

	s.Metrics.TeacherLoads[l.Teacher]--

	return nil
}

func (s *Schedule) Delete(day, pair int, query interface{}) error {
	switch q := query.(type) {
	case *model.Cabinet:
		for i := range s.Main[day][pair] {
			if s.Main[day][pair][i] == nil {
				return errors.New("nil pointer dereference: Lecture is nil")
			}
			if s.Main[day][pair][i].Cabinet == nil {
				return errors.New("nil pointer dereference: Cabinet is nil")
			}
			if s.Main[day][pair][i].Cabinet.ID == q.ID {
				if err := s._incrementObjectMetrics(s.Main[day][pair][i]); err != nil {
					return err
				}
				s.Main[day][pair] = append(s.Main[day][pair][:i], s.Main[day][pair][i+1:]...)
				return nil
			}
		}
	case *model.Teacher:
		for i := range s.Main[day][pair] {
			if s.Main[day][pair][i] == nil {
				return errors.New("nil pointer dereference: Lecture is nil")
			}
			if s.Main[day][pair][i].Teacher == nil {
				return errors.New("nil pointer dereference: Teacher is nil")
			}
			if s.Main[day][pair][i].Teacher.ID == q.ID {
				if err := s._incrementObjectMetrics(s.Main[day][pair][i]); err != nil {
					return err
				}
				s.Main[day][pair] = append(s.Main[day][pair][:i], s.Main[day][pair][i+1:]...)
				return nil
			}
		}
	case *model.Group:
		for i := range s.Main[day][pair] {
			if s.Main[day][pair][i] == nil {
				return errors.New("nil pointer dereference: Lecture is nil")
			}
			if s.Main[day][pair][i].Groups == nil {
				return errors.New("nil pointer dereference: Group is nil")
			}
			for _, group := range s.Main[day][pair][i].Groups {
				if group.ID == q.ID {
					if err := s._incrementObjectMetrics(s.Main[day][pair][i]); err != nil {
						return err
					}
					s.Main[day][pair] = append(s.Main[day][pair][:i], s.Main[day][pair][i+1:]...)
					return nil
				}
			}
		}

	default:
		return errors.New("wrong query type")
	}

	return nil
}

func (s *Schedule) Insert(day, pair int, lecture *Lecture) error {
	if lecture == nil {
		return errors.New("nil pointer dereference: Lecture is nil")
	}

	if s.Main == nil {
		return errors.New("nil pointer dereference: Main is nil")
	}

	if s.Main[day] == nil {
		return fmt.Errorf("key %d not found in map", day)
	}

	if s.Main[day][pair] == nil {
		return fmt.Errorf("key %d not found in map", pair)
	}

	for idx, item := range s.Main[day][pair] {
		if item.Cabinet.Name == lecture.Cabinet.Name {
			return fmt.Errorf("cabinet already exists at %d: %v", idx, item)
		}
		if item.Teacher.Name == lecture.Teacher.Name {
			return fmt.Errorf("teacher already exists at %d: %v", idx, item)
		}
		for _, group := range item.Groups {
			if group.Name == lecture.Groups[0].Name {
				return fmt.Errorf("group already exists at %d: %v", idx, item)
			}
		}
	}
	s.Main[day][pair] = append(s.Main[day][pair], lecture)

	if err := s._decrementObjectMetrics(lecture); err != nil {
		return err
	}
	return nil
}

func (s *Schedule) Normalize() {
	NormalizeAllLinks(s.Groups, s.Teachers, s.Plans)
}

// NormalizeTeacherLinks нормализует ссылки учителей на группы
func NormalizeTeacherLinks(teachers []*model.Teacher, groups []*model.Group) {
	for _, teacher := range teachers {
		newLinks := make(map[*model.Group][]*model.Subject)
		for linkedGroup, subjects := range teacher.Links {
			for _, group := range groups {
				if linkedGroup.Name == group.Name {
					newLinks[group] = subjects
					break
				}
			}
		}
		teacher.Links = newLinks
	}
}

// NormalizeSpecializationLinks нормализует ссылки групп на специализации
func NormalizeSpecializationLinks(groups []*model.Group, plans []*model.Specialization) {
	for _, group := range groups {
		for _, plan := range plans {
			if group.Specialization.Name == plan.Name {
				group.Specialization = plan
				break
			}
		}
	}
}

// NormalizeAllLinks нормализует все ссылки в структурах
func NormalizeAllLinks(groups []*model.Group, teachers []*model.Teacher, plans []*model.Specialization) {
	NormalizeTeacherLinks(teachers, groups)
	NormalizeSpecializationLinks(groups, plans)
}

func (s *Schedule) RecoverLectureData(
	mongoLecture *struct {
		Groups  []string
		Teacher string
		Cabinet string
		Subject string
	}) (*Lecture, error) {
	grs := make([]*model.Group, 0)
	var t *model.Teacher
	var c *model.Cabinet
	var sub *model.Subject

	for _, group := range s.Groups {
		for _, name := range mongoLecture.Groups {
			if group.Name == name {
				grs = append(grs, group)
				break
			}
		}
	}
	for _, teacher := range s.Teachers {
		if teacher.Name == mongoLecture.Teacher {
			t = teacher
			break
		}
	}
	for _, cabinet := range s.Cabinets {
		if cabinet.Name == mongoLecture.Cabinet {
			c = cabinet
			break
		}
	}
	for _, plan := range s.Plans {
		for subject := range plan.EduPlan {
			if subject.Name == mongoLecture.Subject {
				sub = subject
				break
			}
		}
	}
	l := &Lecture{Groups: grs, Teacher: t, Cabinet: c, Subject: sub}
	if len(grs) == 0 || t == nil || c == nil || sub == nil {
		return nil, errors.New("some data lost")
	}
	return l, nil
}

func (s *Schedule) RecoverObject(name, t string) interface{} {
	switch t {
	case "group":
		for _, group := range s.Groups {
			if group.Name == name {
				return group
			}
		}
	case "teacher":
		for _, teacher := range s.Teachers {
			if teacher.Name == name {
				return teacher
			}
		}
	case "cabinet":
		for _, cabinet := range s.Cabinets {
			if cabinet.Name == name {
				return cabinet
			}
		}
	default:
		return nil
	}
	return nil
}
