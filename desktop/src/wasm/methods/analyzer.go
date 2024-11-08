package methods

import (
	"fmt"
)

// Analyze evaluates the current state of the schedule
func (cs *CustomSchedule) Analyze() (Reviewscomms, error) {
	// Initialize the reviews map with all state codes
	reviews := make(Reviewscomms)
	reviews[OK] = []StateInst{}
	reviews[RED] = []StateInst{}
	reviews[ORANGE] = []StateInst{}
	reviews[YELLOW] = []StateInst{}

	// RED: Detect conflicting lectures at the same time slot
	for dayIndex, day := range cs.Main {
		for pairIndex, pair := range day {
			// Collect non-flowable lectures at this time slot
			nonFlowableLectures := []*MongoLecture{}
			for _, lecture := range pair {
				if lecture != nil && !isFlowable(lecture) {
					nonFlowableLectures = append(nonFlowableLectures, lecture)
				}
			}
			// If more than one non-flowable lecture, it's a conflict
			if len(nonFlowableLectures) > 1 {
				// Build the conflict report
				key := fmt.Sprintf("Conflict at day %d, pair %d", dayIndex+1, pairIndex+1)
				value := "Non-flowable lectures overlapping: "
				for _, lecture := range nonFlowableLectures {
					value += fmt.Sprintf("%s (%s), ", lecture.Subject, lecture.Teacher)
				}
				// Add to RED reviews
				reviews[RED] = append(reviews[RED], StateInst{Key: key, Value: value})
			}
		}
	}

	// ORANGE: Detect teacher overload
	teacherLoad := calculateTeacherLoad(cs)
	for teacher, load := range teacherLoad {
		if load > cs.MaxGroupLecturesFor2Weeks {
			key := "Teacher Overload"
			value := fmt.Sprintf("Teacher %s has %d lectures, exceeding limit of %d", teacher, load, cs.MaxGroupLecturesFor2Weeks)
			reviews[ORANGE] = append(reviews[ORANGE], StateInst{Key: key, Value: value})
		}
	}

	// YELLOW: Detect average daily overload
	for dayIndex, day := range cs.Main {
		lectureCount := 0
		for _, pair := range day {
			lectureCount += len(pair)
		}
		avgLectures := float64(lectureCount) / float64(len(cs.Groups))
		if avgLectures > float64(cs.MaxGroupLecturesForDay) {
			key := fmt.Sprintf("Daily Overload on day %d", dayIndex+1)
			value := fmt.Sprintf("Average lectures per group: %.2f", avgLectures)
			reviews[YELLOW] = append(reviews[YELLOW], StateInst{Key: key, Value: value})
		}
	}

	// Return the compiled reviews
	return reviews, nil
}

// isFlowable determines if a lecture is flowable
func isFlowable(lecture *MongoLecture) bool {
	// Define logic to determine if a lecture is flowable
	// For example, check if the cabinet or subject is in a list of flowable items

	flowableSubjects := map[string]bool{
		"Physical Education": true,
		"Art":                true,
	}

	flowableCabinets := map[string]bool{
		"Gym":        true,
		"Art Studio": true,
	}

	if flowableSubjects[lecture.Subject] || flowableCabinets[lecture.Cabinet] {
		return true
	}

	return false
}

// calculateTeacherLoad computes the total lectures for each teacher
func calculateTeacherLoad(cs *CustomSchedule) map[string]int {
	teacherLoad := make(map[string]int)
	for _, day := range cs.Main {
		for _, pair := range day {
			for _, lecture := range pair {
				if lecture != nil {
					teacherLoad[lecture.Teacher]++
				}
			}
		}
	}
	return teacherLoad
}

// AnalyzeSchedule is a function that accepts JSON data and returns the analysis result
func AnalyzeSchedule(jsonData string) (Reviewscomms, error) {
	// Load schedule from JSON data
	schedule, err := LoadCustomScheduleFromJSON(jsonData)
	if err != nil {
		return nil, err
	}

	// Analyze the schedule
	reviews, err := schedule.Analyze()
	if err != nil {
		return nil, err
	}

	return reviews, nil
}
