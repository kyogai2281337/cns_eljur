//go:build wasm
// +build wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"desktop/methods"
)

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c // Не даем программе завершиться
}

func registerCallbacks() {
	// Экспортируем функцию analyzeSchedule в JavaScript
	js.Global().Set("analyzeSchedule", js.FuncOf(analyzeSchedule))
}

func analyzeSchedule(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Неверное количество аргументов"
	}
	jsonData := args[0].String()

	// Загружаем расписание из JSON
	schedule, err := methods.LoadCustomScheduleFromJSON(jsonData)
	if err != nil {
		return fmt.Sprintf("Ошибка при загрузке расписания: %v", err)
	}

	// Выполняем анализ расписания
	err = schedule.Analyze()
	if err != nil {
		return fmt.Sprintf("Ошибка при анализе: %v", err)
	}

	// При необходимости преобразуем результат обратно в JSON
	resultJSON, err := json.Marshal(schedule)
	if err != nil {
		return fmt.Sprintf("Ошибка при преобразовании результата в JSON: %v", err)
	}
	return string(resultJSON)
}
