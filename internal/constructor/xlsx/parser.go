package xlsx

import (
	"fmt"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/tealeg/xlsx"
)

func LoadDump(sch constructor.SchCabSorted, fileName string) error {
	file := xlsx.NewFile()

	unsort, err := file.AddSheet("Неотсортированный массив")
	if err != nil {
		return err
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
				row.AddCell().Value = fmt.Sprintf("Cabinet %d", cab.Name)
				row.AddCell().Value = lecture.Teacher.Name
				row.AddCell().Value = lecture.Group.Name
				row.AddCell().Value = lecture.Subject.Name
				row.AddCell().Value = fmt.Sprintf("%v", lecture.Cabinet.Type.String())
			}
		}
	}
	// Сохраняем файл
	fn := fmt.Sprintf("%s.xlsx", fileName)
	err = file.Save(fn)
	if err != nil {
		return ErrUnableToSaveFile
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

func LoadFile(sch constructor.SchCabSorted, filename string, blockSize BlockSize) error {
	file := xlsx.NewFile()
	cabinets := make(map[int]map[int][]string)
	var teachers []string
	var groups []string

	for _, day := range sch.Days {
		for _, pair := range day {
			for _, lecture := range pair {
				if !StringFieldExists(teachers, lecture.Teacher.Name) {
					teachers = append(teachers, lecture.Teacher.Name)
				}
				if !StringFieldExists(groups, lecture.Group.Name) {
					groups = append(groups, lecture.Group.Name)
				}
			}
		}
	}

	// Заполнение данных по кабинетам
	for di, day := range sch.Days {
		for pi, pair := range day {
			for cab, lecture := range pair {
				if _, exists := cabinets[cab.Name]; !exists {
					cabinets[cab.Name] = make(map[int][]string)
					for i := 0; i < 7; i++ {
						cabinets[cab.Name][i] = make([]string, 6) // 6 строк для каждого кабинета
					}
				}
				cabinets[cab.Name][di][pi] = fmt.Sprintf("%s\n%s\n%s", lecture.Teacher.Name, lecture.Group.Name, lecture.Subject.Name)
			}
		}
	}

	cabSheet, err := file.AddSheet("По кабинетам")
	if err != nil {
		return ErrUnableToCreateSheet
	}
	teacherSheet, err := file.AddSheet("По преподавателям")
	if err != nil {
		return ErrUnableToCreateSheet
	}
	groupSheet, err := file.AddSheet("По группам")
	if err != nil {
		return ErrUnableToCreateSheet
	}

	// Добавляем названия кабинетов и дней недели
	for cab, schedule := range cabinets {
		cabRow := cabSheet.AddRow()
		cabCell := cabRow.AddCell()
		cabCell.Value = fmt.Sprintf("Кабинет %d:", cab)
		cabCell.GetStyle().Font.Bold = true

		// Создаем заголовок для дней недели
		dayRow := cabSheet.AddRow()
		daysOfWeek := []string{"ПН", "ВТ", "СР", "ЧТ", "ПТ", "СБ"}
		for _, day := range daysOfWeek {
			dayCell := dayRow.AddCell()
			dayCell.Value = day
		}

		// Заполнение данных по дням недели
		for pi := 0; pi < 6; pi++ {
			row := cabSheet.AddRow()
			for di := 0; di < 6; di++ {
				cell := row.AddCell()
				cell.Value = schedule[di][pi]

				// Обводка ячеек
				style := xlsx.NewStyle()
				border := xlsx.Border{
					Left:   "thin",
					Right:  "thin",
					Top:    "thin",
					Bottom: "thin",
				}
				style.Border = border
				style.Alignment = xlsx.Alignment{
					Horizontal: "center",
					Vertical:   "center",
					WrapText:   true,
				}
				cell.SetStyle(style)
			}
		}
	}

	// Заполнение данных по преподавателям
	for _, day := range sch.Days {
		for _, pair := range day {
			for _, lecture := range pair {
				if !StringFieldExists(teachers, lecture.Teacher.Name) {
					teachers = append(teachers, lecture.Teacher.Name)
				}
			}
		}
	}

	for _, teacher := range teachers {
		row := teacherSheet.AddRow()
		row.AddCell().Value = teacher
	}

	// Заполнение данных по группам
	for _, day := range sch.Days {
		for _, pair := range day {
			for _, lecture := range pair {
				if !StringFieldExists(groups, lecture.Group.Name) {
					groups = append(groups, lecture.Group.Name)
				}
			}
		}
	}

	for _, group := range groups {
		row := groupSheet.AddRow()
		row.AddCell().Value = group
	}

	// Сохраняем файл
	fn := fmt.Sprintf("%s.xlsx", filename)
	err = file.Save(fn)
	if err != nil {
		return ErrUnableToSaveFile
	}
	fmt.Println("Excel файл создан успешно.")
	return nil
}
