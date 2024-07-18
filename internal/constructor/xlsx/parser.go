package xlsx

import (
	"fmt"
	"os"
	"path/filepath"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/tealeg/xlsx"
)

func LoadDump(sch constructor.SchCabSorted, fileName string) error {
	file := xlsx.NewFile()

	unsort, err := file.AddSheet("Неотсортированный массив")
	if err != nil {
		return fmt.Errorf("error adding sheet: %w", err)
	}

	head := unsort.AddRow()
	head.AddCell().Value = "День"
	head.AddCell().Value = "Пара"
	head.AddCell().Value = "Кабинет"
	head.AddCell().Value = "Преподаватель"
	head.AddCell().Value = "Группа"
	head.AddCell().Value = "Предмет"
	head.AddCell().Value = "Тип кабинета"

	for dayIndex, day := range sch.Days {
		dayRow := fmt.Sprintf("Day %d", dayIndex+1)
		for pairIndex, pair := range day {
			pairRow := fmt.Sprintf("Pair %d", pairIndex+1)
			for cab, lecture := range pair {
				row := unsort.AddRow()
				row.AddCell().Value = dayRow
				row.AddCell().Value = pairRow
				row.AddCell().Value = cab.Name
				row.AddCell().Value = lecture.Teacher.Name
				row.AddCell().Value = lecture.Group.Name
				row.AddCell().Value = lecture.Subject.Name
				row.AddCell().Value = lecture.Cabinet.Type.String()
			}
		}
	}

	// Проверяем и создаем директорию, если она не существует
	dir := filepath.Dir(fileName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}

	// Сохраняем файл
	err = file.Save(fileName)
	if err != nil {
		return fmt.Errorf("error saving file: %w", err)
	}

	fmt.Println("Excel файл создан успешно.")
	return nil
}

// func LoadFile(sch constructor.SchCabSorted, filename string, blockSize BlockSize) error {
// 	file := xlsx.NewFile()
// 	cabinets := make(map[int][][]string)
// 	var teachers []string
// 	var groups []string

// 	for _, day := range sch.Days {
// 		for _, pair := range day {
// 			for _, lecture := range pair {
// 				if !StringFieldExists(teachers, lecture.Teacher.Name) {
// 					teachers = append(teachers, lecture.Teacher.Name)
// 				}
// 				if !StringFieldExists(groups, lecture.Group.Name) {
// 					groups = append(groups, lecture.Group.Name)
// 				}
// 			}
// 		}
// 	}

// 	// Заполнение данных по кабинетам
// 	for di, day := range sch.Days {
// 		for pi, pair := range day {
// 			for cab, lecture := range pair {
// 				if _, exists := cabinets[cab.Name]; !exists {
// 					cabinets[cab.Name] = make([][]string, 6)
// 					for i := range cabinets[cab.Name] {
// 						cabinets[cab.Name][i] = make([]string, 7) // 7 столбцов для каждого дня
// 					}
// 				}
// 				cabinets[cab.Name][pi][di] = fmt.Sprintf("%s\n%s\n%s", lecture.Teacher.Name, lecture.Group.Name, lecture.Subject.Name)
// 			}
// 		}
// 	}

// 	cabSheet, err := file.AddSheet("По кабинетам")
// 	if err != nil {
// 		return ErrUnableToCreateSheet
// 	}
// 	teacherSheet, err := file.AddSheet("По преподавателям")
// 	if err != nil {
// 		return ErrUnableToCreateSheet
// 	}
// 	groupSheet, err := file.AddSheet("По группам")
// 	if err != nil {
// 		return ErrUnableToCreateSheet
// 	}

// 	// Добавляем названия дней недели в первую строку
// 	daysOfWeek := []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"}
// 	headerRow := cabSheet.AddRow()
// 	headerRow.AddCell().Value = "Кабинет"
// 	for _, day := range daysOfWeek {
// 		headerRow.AddCell().Value = day
// 	}

// 	for cab, schedule := range cabinets {
// 		cabRow := cabSheet.AddRow()
// 		cabCell := cabRow.AddCell()
// 		cabCell.Value = fmt.Sprintf("Cabinet %d", cab)
// 		cabCell.GetStyle().Font.Bold = true

// 		for _, daySchedule := range schedule {
// 			for _, lecture := range daySchedule {
// 				cell := cabRow.AddCell()
// 				cell.Value = lecture

// 				// Обводка ячеек
// 				style := xlsx.NewStyle()
// 				border := xlsx.Border{
// 					Left:   "thin",
// 					Right:  "thin",
// 					Top:    "thin",
// 					Bottom: "thin",
// 				}
// 				style.Border = border
// 				cell.SetStyle(style)
// 			}
// 		}
// 	}

// 	// Заполнение данных по преподавателям
// 	for _, day := range sch.Days {
// 		for _, pair := range day {
// 			for _, lecture := range pair {
// 				if !StringFieldExists(teachers, lecture.Teacher.Name) {
// 					teachers = append(teachers, lecture.Teacher.Name)
// 				}
// 			}
// 		}
// 	}

// 	for _, teacher := range teachers {
// 		row := teacherSheet.AddRow()
// 		row.AddCell().Value = teacher
// 	}

// 	// Заполнение данных по группам
// 	for _, day := range sch.Days {
// 		for _, pair := range day {
// 			for _, lecture := range pair {
// 				if !StringFieldExists(groups, lecture.Group.Name) {
// 					groups = append(groups, lecture.Group.Name)
// 				}
// 			}
// 		}
// 	}

// 	for _, group := range groups {
// 		row := groupSheet.AddRow()
// 		row.AddCell().Value = group
// 	}

// 	// Сохраняем файл
// 	fn := fmt.Sprintf("%s.xlsx", filename)
// 	err = file.Save(fn)
// 	if err != nil {
// 		return ErrUnableToSaveFile
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

func LoadFile(sch constructor.SchCabSorted, filename string) error {
	file := xlsx.NewFile()
	cabinets := make(map[string]map[int][]string)
	teachers := make(map[string]map[int][]string)
	groups := make(map[string]map[int][]string)

	daysInWeek := 6
	pairsPerDay := 6
	daysOfWeek := []string{"ПН", "ВТ", "СР", "ЧТ", "ПТ", "СБ"}

	// Заполнение данных по кабинетам, преподавателям и группам
	for di, day := range sch.Days {
		for pi, pair := range day {
			for cab, lecture := range pair {
				if _, exists := cabinets[cab.Name]; !exists {
					cabinets[cab.Name] = make(map[int][]string)
					for i := 0; i < daysInWeek; i++ {
						cabinets[cab.Name][i] = make([]string, pairsPerDay)
					}
				}
				cabinets[cab.Name][di][pi] = fmt.Sprintf("Teach %s\nGroup %s\nSubject %s", lecture.Teacher.Name, lecture.Group.Name, lecture.Subject.Name)

				if _, exists := teachers[lecture.Teacher.Name]; !exists {
					teachers[lecture.Teacher.Name] = make(map[int][]string)
					for i := 0; i < daysInWeek; i++ {
						teachers[lecture.Teacher.Name][i] = make([]string, pairsPerDay)
					}
				}
				teachers[lecture.Teacher.Name][di][pi] = fmt.Sprintf("Cab %v\nGroup %s\nSubject %s", lecture.Cabinet.Name, lecture.Group.Name, lecture.Subject.Name)

				if _, exists := groups[lecture.Group.Name]; !exists {
					groups[lecture.Group.Name] = make(map[int][]string)
					for i := 0; i < daysInWeek; i++ {
						groups[lecture.Group.Name][i] = make([]string, pairsPerDay)
					}
				}
				groups[lecture.Group.Name][di][pi] = fmt.Sprintf("Cab %v\nTeach %s\nSubject %s", lecture.Cabinet.Name, lecture.Teacher.Name, lecture.Subject.Name)
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
