package constructor

import (
	"fmt"

	"github.com/kyogai2281337/cns_eljur/pkg/set"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

type Vulnerabilities struct {
	StudentWindows    int
	TeacherWindows    int
	MaxStudentWindows int
	MaxTeacherWindows int
}

func (s *SchCabSorted) FindVulnerabilities(native_groups *set.Set, native_teachers *set.Set) *Vulnerabilities {
	vulnerabilities := &Vulnerabilities{}
	// Iterate over each day
	for _, pairs := range s.Days {
		groupWindows := make(map[*model.Group]int)
		teacherWindows := make(map[*model.Teacher]int)
		firstPairForGroup := make(map[*model.Group]int)
		lastPairForGroup := make(map[*model.Group]int)
		firstPairForTeacher := make(map[*model.Teacher]int)
		lastPairForTeacher := make(map[*model.Teacher]int)

		// Initialize first and last pair maps
		for g := range native_groups.Set {
			group := g.(*model.Group)
			firstPairForGroup[group] = len(pairs)
			lastPairForGroup[group] = -1
		}
		for t := range native_teachers.Set {
			teacher := t.(*model.Teacher)
			firstPairForTeacher[teacher] = len(pairs)
			lastPairForTeacher[teacher] = -1
		}

		// Find first and last pair for each group and teacher
		for pairIndex, pair := range pairs {
			for _, lecture := range pair {
				if lecture != nil {
					group := lecture.Group
					teacher := lecture.Teacher
					if pairIndex < firstPairForGroup[group] {
						firstPairForGroup[group] = pairIndex
					}
					if pairIndex > lastPairForGroup[group] {
						lastPairForGroup[group] = pairIndex
					}
					if pairIndex < firstPairForTeacher[teacher] {
						firstPairForTeacher[teacher] = pairIndex
					}
					if pairIndex > lastPairForTeacher[teacher] {
						lastPairForTeacher[teacher] = pairIndex
					}
				}
			}
		}

		// Count windows for groups
		for group, firstPair := range firstPairForGroup {
			lastPair := lastPairForGroup[group]
			for pairIndex := firstPair + 1; pairIndex < lastPair; pairIndex++ {
				hasGroupLecture := false
				for _, lecture := range pairs[pairIndex] {
					if lecture != nil && lecture.Group == group {
						hasGroupLecture = true
						break
					}
				}
				if !hasGroupLecture {
					groupWindows[group]++
					vulnerabilities.StudentWindows++
				}
			}
		}

		// Count windows for teachers
		for teacher, firstPair := range firstPairForTeacher {
			lastPair := lastPairForTeacher[teacher]
			for pairIndex := firstPair + 1; pairIndex < lastPair; pairIndex++ {
				hasLecture := false
				for _, lecture := range pairs[pairIndex] {
					if lecture != nil && lecture.Teacher == teacher {
						hasLecture = true
						break
					}
				}
				if !hasLecture {
					teacherWindows[teacher]++
					vulnerabilities.TeacherWindows++
				}
			}
		}

		// Update max windows
		for _, count := range groupWindows {
			if count > vulnerabilities.MaxStudentWindows {
				vulnerabilities.MaxStudentWindows = count
			}
		}
		for _, count := range teacherWindows {
			if count > vulnerabilities.MaxTeacherWindows {
				vulnerabilities.MaxTeacherWindows = count
			}
		}
	}

	return vulnerabilities
}

func (v *Vulnerabilities) Out() {
	fmt.Printf("Vulnerabilities:\n")
	fmt.Printf("Total student windows: %v\n", v.StudentWindows)
	fmt.Printf("Max student windows in a day: %v\n", v.MaxStudentWindows)
	fmt.Printf("Total teacher windows: %v\n", v.TeacherWindows)
	fmt.Printf("Max teacher windows in a day: %v\n", v.MaxTeacherWindows)
}

func (s *SchCabSorted) ChangeTeacherOnPair(new *model.Teacher, cab *model.Cabinet, day int, pair int) error {
	if _, ok := s.Days[day][pair][cab]; !ok {
		return fmt.Errorf("there is no pair %v in day %v", pair, day)
	}

	for _, lecture := range s.Days[day][pair] {
		if lecture.Teacher == new {
			return fmt.Errorf("teacher %v is already in pair %v in day %v", new.Name, pair, day)
		}
	}
	s.Days[day][pair][cab].Teacher.RecommendSchCap_--
	s.Days[day][pair][cab].Teacher = new
	new.RecommendSchCap_++
	return nil
}

func (s *SchCabSorted) Delete(day int, pair int, query interface{}) error {
	lectures := s.Days[day][pair]
	switch query := query.(type) {
	case *model.Teacher:
		for _, lecture := range lectures {
			if lecture.Teacher == query {
				delete(s.Days[day][pair], lecture.Cabinet)
			}
		}
	case *model.Group:
		for _, lecture := range lectures {
			if lecture.Group == query {
				delete(s.Days[day][pair], lecture.Cabinet)
			}
		}
	case *model.Cabinet:
		delete(s.Days[day][pair], query)
	default:
		return fmt.Errorf("unknown query type: %T", query)
	}

	return nil
}

func (s *SchCabSorted) Find(day, pair int, query interface{}) (*Lecture, error) {
	switch query := query.(type) {
	case *model.Teacher:
		for _, lecture := range s.Days[day][pair] {
			if lecture.Teacher == query {
				return lecture, nil
			}
		}
	case *model.Group:
		for _, lecture := range s.Days[day][pair] {
			if lecture.Group == query {
				return lecture, nil
			}
		}
	case *model.Cabinet:
		if lecture, ok := s.Days[day][pair][query]; ok {
			return lecture, nil
		}
	}
	return nil, fmt.Errorf("unknown query type: %T", query)
}

func (s *SchCabSorted) FindAvailableCabinets(allCabinets *set.Set, day int, pair int) *set.Set {
	availableCabinets := make(map[interface{}]struct{})
	for cabinet := range allCabinets.Set {
		availableCabinets[cabinet.(*model.Cabinet)] = struct{}{}
	}

	for _, lecture := range s.Days[day][pair] {
		delete(availableCabinets, lecture.Cabinet)
	}

	return &set.Set{Set: availableCabinets}
}

func (s *SchCabSorted) FindAvailableTeachers(allTeachers *set.Set, day int, pair int) *set.Set {
	availableTeachers := make(map[interface{}]struct{})
	for teacher := range allTeachers.Set {
		availableTeachers[teacher.(*model.Teacher)] = struct{}{}
	}

	for _, lecture := range s.Days[day][pair] {
		delete(availableTeachers, lecture.Teacher)
	}

	return &set.Set{Set: availableTeachers}
}

// func Remove(list []*Lecture, item *Lecture) []*Lecture {
// 	for i, v := range list {
// 		if v == item {
// 			copy(list[i:], list[i+1:])
// 			list[len(list)-1] = nil // обнуляем "хвост"
// 			list = list[:len(list)-1]
// 		}
// 	}
// 	return list
// }

func (v *Vulnerabilities) CondToStop(msw, mtw, tsw, ttw int) bool {
	return v.MaxStudentWindows > msw && v.MaxTeacherWindows > mtw && v.TeacherWindows > tsw && v.StudentWindows > ttw
}
