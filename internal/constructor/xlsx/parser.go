package xlsx

import (
	"fmt"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/tealeg/xlsx"
)

func LoadFile(sch *constructor.SchCabSorted, fileName string) error {
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

// func Parse(file string) (*constructor.SchCabSorted, error) {
// 	xlFile, err := xlsx.OpenFile(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return sch, nil
// }
