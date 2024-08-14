package constructor_test

import (
	"testing"

	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor/logic"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

func TestShuffleCabArray(t *testing.T) {
	cabs := []*model.Cabinet{
		{Name: "Cab1"},
		{Name: "Cab2"},
		{Name: "Cab3"},
	}
	shuffled := constructor.ShuffleCabArray(cabs)
	if len(shuffled) != len(cabs) {
		t.Errorf("Expected %d cabinets, got %d", len(cabs), len(shuffled))
	}
}

func TestShuffleSpeArray(t *testing.T) {
	specializations := []*model.Specialization{
		{Name: "Spec1"},
		{Name: "Spec2"},
		{Name: "Spec3"},
	}
	shuffled := constructor.ShuffleSpeArray(specializations)
	if len(shuffled) != len(specializations) {
		t.Errorf("Expected %d specializations, got %d", len(specializations), len(shuffled))
	}
}

func TestShuffleGroupArray(t *testing.T) {
	groups := []*model.Group{
		{Name: "Group1"},
		{Name: "Group2"},
		{Name: "Group3"},
	}
	shuffled := constructor.ShuffleGroupArray(groups)
	if len(shuffled) != len(groups) {
		t.Errorf("Expected %d groups, got %d", len(groups), len(shuffled))
	}
}

func TestShuffleTeachArray(t *testing.T) {
	teachers := []*model.Teacher{
		{Name: "Teacher1"},
		{Name: "Teacher2"},
		{Name: "Teacher3"},
	}
	shuffled := constructor.ShuffleTeachArray(teachers)
	if len(shuffled) != len(teachers) {
		t.Errorf("Expected %d teachers, got %d", len(teachers), len(shuffled))
	}
}
func TestMakeSchedule(t *testing.T) {
	groups := []*model.Group{
		{Name: "Group1"},
		{Name: "Group2"},
	}
	teachers := []*model.Teacher{
		{Name: "Teacher1"},
		{Name: "Teacher2"},
	}
	cabinets := []*model.Cabinet{
		{Name: "Cab1"},
		{Name: "Cab2"},
	}
	specializations := []*model.Specialization{
		{Name: "Spec1"},
	}

	schedule := constructor.MakeSchedule("Test Schedule", 6, 6, groups, teachers, cabinets, specializations, 4, 18)
	if schedule == nil {
		t.Fatalf("Expected non-nil schedule")
	}
	if schedule.Name != "Test Schedule" {
		t.Errorf("Expected schedule name 'Test Schedule', got %s", schedule.Name)
	}
}

func TestScheduleMake(t *testing.T) {
	groups := []*model.Group{
		{Name: "Group1"},
		{Name: "Group2"},
	}
	teachers := []*model.Teacher{
		{Name: "Teacher1"},
		{Name: "Teacher2"},
	}
	cabinets := []*model.Cabinet{
		{Name: "Cab1"},
		{Name: "Cab2"},
	}
	specializations := []*model.Specialization{
		{Name: "Spec1"},
	}

	schedule := constructor.MakeSchedule("Test Schedule", 6, 6, groups, teachers, cabinets, specializations, 4, 18)
	err := schedule.Make()
	if err != nil {
		t.Fatalf("Failed to make schedule: %v", err)
	}

	if len(schedule.Main) == 0 {
		t.Errorf("Expected non-empty schedule")
	}
}
