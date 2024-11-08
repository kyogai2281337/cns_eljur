package methods

// Additional analysis functions and logic can be added here.

// For example, helper functions to maintain data consistency.

// calculateTeacherLoad computes the total lectures for each teacher.
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
