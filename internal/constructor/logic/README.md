***Логический пакет конструктора***:
Я не хочу это всё описывать, но мне придётся.
Начинаем с описи модуля:

---

Structure:

```go

type Schedule struct { // size=256 (0x100)
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
func (s *Schedule) Delete(day int, pair int, query interface{}) error
func (s *Schedule) Insert(day int, pair int, lecture *Lecture) error
func (s *Schedule) Make() error
func (s *Schedule) MakeReview() error
func (s *Schedule) Normalize()
func (s *Schedule) RecoverLectureData(mongoLecture *struct{Groups []string; Teacher string; Cabinet string; Subject string}) *Lecture
func (s *Schedule) RecoverObject(name string, t string) interface{}
func (s *Schedule) String() string
// incapsulated next
func (s *Schedule) _cleanUpMetaDay()
func (s *Schedule) _cleanUpMetaPair()
func (s *Schedule) _decrementObjectMetrics(l *Lecture) error
func (s *Schedule) _findLecDataFlow(cabinet *model.Cabinet, teacher *model.Teacher) *Lecture
func (s *Schedule) _incrementObjectMetrics(l *Lecture) error
func (s *Schedule) _isAvailableCabinet(cabinet *model.Cabinet) bool
func (s *Schedule) _isAvailableGroup(group *model.Group) bool
func (s *Schedule) _isAvailableTeacher(teacher *model.Teacher) bool
func (s *Schedule) _shuffleArrs()
```

На данный момент реализовано многое, но требуется провести пару изменений с оптимизацией и добавить механики для повышения гибкости алгоритма.

Из планов:

1. Реализация обработки момента постоянного корелирования нескольких групп в одном потоке, пока не знаю как впихнуть, но думаю осилю;
2. Улучшение асимптотических параметров при выполнении операций генерации и оценки состояния алгоритма.

Usage:

```go
func main() {
	// subjects

	sa := model.Subject{
		Name:             "Go",
		RecommendCabType: model.Flowable,
	}

	//specializations

	speca := model.Specialization{
		Name:    "IT",
		Course:  1,
		EduPlan: map[*model.Subject]int{&sa: 18},
	}

	//cabinets

	ca := model.Cabinet{
		Name:     "207",
		Type:     model.Flowable,
		Capacity: 2,
	}

	// groups

	g1 := model.Group{
		Specialization: &speca,
		Name:           "201IT",
		MaxPairs:       18,
	}

	g2 := model.Group{
		Specialization: &speca,
		Name:           "202IT",
		MaxPairs:       18,
	}

	// teachers

	t1 := model.Teacher{
		Name: "Ivan Ivanov",
		Links: map[*model.Group][]*model.Subject{
			&g1: {
				&sa,
			},
			&g2: {
				&sa,
			},
		},
		RecommendSchCap_: 18,
	}

	schedule := constructor.MakeSchedule("", 6, 7, []*model.Group{&g1, &g2}, []*model.Teacher{&t1}, []*model.Cabinet{&ca}, []*model.Specialization{&speca}, 4, 18)
	err := schedule.Make()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(schedule)
	xlsx.LoadFile(schedule, "schedule.xlsx")

}

```

Правки, форки в алгоритм ПРИВЕТСТВУЮТСЯ!!! Можно рассмотреть аналоги на основе других ЯП.


Разложение вложенных структур:

Windows:

```go
type Windows struct { // size=16 (0x10)
    Groups   map[*model.Group][]int
    Teachers map[*model.Teacher][]int
}
```

Необходима для того, чтобы реестрировать окна для преподавателей и групп по дням недели, реализовано с учетом работы с атомарными операциями.

Metrics:

```go
type Metrics struct { // size=24 (0x18)
    Plans        map[*model.Group]map[*model.Subject]int
    Wins         *Windows
    TeacherLoads map[*model.Teacher]int
}
```

Необходима для хранения состояния выполнения учебных планов и оценки нагрузок преподавателей, в зависимости от их первозданного состояния(при смене нагрузки в БД ничего не поменяется в конструкторе) => TODO.

FlowHolder:

В данный момент эта структура находится в разработке, но предполагается, что в будущем вопрос с постоянными одними и теми же группами в потоке будет закрыт.
