package structures_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/kyogai2281337/cns_eljur/internal/constructor/structures"
	constructor_logic_entrypoint "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/scd"
)

func TestUpdateMapping(t *testing.T) {
	t.Log("Testing update mapping...")

	// TODO: DO that, to understand, how it works and why it isn`t works
	// * Мне если честно так впадлу эти тесты писать, типо блять я сидел, ковырял двое суток компоуз перезапускал,
	// * ради того, чтобы эта поебота в конце концов мне дала понять, что я дефер вьебал прям перед ебучей
	// * горутиной, а потом еще и дерьмом ебаным с маппингом приправилб да в пизду.
	// * Какой-то мусор, но я понял, что такое маппинг и как с этим работать.

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

func TestDeleteMapping(t *testing.T) {
	t.Log("Part 1\nMarshalling...")

	styled := structures.UpdateDeleteRequest{
		Type: constructor_logic_entrypoint.DirDelete,
		Data: structures.DeleteReqData{
			Day:  1,
			Pair: 1,
			Type: "group",
			Name: "test",
		},
	}
	raw, err := json.Marshal(styled)
	if err != nil {
		t.Errorf("Error marshalling: %s", err.Error())
	}
	if string(raw) != `{"type":2,"data":{"day":1,"pair":1,"type":"group","name":"test"}}` {
		t.Errorf("Wrong result: %s", string(raw))
	}

	t.Log("Part 2\nUnmarshalling...")

	after := structures.UpdateDeleteRequest{}
	if err := json.Unmarshal(raw, &after); err != nil {
		t.Errorf("Error unmarshalling: %s", err.Error())
	}
	if after != styled {
		t.Errorf("Wrong result, expected %+v, got %+v", styled, after)
	}
	t.Log("Success!")
}

func TestInsertMapping(t *testing.T) {
	t.Log("Part 1\nMarshalling...")

	styled := structures.UpdateInsertRequest{
		Type: constructor_logic_entrypoint.DirInsert,
		Data: structures.InsertReqData{
			Day:  1,
			Pair: 1,
			Lecture: struct {
				Groups  []string `json:"groups"`
				Teacher string   `json:"teacher"`
				Cabinet string   `json:"cabinet"`
				Subject string   `json:"subject"`
			}{
				Groups:  []string{"group1", "group2"},
				Teacher: "teach",
				Cabinet: "cab",
				Subject: "sub",
			},
		},
	}
	raw, err := json.Marshal(styled)
	if err != nil {
		t.Errorf("Error marshalling: %s", err.Error())
	}
	if string(raw) != `{"type":1,"data":{"day":1,"pair":1,"lecture":{"groups":["group1","group2"],"teacher":"teach","cabinet":"cab","subject":"sub"}}}` {
		t.Errorf("Wrong result: %s", string(raw))
	}

	t.Log("Part 2\nUnmarshalling...")

	after := structures.UpdateInsertRequest{}
	if err := json.Unmarshal(raw, &after); err != nil {
		t.Errorf("Error unmarshalling: %s", err.Error())
	}

	// ! Anonimous function, needed for checking equality
	cond := func(after, styled structures.UpdateInsertRequest) bool {
		if after.Type != styled.Type ||
			after.Data.Day != styled.Data.Day ||
			after.Data.Pair != styled.Data.Pair ||
			!reflect.DeepEqual(after.Data.Lecture, styled.Data.Lecture) {
			return false
		}

		return true
	}

	if !cond(after, styled) {
		t.Errorf("Wrong result, expected %+v, got %+v", styled, after)
	}
	t.Log("Success!")
}
