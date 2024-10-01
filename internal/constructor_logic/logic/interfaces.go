package constructor

type ScheduleRepository interface {
	Find(int, int, interface{}) (*Lecture, error) // Найти лекцию среди помоев
	Insert(int, int, *Lecture) error              // Добавить лекцию в расписание
	Delete(int, int, interface{}) error           // Удалить лекцию из расписания
	Make() error                                  // Сгенерировать расписание
	MakeReview() error                            // Сгенерировать отчёт

	String() string
}
