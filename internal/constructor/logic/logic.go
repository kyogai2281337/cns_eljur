package constructor

import (
	"container/heap"
	"fmt"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"

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
type SchCabSorted struct {
	Days [][]map[*model.Cabinet]*Lecture
}

// Инициализация пустого расписания
func NewSchCab(days int, pairs int) *SchCabSorted {

	arr := make([][]map[*model.Cabinet]*Lecture, days)
	for i := range arr {
		arr[i] = make([]map[*model.Cabinet]*Lecture, pairs)
		for j := range arr[i] {
			arr[i][j] = make(map[*model.Cabinet]*Lecture)
		}
	}
	return &SchCabSorted{
		Days: arr,
	}
}

func (s *SchCabSorted) AssignLecturesViaCabinet(native_groups *set.Set, native_teachers *set.Set, native_cabinets *set.Set) error {
	// Инициализация групп и их учебных планов
	groupHeap := &GroupHeap{}
	heap.Init(groupHeap)
	for g := range native_groups.Set {
		group := g.(*model.Group)
		group.SpecPlan = make(map[*model.Subject]int)
		for sub, count := range group.Specialization.EduPlan {
			group.SpecPlan[sub] = count
		}
		heap.Push(groupHeap, group)
	}

	// Инициализация учителей
	teachSet := set.New()
	for t := range native_teachers.Set {
		teachSet.Push(t)
	}

	// Инициализация цикла для расчёта расписания
	for day := range s.Days {
		fmt.Printf("Start day: %v\n", day)

		// Создаем приоритетную очередь групп, использующуюся в пределах дня
		groupHeapDay := &GroupHeap{}
		heap.Init(groupHeapDay)
		for g := range native_groups.Set {
			group := g.(*model.Group)
			heap.Push(groupHeapDay, group)
		}

		// Создаем отсортированный список кабинетов
		cabinetsList := set.New()
		for cab := range native_cabinets.Set {
			cabinetsList.Push(cab)
		}

		// Мапа для подсчета пар для каждой группы в день
		groupsPairsCount := make(map[*model.Group]int)

		for pair := range s.Days[day] {
			fmt.Printf("\tStart pair: %v\n", pair)

			// Создаем копии учителей для текущей пары
			teachers := set.New()
			for t := range teachSet.Set {
				teacher := t.(*model.Teacher)
				teachers.Push(teacher)
			}

			// Создаем копии кабинетов для текущей пары
			cabinets := set.New()
			for cab := range cabinetsList.Set {
				cabinets.Push(cab)
			}

			// Создаём копии доступных групп для текущей пары
			gH := &GroupHeap{}
			heap.Init(gH)
			for _, g := range *groupHeapDay {
				group := g
				heap.Push(gH, group)
			}
			// создание переменной для хранения назначенных групп, чтобы проверять их наличие в коде в пределах пары
			remGroups := &GroupHeap{}
			heap.Init(remGroups)
			for _, g := range *gH {
				heap.Push(remGroups, g)
			}

			// Перебор MDMI типа для кабинетов > групп
			for cabInterface := range cabinets.Set {
				cab := cabInterface.(*model.Cabinet)
				var foundGroup *model.Group
				fmt.Printf("\t  Start cab: %v\n", cab)

				// Поиск подходящей группы для кабинета
				for remGroups.Len() > 0 {
					group := heap.Pop(remGroups).(*model.Group)

					// Проверяем, не достигла ли группа максимального количества пар в день
					if groupsPairsCount[group] >= MaxPairsDayStud {
						fmt.Printf("\t\tskip group %v because it already has %v pairs\n", group.Name, groupsPairsCount[group])
						continue
					}

					// Проверяем, не была ли группа уже назначена на текущую пару
					if _, exists := s.Days[day][pair][cab]; exists {
						fmt.Printf("\t\tskip group %v because it's already assigned in this pair\n", group.Name)
						continue
					}
					if canGroupBeInCabinet(group, cab) {
						foundGroup = group
						break
					}
				}

				if foundGroup == nil {
					fmt.Printf("\t\tskip cab %v because no group for it\n", cab)
					continue
				}

				// Поиск подходящей предметной области
				prefSub := findSubjectForCabinet(cab, foundGroup.SpecPlan)
				if prefSub == nil {
					fmt.Printf("\t\tskip cab %v because no subject for it\n", cab)
					continue
				}

				// Поиск подходящего учителя для предметной области
				t := findTeacherForSubject(teachers, foundGroup, prefSub)
				if t == nil {
					fmt.Printf("\t\tskip cab %v because no teacher for it\n", cab)
					continue
				}

				// Назначаем лекцию
				s.Days[day][pair][cab] = &Lecture{
					Cabinet: cab,
					Teacher: t,
					Group:   foundGroup,
					Subject: prefSub,
				}

				// Обновляем данные групп, учителей и кабинетов
				native_groups.Remove(foundGroup)
				foundGroup.SpecPlan[prefSub]--
				if foundGroup.SpecPlan[prefSub] <= 0 {
					delete(foundGroup.SpecPlan, prefSub)
				}
				native_groups.Push(foundGroup)

				groupsPairsCount[foundGroup]++

				native_teachers.Remove(t)
				t.RecommendSchCap_--
				if t.RecommendSchCap_ > 0 {
					native_teachers.Push(t)
				} else {
					fmt.Println("teacher is overloaded")
				}

				cabinets.Remove(cab)
				teachers.Remove(t)

				// Проверяем, достигла ли группа максимального количества пар в день
				if groupsPairsCount[foundGroup] >= 4 {
					fmt.Printf("\t\t\tend group %v because it already has %v pairs\n", foundGroup.Name, groupsPairsCount[foundGroup])
				} else {
					heap.Push(gH, foundGroup)
				}

				// Продолжение поиска лекций для текущей пары
				if gH.Len() == 0 {
					fmt.Printf("\tend pair %v because all groups finished\n", pair)
					break
				}
			}

			// Проверка окончания пар в текущий день
			if gH.Len() == 0 {
				fmt.Printf("end day %v because all groups finished\n", day)
				break
			}
		}
	}

	fmt.Printf("Scheduling review for the last day:\n")
	for g := range native_groups.Set {
		group := g.(*model.Group)
		fmt.Printf("\tGroup %s:\n", group.Name)
		for key, val := range group.SpecPlan {
			fmt.Printf("\t\tSubject: %v;\n\t\tPairs left: %v\n\n", key.Name, val)
		}
	}
	return nil
}

// Поиск предмета для кабинета
func findSubjectForCabinet(cabinet *model.Cabinet, subjects map[*model.Subject]int) (bestSubject *model.Subject) {
	var maxCount int

	// Перебор предметов для данного типа кабинета
	for sub, val := range subjects {
		if sub.RecommendCabType == cabinet.Type && val > 0 {
			if bestSubject == nil || val > maxCount {
				bestSubject = sub
				maxCount = val
			}
		}
	}

	// Если не найдено подходящего предмета для данного типа кабинета,
	// и тип кабинета Normal, ищем любой предмет с наибольшим количеством пар
	if bestSubject == nil && cabinet.Type == model.Normal {
		for sub, val := range subjects {
			if val > maxCount {
				bestSubject = sub
				maxCount = val
			}
		}
	}

	return bestSubject
}

// Проверка возможности группы быть в кабинете
func canGroupBeInCabinet(group *model.Group, cabinet *model.Cabinet) bool {
	// Поиск предмета, рекомендованного для данного типа кабинета, и проверка наличия оставшихся пар
	for sub, val := range group.SpecPlan {
		if sub.RecommendCabType == cabinet.Type && val > 0 {
			return true
		}
	}

	// Если не найдено подходящего предмета для данного типа кабинета,
	// и тип кабинета Normal, проверяем наличие любых оставшихся пар
	if cabinet.Type == model.Normal {
		for _, val := range group.SpecPlan {
			if val > 0 {
				return true
			}
		}
	}

	return false
}

// Поиск учителя для предмета
func findTeacherForSubject(teachers *set.Set, group *model.Group, subject *model.Subject) *model.Teacher {
	for t := range teachers.Set {
		teacher := t.(*model.Teacher)

		if subjects, ok := teacher.Links[group]; ok {
			for _, sub := range subjects {
				if sub == subject {
					return teacher
				}
			}
		}
	}
	return nil
}

// Проверка нагрузки преподавателей
func (s *SchCabSorted) CheckTeacherLoad(teachers *set.Set) {
	for teacher := range teachers.Set {
		totalLoad := 0
		for day := range s.Days {
			for pair := range s.Days[day] {
				for _, lecture := range s.Days[day][pair] {
					if lecture.Teacher == teacher.(*model.Teacher) {
						totalLoad++
					}
				}
			}
		}
		if totalLoad < MinTeacherLoad || totalLoad > MaxTeacherLoad {
			fmt.Printf("Teacher %s has an incorrect load: %d hours\n", teacher.(*model.Teacher).Name, totalLoad)
			// Дополнительная логика для перераспределения нагрузки
		}
	}
}

// Вывод расписания
func (s *SchCabSorted) Out() {
	for day := range s.Days {
		fmt.Printf("Day %v-----\n", day+1)
		for pair := range s.Days[day] {
			fmt.Printf("\tpair %v-----\n", pair+1)
			for cab := range s.Days[day][pair] {
				lecture := s.Days[day][pair][cab]
				fmt.Printf("\t\t-cab %v: \tLecture: %v; Teacher: %v; Group: %v; Subject: %v\n", cab, lecture.Cabinet, lecture.Teacher.Name, lecture.Group.Name, lecture.Subject)
			}
		}
	}
}
