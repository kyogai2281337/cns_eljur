//go:build wasm
// +build wasm

package main_test

import (
	"testing"

	"frontend/methods"
)

func TestAdd(t *testing.T) {
	if methods.ManTask(1, 2) != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", methods.ManTask(1, 2))
	}
	t.Log("Add(1, 2) = 3")
}
