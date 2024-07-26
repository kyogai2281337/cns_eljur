package constructor

import (
	"errors"
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
			//fmt.Printf("Teacher %s is not available\n", teacher.Name)
			return false
		}
	}
	return true
}

func (s *Schedule) _isAvailableCabinet(cabinet *model.Cabinet) bool {
	for _, _metaPairCabinet := range s._metaCabinetPair {
		fmt.Println(_metaPairCabinet.Name)
		if _metaPairCabinet == cabinet {
			//fmt.Printf("Cabinet %s is not available\n", cabinet.Name)
			return false
		}
	}
	return true
}

func (s *Schedule) _isAvailableGroup(group *model.Group) bool {
	if s._metaGroupDay[group] >= s.MaxGroupLecturesForDay {
		fmt.Printf("Group %s has reached the maximum number of lectures for the day\n", group.Name)
		return false
	}
	for _, _metaPairGroup := range s._metaGroupPair {
		if group == _metaPairGroup {
			//fmt.Printf("Group %s is not available\n", group.Name)
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
					//fmt.Printf("Found available teacher %s for group %s and subject %s\n", teacher.Name, group.Name, subject.Name)
					return teacher
				}
			}
		}
	}
	//fmt.Printf("No available teacher found for group %s and subject %s\n", group.Name, subject.Name)
	return nil
}

func (s *Schedule) _findLectureData(cabinet *model.Cabinet, group *model.Group) *Lecture {
	//fmt.Printf("Trying to find lecture data for group %s in cabinet %s\n", group.Name, cabinet.Name)
	for subject, lessonsCount := range s.Metrics.Plans[group] {
		//fmt.Printf("Checking subject %s for group %s with %d lessons left\n", subject.Name, group.Name, lessonsCount)
		if subject.RecommendCabType == cabinet.Type && lessonsCount > 0 {
			teacher := s._findTeacherOnGroup(group, subject)
			if teacher != nil {
				if s._isAvailableGroup(group) && s._isAvailableTeacher(teacher) && s._isAvailableCabinet(cabinet) {
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
					//fmt.Printf("Group %s, teacher %s, or cabinet %s is not available\n", group.Name, teacher.Name, cabinet.Name)
				}
			} else {
				//fmt.Printf("No available teacher found for subject %s and group %s\n", subject.Name, group.Name)
			}
		} else {
			//fmt.Printf("Subject %s is not suitable for cabinet %s or no lessons left\n", subject.Name, cabinet.Name)
		}
	}
	//fmt.Printf("No lecture data found for group %s in cabinet %s\n", group.Name, cabinet.Name)
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
							// Проверка, что лекция не назначена в это время
							alreadyAssigned := false
							for _, existingLecture := range s.Main[day][pair] {
								if existingLecture.Cabinet == lecture.Cabinet || existingLecture.Teacher == lecture.Teacher || existingLecture.Group == lecture.Group {
									alreadyAssigned = true
									break
								}
							}
							if !alreadyAssigned {
								s.Main[day][pair] = append(s.Main[day][pair], lecture)
								fmt.Printf("Lecture added: %s, %s, %s, %s\n", lecture.Group.Name, lecture.Subject.Name, lecture.Teacher.Name, lecture.Cabinet.Name)
							}
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

func (s *Schedule) MakeReview() error {
	// Definition of META structs
	_PairGroups := make(map[*model.Group]int)
	_PairTeachers := make(map[*model.Teacher]int)
	for _, groups := range s.Groups {
		_PairGroups[groups] = 0
	}
	for _, teachers := range s.Teachers {
		_PairTeachers[teachers] = 0
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
				if _PairGroups[lecture.Group] == 0 {
					_PairGroups[lecture.Group] = pair
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
				if _PairGroups[lecture.Group]+1 < currentPair {
					s.Metrics.Wins.Groups[lecture.Group][currentDay] += (_PairGroups[lecture.Group] + 1 - currentPair)
				}
				if _PairTeachers[lecture.Teacher]+1 < currentPair {
					s.Metrics.Wins.Teachers[lecture.Teacher][currentDay] += (_PairTeachers[lecture.Teacher] + 1 - currentPair)
				}

				_PairGroups[lecture.Group], _PairTeachers[lecture.Teacher] = currentPair, currentPair
			}
		}
	}

	return nil
}

func (s *Schedule) _incrementObjectMetrics(l *Lecture) error {
	s.Metrics.Plans[l.Group][l.Subject]++
	s.Metrics.TeacherLoads[l.Teacher]++
	err := s.MakeReview()
	if err != nil {
		return err
	}
	return nil
}

func (s *Schedule) _decrementObjectMetrics(l *Lecture) error {
	s.Metrics.Plans[l.Group][l.Subject]--
	s.Metrics.TeacherLoads[l.Teacher]--
	err := s.MakeReview()
	if err != nil {
		return err
	}
	return nil
}

func (s *Schedule) Delete(day, pair int, query interface{}) error {
	switch q := query.(type) {
	case *model.Cabinet:
		for i := range s.Main[day][pair] {
			if s.Main[day][pair][i].Cabinet.ID == q.ID {
				s._incrementObjectMetrics(s.Main[day][pair][i])
				s.Main[day][pair] = append(s.Main[day][pair][:i], s.Main[day][pair][i+1:]...)
				return nil
			}
		}
	case *model.Teacher:
		for i := range s.Main[day][pair] {
			if s.Main[day][pair][i].Teacher.ID == q.ID {
				s._incrementObjectMetrics(s.Main[day][pair][i])
				s.Main[day][pair] = append(s.Main[day][pair][:i], s.Main[day][pair][i+1:]...)
				return nil
			}
		}
	case *model.Group:
		for i := range s.Main[day][pair] {
			if s.Main[day][pair][i].Group.ID == q.ID {
				s._incrementObjectMetrics(s.Main[day][pair][i])
				s.Main[day][pair] = append(s.Main[day][pair][:i], s.Main[day][pair][i+1:]...)
				return nil
			}
		}

	default:
		return errors.New("wrong query type")
	}

	return nil
}

func (s *Schedule) Insert(day, pair int, lecture *Lecture) error {
	s.Main[day][pair] = append(s.Main[day][pair], lecture)

	if err := s._decrementObjectMetrics(lecture); err != nil {
		return err
	}
	return nil
}
