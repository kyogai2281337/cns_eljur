package structures_test

import (
	"encoding/json"
	"testing"

	"github.com/kyogai2281337/cns_eljur/internal/constructor/structures"
)

func TestUpdateMapping(t *testing.T) {
	t.Log("Testing update mapping...")

	// TODO: DO that, to understand, how it works and why it isn`t works

	t.Log("Success!")
}

func TestRenameMapping(t *testing.T) {
	t.Log("Part 1\nMarshalling...")

	styled := structures.UpdateRenameRequest{
		Type: 4,
		Data: structures.RenameReqData{
			Name: "test",
		},
	}
	raw, err := json.Marshal(styled)
	if err != nil {
		t.Errorf("Error marshalling: %s", err.Error())
	}
	if string(raw) != `{"type":4,"data":{"name":"test"}}` {
		t.Errorf("Wrong result: %s", string(raw))
	}
	t.Log("Part 2\nUnmarshalling...")

	after := structures.UpdateRenameRequest{}
	if err := json.Unmarshal(raw, &after); err != nil {
		t.Errorf("Error unmarshalling: %s", err.Error())
	}
	if after != styled {
		t.Errorf("Wrong result, expected %+v, got %+v", styled, after)
	}
	t.Log("Success!")
}
