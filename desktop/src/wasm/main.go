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
	<-c
}

func registerCallbacks() {
	js.Global().Set("analyzeSchedule", js.FuncOf(analyzeSchedule))
}

func analyzeSchedule(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Invalid number of arguments"
	}
	jsonData := args[0].String()
	reviews, err := methods.AnalyzeSchedule(jsonData)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	// Convert reviews to JSON
	reviewsJSON, err := json.Marshal(reviews)
	if err != nil {
		return fmt.Sprintf("Error converting reviews to JSON: %v", err)
	}
	return string(reviewsJSON)
}
