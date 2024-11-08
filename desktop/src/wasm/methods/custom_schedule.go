package methods

import (
	"encoding/json"
	"fmt"
)

// CustomSchedule — пользовательская структура с требуемыми методами
type CustomSchedule struct {
	Name     string   `json:"name"`
	Groups   []string `json:"groups"`
	Teachers []string `json:"teachers"`
	Cabinets []string `json:"cabinets"`
	Plans    []string `json:"plans"`
	Days     int      `json:"days"`
	Pairs    int      `json:"pairs"`
	//Metrics                   *MongoMetrics       `json:"metrics"`
	Main                      [][][]*MongoLecture `json:"schedule"`
	MaxGroupLecturesFor2Weeks int                 `json:"weeklimit"`
	MaxGroupLecturesForDay    int                 `json:"daylimit"`
}

// LoadCustomScheduleFromJSON анмаршалит JSON в CustomSchedule
func LoadCustomScheduleFromJSON(jsonData string) (*CustomSchedule, error) {
	// Анмаршалим в структуру MongoSchedule
	var mongoSchedule MongoSchedule
	err := json.Unmarshal([]byte(jsonData), &mongoSchedule)
	if err != nil {
		return nil, err
	}

	// Преобразуем MongoSchedule в CustomSchedule
	customSchedule := &CustomSchedule{
		Name:                      mongoSchedule.Name,
		Groups:                    mongoSchedule.Groups,
		Teachers:                  mongoSchedule.Teachers,
		Cabinets:                  mongoSchedule.Cabinets,
		Plans:                     mongoSchedule.Plans,
		Days:                      mongoSchedule.Days,
		Pairs:                     mongoSchedule.Pairs,
		Metrics:                   mongoSchedule.Metrics,
		Main:                      mongoSchedule.Main,
		MaxGroupLecturesFor2Weeks: mongoSchedule.MaxGroupLecturesFor2Weeks,
		MaxGroupLecturesForDay:    mongoSchedule.MaxGroupLecturesForDay,
	}

	return customSchedule, nil
}

// Analyze реализует основную логику оценки расписания
// Реализованная задача: метод Analyze
func (cs *CustomSchedule) Analyze() error {
	// Основная логика анализа расписания
	// Например, проверка на конфликты лекций в одно и то же время

	for dayIndex, day := range cs.Main {
		for pairIndex, pair := range day {
			if len(pair) > 1 {
				// Обнаружен конфликт
				fmt.Printf("Обнаружен конфликт на день %d, пара %d\n", dayIndex+1, pairIndex+1)
			}
		}
	}

	// Дополнительная логика анализа может быть добавлена здесь

	return nil
}

// Insert добавляет новую лекцию в расписание
// Реализованная задача: метод Insert
func (cs *CustomSchedule) Insert(newLecture *MongoLecture, day, pair uint8) error {
	// Проверяем допустимость значений day и pair
	if int(day) >= cs.Days || int(pair) >= cs.Pairs {
		return fmt.Errorf("день или пара вне допустимого диапазона")
	}

	// Добавляем новую лекцию
	cs.Main[day][pair] = append(cs.Main[day][pair], newLecture)

	// Поддержание консистентности данных при необходимости

	return nil
}

// Rename переименовывает расписание
// Реализованная задача: метод Rename
func (cs *CustomSchedule) Rename(newName string) error {
	cs.Name = newName
	return nil
}

// Delete удаляет лекции на указанную дату и пару
// Реализованная задача: метод Delete
func (cs *CustomSchedule) Delete(day, pair uint8) error {
	// Проверяем допустимость значений day и pair
	if int(day) >= cs.Days || int(pair) >= cs.Pairs {
		return fmt.Errorf("день или пара вне допустимого диапазона")
	}

	// Удаляем лекции
	cs.Main[day][pair] = []*MongoLecture{}

	// Поддержание консистентности данных при необходимости

	return nil
}
