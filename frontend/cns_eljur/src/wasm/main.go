//go:build wasm
// +build wasm

package main

import (
	"syscall/js"

	"frontend/methods"
)

// ValueOf returns x as a JavaScript value:
//
//	| Go                     | JavaScript             |
//	| ---------------------- | ---------------------- |
//	| js.Value               | [its value]            |
//	| js.Func                | function               |
//	| nil                    | null                   |
//	| bool                   | boolean                |
//	| integers and floats    | number                 |
//	| string                 | string                 |
//	| []interface{}          | new array              |
//	| map[string]interface{} | new object             |
//
// Panics if x is not one of the expected types.

func main() {
	regCBacks()
	select {}
}

type JSfunc func(this js.Value, args []js.Value) any

// Inits a JS function callback that takes two arguments and returns their sum.
// In executable context
func regCBacks() {
	js.Global().Set("inc", js.FuncOf(add))
}

// add is a JS function callback that takes two arguments and returns their sum.
// If the number of arguments is not 2, it returns an error string.
func add(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return js.ValueOf("Неверное количество аргументов")
	}
	return methods.ManTask(args[0].Int(), args[1].Int())
}
