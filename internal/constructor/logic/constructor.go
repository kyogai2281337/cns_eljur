package constructor

import (
	"fmt"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

func (s *Schedule) _cleanUpMetaDay() {
	fmt.Println("Cleaning up meta day data...")
	s._cleanUpMetaPair()
	s._metaGroupDay = make(map[*model.Group]int)
	s._metaTeachDay = make(map[*model.Teacher]int)
}

func (s *Schedule) _cleanUpMetaPair() {
	fmt.Println("Cleaning up meta pair data...")
	s._metaCabinetPair = make([]*model.Cabinet, 0)
	s._metaTeachPair = make([]*model.Teacher, 0)
	s._metaGroupPair = make([]*model.Group, 0)
}

func (s *Schedule) _isAvailableTeacher(teacher *model.Teacher) bool {
	for _, _metaPairTeacher := range s._metaTeachPair {
		if _metaPairTeacher == teacher {
			fmt.Printf("Teacher %s is not available\n", teacher.Name)
			return false
		}
	}
	return true
}

func (s *Schedule) _isAvailableCabinet(cabinet *model.Cabinet) bool {
	for _, _metaPairCabinet := range s._metaCabinetPair {
		if _metaPairCabinet == cabinet {
			fmt.Printf("Cabinet %s is not available\n", cabinet.Name)
			return false
		}
	}
	return true
}

func (s *Schedule) _isAvailableGroup(group *model.Group) bool {
	for _, _metaPairGroup := range s._metaGroupPair {
		if group == _metaPairGroup || s._metaGroupDay[group] >= s.MaxGroupLecturesForDay {
			fmt.Printf("Group %s is not available\n", group.Name)
			return false
		}
	}
	return true
}

func (s *Schedule) _findTeacherOnGroup(group *model.Group, subject *model.Subject) *model.Teacher {
	for _, teacher := range s.Teachers {
		subjects, ok := teacher.Links[group]
		if ok {
			for _, sub := range subjects {
				if sub == subject && s._isAvailableTeacher(teacher) {
					fmt.Printf("Found available teacher %s for group %s and subject %s\n", teacher.Name, group.Name, subject.Name)
					return teacher
				}
			}
		}
	}
	fmt.Printf("No available teacher found for group %s and subject %s\n", group.Name, subject.Name)
	return nil
}

func (s *Schedule) _findLectureData(cabinet *model.Cabinet, group *model.Group) *Lecture {
	fmt.Printf("Trying to find lecture data for group %s in cabinet %s\n", group.Name, cabinet.Name)
	for subject, lessonsCount := range s.Metrics.Plans[group] {
		fmt.Printf("Checking subject %s for group %s with %d lessons left\n", subject.Name, group.Name, lessonsCount)
		if subject.RecommendCabType == cabinet.Type && lessonsCount > 0 {
			teacher := s._findTeacherOnGroup(group, subject)
			if teacher != nil {
				fmt.Printf("Creating lecture for group %s, subject %s, teacher %s, cabinet %s\n", group.Name, subject.Name, teacher.Name, cabinet.Name)

				s._metaCabinetPair = append(s._metaCabinetPair, cabinet)
				s._metaTeachPair = append(s._metaTeachPair, teacher)
				s._metaGroupPair = append(s._metaGroupPair, group)

				s._metaGroupDay[group]++
				s._metaTeachDay[teacher]++

				s.Metrics.Plans[group][subject]--
				s.Metrics.TeacherLoads[teacher]--

				return MakeLecture(subject, cabinet, teacher, group)
			} else {
				fmt.Printf("No available teacher found for subject %s and group %s\n", subject.Name, group.Name)
			}
		} else {
			fmt.Printf("Subject %s is not suitable for cabinet %s or no lessons left\n", subject.Name, cabinet.Name)
		}
	}
	fmt.Printf("No lecture data found for group %s in cabinet %s\n", group.Name, cabinet.Name)
	return nil
}

func (s *Schedule) Make() error {
	for day := range s.Main {
		fmt.Printf("Processing day %d\n", day)
		for pair := range s.Main[day] {
			fmt.Printf("Processing pair %d\n", pair)
			for _, cabinet := range s.Cabinets {
				if s._isAvailableCabinet(cabinet) {

					for _, group := range s.Groups {
						if !s._isAvailableGroup(group) {
							continue
						}
						lecture := s._findLectureData(cabinet, group)
						if lecture != nil {
							s.Main[day][pair] = append(s.Main[day][pair], lecture)
							fmt.Printf("Lecture added: %s, %s, %s, %s\n", lecture.Group.Name, lecture.Subject.Name, lecture.Teacher.Name, lecture.Cabinet.Name)
						}
					}
				}
			}
			s._cleanUpMetaPair()
		}
		for group, i := range s._metaGroupDay {
			fmt.Printf("Group %s has %d lectures\n", group.Name, i)
		}
		s._cleanUpMetaDay()
	}
	return nil
}

func (s *Schedule) String() string {
	response := "_______________________\n\tSCHEDULE_REVIEW:\n_______________________\n"
	for d, day := range s.Main {
		response += "Day: " + fmt.Sprintf("%v\n", d)
		for p, pair := range day {
			response += "\tPair: " + fmt.Sprintf("%v\n", p)
			for k, lecture := range pair {
				response += "\t\t" + fmt.Sprintf("%v. ", k) + lecture.Cabinet.Name + " " + lecture.Teacher.Name + " " + lecture.Group.Name + " " + lecture.Subject.Name + "\n"
			}
		}
	}
	response += "_______________________\n\tPLAN_REVIEW:\n_______________________\n"

	for _, group := range s.Groups {
		response += group.Name + " " + group.Specialization.Name + " "
		for subject, i := range s.Metrics.Plans[group] {
			response += subject.Name + " " + fmt.Sprintf("%v ", i)
		}
		response += "\n"
	}
	return response
}
