package methods

import (
	"encoding/json"
	"fmt"
)

// CustomSchedule is your custom struct mirroring MongoSchedule
type CustomSchedule struct {
	Name                      string
	Groups                    []string
	Teachers                  []string
	Cabinets                  []string
	Plans                     []string
	Days                      int
	Pairs                     int
	Metrics                   *MongoMetrics
	Main                      [][][]*MongoLecture
	MaxGroupLecturesFor2Weeks int
	MaxGroupLecturesForDay    int
}

// LoadCustomScheduleFromJSON loads the schedule from JSON data
func LoadCustomScheduleFromJSON(jsonData string) (*CustomSchedule, error) {
	// Unmarshal into MongoSchedule struct
	var mongoSchedule MongoSchedule
	err := json.Unmarshal([]byte(jsonData), &mongoSchedule)
	if err != nil {
		return nil, err
	}

	// Convert MongoSchedule to CustomSchedule
	customSchedule := &CustomSchedule{
		Name:                      mongoSchedule.Name,
		Groups:                    mongoSchedule.Groups,
		Teachers:                  mongoSchedule.Teachers,
		Cabinets:                  mongoSchedule.Cabinets,
		Plans:                     mongoSchedule.Plans,
		Days:                      mongoSchedule.Days,
		Pairs:                     mongoSchedule.Pairs,
		Metrics:                   mongoSchedule.Metrics,
		Main:                      mongoSchedule.Main,
		MaxGroupLecturesFor2Weeks: mongoSchedule.MaxGroupLecturesFor2Weeks,
		MaxGroupLecturesForDay:    mongoSchedule.MaxGroupLecturesForDay,
	}

	return customSchedule, nil
}

// Insert adds a new lecture to the schedule
func (cs *CustomSchedule) Insert(newLecture *MongoLecture, day int, pair int) error {
	// Ensure day and pair are within bounds
	if day < 0 || day >= cs.Days || pair < 0 || pair >= cs.Pairs {
		return fmt.Errorf("day or pair index out of range")
	}

	// Insert the new lecture
	cs.Main[day][pair] = append(cs.Main[day][pair], newLecture)
	return nil
}

// Rename changes the name of the schedule
func (cs *CustomSchedule) Rename(newName string) error {
	cs.Name = newName
	return nil
}

// Delete removes lectures at a specific day and pair
func (cs *CustomSchedule) Delete(day int, pair int) error {
	// Ensure day and pair are within bounds
	if day < 0 || day >= cs.Days || pair < 0 || pair >= cs.Pairs {
		return fmt.Errorf("day or pair index out of range")
	}

	// Delete all lectures at the specified day and pair
	cs.Main[day][pair] = []*MongoLecture{}
	return nil
}
