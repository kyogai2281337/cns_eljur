package xlsx

import (
	"fmt"
	"os"
	"path/filepath"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/tealeg/xlsx"
)

// func LoadDump(sch *constructor.Schedule, fileName string) error {
// 	file := xlsx.NewFile()

// 	unsort, err := file.AddSheet("Неотсортированный массив")
// 	if err != nil {
// 		return fmt.Errorf("error adding sheet: %w", err)
// 	}

// 	head := unsort.AddRow()
// 	head.AddCell().Value = "День"
// 	head.AddCell().Value = "Пара"
// 	head.AddCell().Value = "Кабинет"
// 	head.AddCell().Value = "Преподаватель"
// 	head.AddCell().Value = "Группа"
// 	head.AddCell().Value = "Предмет"
// 	head.AddCell().Value = "Тип кабинета"

// 	for dayIndex, day := range sch.Main {
// 		dayRow := fmt.Sprintf("Day %d", dayIndex+1)
// 		for pairIndex, pair := range day {
// 			pairRow := fmt.Sprintf("Pair %d", pairIndex+1)
// 			for _, lecture := range pair {
// 				row := unsort.AddRow()
// 				row.AddCell().Value = dayRow
// 				row.AddCell().Value = pairRow
// 				row.AddCell().Value = lecture.Cabinet.Name
// 				row.AddCell().Value = lecture.Teacher.Name
// 				row.AddCell().Value = lecture.Group.Name
// 				row.AddCell().Value = lecture.Subject.Name
// 				row.AddCell().Value = lecture.Cabinet.Type.String()
// 			}
// 		}
// 	}

// 	// Проверяем и создаем директорию, если она не существует
// 	dir := filepath.Dir(fileName)
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
// 			return fmt.Errorf("error creating directory: %w", err)
// 		}
// 	}

// 	// Сохраняем файл
// 	err = file.Save(fileName)
// 	if err != nil {
// 		return fmt.Errorf("error saving file: %w", err)
// 	}

// 	fmt.Println("Excel файл создан успешно.")
// 	return nil
// }

func createBorderStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	style.Border = xlsx.Border{
		Left:   "thin",
		Right:  "thin",
		Top:    "thin",
		Bottom: "thin",
	}
	style.Alignment = xlsx.Alignment{
		Horizontal: "center",
		Vertical:   "center",
		WrapText:   true,
	}
	return style
}

func addHeader(sheet *xlsx.Sheet, title string, daysOfWeek []string) {
	headerRow := sheet.AddRow()
	cell := headerRow.AddCell()
	cell.Value = title
	cell.GetStyle().Font.Bold = true

	dayRow := sheet.AddRow()
	for _, day := range daysOfWeek {
		dayCell := dayRow.AddCell()
		dayCell.Value = day
	}
}

func addScheduleData(sheet *xlsx.Sheet, schedule map[int][]string, daysInWeek int, pairsPerDay int, style *xlsx.Style) {
	for pi := 0; pi < pairsPerDay; pi++ {
		row := sheet.AddRow()
		for di := 0; di < daysInWeek; di++ {
			cell := row.AddCell()
			cell.Value = schedule[di][pi]
			cell.SetStyle(style)
		}
	}
}

func createSheetWithData(file *xlsx.File, sheetName string, data map[string]map[int][]string, titlePrefix string, daysOfWeek []string, daysInWeek int, pairsPerDay int, style *xlsx.Style) error {
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnableToCreateSheet, sheetName)
	}

	for item, schedule := range data {
		addHeader(sheet, fmt.Sprintf("%s %s:", titlePrefix, item), daysOfWeek)
		addScheduleData(sheet, schedule, daysInWeek, pairsPerDay, style)
	}

	return nil
}

func LoadFile(sch *constructor.Schedule, filename string) error {
	file := xlsx.NewFile()
	cabinets := make(map[string]map[int][]string)
	teachers := make(map[string]map[int][]string)
	groups := make(map[string]map[int][]string)

	daysInWeek := 6
	pairsPerDay := 6
	daysOfWeek := []string{"ПН", "ВТ", "СР", "ЧТ", "ПТ", "СБ"}

	// Заполнение данных по кабинетам, преподавателям и группам
	for di, day := range sch.Main {
		for pi, pair := range day {
			for _, lecture := range pair {
				// Формируем список групп
				groupsString := ""
				for _, group := range lecture.Groups { // предполагаю, что lecture.Groups это список групп
					groupsString += group.Name + "\n"
				}

				// Убираем лишний перенос строки в конце
				if len(groupsString) > 0 {
					groupsString = groupsString[:len(groupsString)-1]
				}

				// Формируем информацию для кабинетов
				if _, exists := cabinets[lecture.Cabinet.Name]; !exists {
					cabinets[lecture.Cabinet.Name] = make(map[int][]string)
					for i := 0; i < daysInWeek; i++ {
						cabinets[lecture.Cabinet.Name][i] = make([]string, pairsPerDay)
					}
				}
				cabinets[lecture.Cabinet.Name][di][pi] = fmt.Sprintf("Teach %s\nGroup(s):\n%s\nSubject %s", lecture.Teacher.Name, groupsString, lecture.Subject.Name)

				// Формируем информацию для преподавателей
				if _, exists := teachers[lecture.Teacher.Name]; !exists {
					teachers[lecture.Teacher.Name] = make(map[int][]string)
					for i := 0; i < daysInWeek; i++ {
						teachers[lecture.Teacher.Name][i] = make([]string, pairsPerDay)
					}
				}
				teachers[lecture.Teacher.Name][di][pi] = fmt.Sprintf("Cab %v\nGroup(s):\n%s\nSubject %s", lecture.Cabinet.Name, groupsString, lecture.Subject.Name)

				// Формируем информацию для групп
				for _, group := range lecture.Groups {
					if _, exists := groups[group.Name]; !exists {
						groups[group.Name] = make(map[int][]string)
						for i := 0; i < daysInWeek; i++ {
							groups[group.Name][i] = make([]string, pairsPerDay)
						}
					}
					groups[group.Name][di][pi] = fmt.Sprintf("Cab %v\nTeach %s\nSubject %s", lecture.Cabinet.Name, lecture.Teacher.Name, lecture.Subject.Name)
				}
			}
		}
	}

	borderStyle := createBorderStyle()

	// Создание и заполнение листов
	if err := createSheetWithData(file, "По кабинетам", cabinets, "Кабинет", daysOfWeek, daysInWeek, pairsPerDay, borderStyle); err != nil {
		return err
	}
	if err := createSheetWithData(file, "По преподавателям", teachers, "Преподаватель", daysOfWeek, daysInWeek, pairsPerDay, borderStyle); err != nil {
		return err
	}
	if err := createSheetWithData(file, "По группам", groups, "Группа", daysOfWeek, daysInWeek, pairsPerDay, borderStyle); err != nil {
		return err
	}

	// Проверяем и создаем директорию, если она не существует
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}

	// Сохраняем файл
	if err := file.Save(filename); err != nil {
		return fmt.Errorf("error saving file: %w", err)
	}

	fmt.Println("Excel файл создан успешно.")
	return nil
}
