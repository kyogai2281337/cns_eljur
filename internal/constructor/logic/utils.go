package constructor

import (
	"fmt"

	"github.com/kyogai2281337/cns_eljur/pkg/set"
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
		groupWindows := make(map[*Group]int)
		teacherWindows := make(map[*Teacher]int)
		firstPairForGroup := make(map[*Group]int)
		lastPairForGroup := make(map[*Group]int)
		firstPairForTeacher := make(map[*Teacher]int)
		lastPairForTeacher := make(map[*Teacher]int)

		// Initialize first and last pair maps
		for g := range native_groups.Set {
			group := g.(*Group)
			firstPairForGroup[group] = len(pairs)
			lastPairForGroup[group] = -1
		}
		for t := range native_teachers.Set {
			teacher := t.(*Teacher)
			firstPairForTeacher[teacher] = len(pairs)
			lastPairForTeacher[teacher] = -1
		}

		// Find first and last pair for each group and teacher
		for pair := range pairs {
			for _, lecture := range pairs[pair] {
				if lecture != nil {
					group := lecture.Group
					teacher := lecture.Teacher
					if pair < firstPairForGroup[group] {
						firstPairForGroup[group] = pair
					}
					if pair > lastPairForGroup[group] {
						lastPairForGroup[group] = pair
					}
					if pair < firstPairForTeacher[teacher] {
						firstPairForTeacher[teacher] = pair
					}
					if pair > lastPairForTeacher[teacher] {
						lastPairForTeacher[teacher] = pair
					}
				}
			}
		}

		// Count windows for groups
		for group, firstPair := range firstPairForGroup {
			lastPair := lastPairForGroup[group]
			for pair := firstPair + 1; pair < lastPair; pair++ {
				hasGroupLecture := false
				for _, lecture := range pairs[pair] {
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
			for pair := firstPair + 1; pair < lastPair; pair++ {
				hasLecture := false
				for _, lecture := range pairs[pair] {
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

func (s *SchCabSorted) ChangeTeacherOnPair(new *Teacher, cab *Cabinet, day int, pair int) error {
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

func (s *SchCabSorted) FindAvailableCabinets(allCabinets *set.Set, day int, pair int) *set.Set {
	availableCabinets := make(map[interface{}]struct{})
	for cabinet := range allCabinets.Set {
		availableCabinets[cabinet.(*Cabinet)] = struct{}{}
	}

	for _, lecture := range s.Days[day][pair] {
		delete(availableCabinets, lecture.Cabinet)
	}

	return &set.Set{Set: availableCabinets}
}

func (s *SchCabSorted) FindAvailableTeachers(allTeachers *set.Set, day int, pair int) *set.Set {
	availableTeachers := make(map[interface{}]struct{})
	for teacher := range allTeachers.Set {
		availableTeachers[teacher.(*Teacher)] = struct{}{}
	}

	for _, lecture := range s.Days[day][pair] {
		delete(availableTeachers, lecture.Teacher)
	}

	return &set.Set{Set: availableTeachers}
}
