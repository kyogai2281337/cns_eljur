package constructor

import (
	"errors"
	"fmt"

	"github.com/kyogai2281337/cns_eljur/pkg/set"
)

var (
	MaxPairs         = 18
	MaxPairsDayTeach = 6
	MaxPairsDayStud  = 4
	MinTeacherLoad   = 15
	MaxTeacherLoad   = 25
)

// Структура расписания
type Schedule struct {
	Days [6][6]*Lecture // 6 дней по 6 пар
}

// Инициализация пустого расписания
func NewSchedule() *Schedule {
	return &Schedule{}
}

// Вспомогательная функция для поиска преподавателя по предмету
func findTeacherForSubject(teachers *set.Set, group *Group, subject *Subject) *Teacher {
	for t := range teachers.Set {
		if subjects, ok := t.(*Teacher).Links[group]; ok {
			for _, sub := range subjects {
				if sub == subject {
					return t.(*Teacher)
				}
			}
		}
	}
	return nil
}

// Вспомогательная функция для поиска подходящего кабинета
func findCabinetForSubject(cabinets *set.Set, subject *Subject) *Cabinet {
	for cab := range cabinets.Set {
		if cab.(*Cabinet).Type == Normal || (cab.(*Cabinet).Type == Flowable && subject.Flowable) {
			return cab.(*Cabinet)
		}
	}
	return nil
}

// Назначение лекций для группы и преподавателей
func (s *Schedule) AssignLectures(groups *set.Set, teachers *set.Set, cabinets *set.Set) error {
	for group := range groups.Set {
		for subject, hours := range group.(*Group).Specialization.EduPlan {
			teacher := findTeacherForSubject(teachers, group.(*Group), subject)
			if teacher == nil {
				return errors.New(fmt.Sprintf("No teacher found for subject %s in group %s", subject.Name, group.(*Group).Name))
			}
			cabinet := findCabinetForSubject(cabinets, subject)
			if cabinet == nil {
				return errors.New(fmt.Sprintf("No suitable cabinet found for subject %s", subject.Name))
			}
			for hours > 0 {
				assigned := false
				for day := range s.Days {
					for pair := range s.Days[day] {
						if s.Days[day][pair] == nil && group.(*Group).MaxPairs > 0 && teacher.RecommendSchCap_ > 0 {
							// Проверка окон для студентов
							if pair > 0 && s.Days[day][pair-1] == nil {
								continue
							}
							if pair < 5 && s.Days[day][pair+1] != nil && s.Days[day][pair+2] == nil {
								continue
							}

							s.Days[day][pair] = &Lecture{
								Cabinet: cabinet,
								Teacher: teacher,
								Group:   group.(*Group),
								Subject: subject,
							}
							group.(*Group).MaxPairs--
							teacher.RecommendSchCap_--
							hours--
							assigned = true
							break
						}
					}
					if assigned {
						break
					}
				}
				if !assigned {
					return errors.New(fmt.Sprintf("Could not assign all hours for subject %s in group %s", subject.Name, group.(*Group).Name))
				}
			}
		}
	}
	return nil
}

// Проверка и исправление окон
func (s *Schedule) CheckAndFixGaps() {
	for day := range s.Days {
		for pair := 1; pair < len(s.Days[day]); pair++ {
			if s.Days[day][pair] == nil && s.Days[day][pair-1] != nil {
				// Найдено окно, нужно переместить занятие вперед
				for next := pair + 1; next < len(s.Days[day]); next++ {
					if s.Days[day][next] != nil {
						s.Days[day][pair] = s.Days[day][next]
						s.Days[day][next] = nil
						break
					}
				}
			}
		}
	}
}

// Проверка нагрузки преподавателей
func (s *Schedule) CheckTeacherLoad(teachers *set.Set) {
	for teacher := range teachers.Set {
		totalLoad := 0
		for day := range s.Days {
			for pair := range s.Days[day] {
				if s.Days[day][pair] != nil && s.Days[day][pair].Teacher == teacher {
					totalLoad++
				}
			}
		}
		if totalLoad < MinTeacherLoad || totalLoad > MaxTeacherLoad {
			fmt.Printf("Teacher %s has an incorrect load: %d hours\n", teacher.(*Teacher).Name, totalLoad)
			// Исправление нагрузки
			// Дополнительная логика для перераспределения нагрузки
		}
	}
}

func (s *Schedule) Out() {
	for day := range s.Days {
		for pair := range s.Days[day] {
			if s.Days[day][pair] != nil {
				fmt.Printf("day %v, pair %v: \tTeacher: %s \tGroup: %s \tSubject: %s \tCabinet: %v\n", day+1, pair+1, s.Days[day][pair].Teacher.Name, s.Days[day][pair].Group.Name, s.Days[day][pair].Subject.Name, s.Days[day][pair].Cabinet.Name)
			}
		}
	}
}
