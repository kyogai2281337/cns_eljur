package methods

import (
	"testing"
)

func TestMongoScheduleReviews(t *testing.T) {
	tests := []struct {
		name         string
		schedule     MongoSchedule
		expectedCode StateCode
	}{
		{
			name: "2 instances on 1 time (RED case)",
			schedule: MongoSchedule{
				Groups:   []string{"group1", "group2"},
				Teachers: []string{"teacher1", "teacher2"},
				Cabinets: []string{"cabinet1"},
				Plans:    []string{"plan1"},
				Days:     5,
				Pairs:    4,
				Main: [][][]*MongoLecture{
					{
						{
							&MongoLecture{
								Cabinet: "cabinet1",
								Teacher: "teacher1",
								Groups:  []string{"group1", "group2"},
								Subject: "Math",
							},
						},
					},
					{
						{
							&MongoLecture{
								Cabinet: "cabinet1",
								Teacher: "teacher1",
								Groups:  []string{"group1"},
								Subject: "Physics",
							},
						},
					},
				},
			},
			expectedCode: RED,
		},
		{
			name: "Overload due to participation (ORANGE case)",
			schedule: MongoSchedule{
				Groups:   []string{"group1"},
				Teachers: []string{"teacher1"},
				Cabinets: []string{"cabinet1"},
				Plans:    []string{"plan1"},
				Days:     5,
				Pairs:    4,
				Metrics: &MongoMetrics{
					TeacherLoads: map[string]int{
						"teacher1": 12,
					},
				},
				Main: [][][]*MongoLecture{
					{
						{
							&MongoLecture{
								Cabinet: "cabinet1",
								Teacher: "teacher1",
								Groups:  []string{"group1"},
								Subject: "Math",
							},
						},
					},
				},
			},
			expectedCode: ORANGE,
		},
		{
			name: "Average daily overload (YELLOW case)",
			schedule: MongoSchedule{
				Groups:   []string{"group1"},
				Teachers: []string{"teacher1"},
				Cabinets: []string{"cabinet1"},
				Plans:    []string{"plan1"},
				Days:     5,
				Pairs:    4,
				Metrics: &MongoMetrics{
					TeacherLoads: map[string]int{
						"teacher1": 6,
					},
				},
				Main: [][][]*MongoLecture{
					{
						{
							&MongoLecture{
								Cabinet: "cabinet1",
								Teacher: "teacher1",
								Groups:  []string{"group1"},
								Subject: "Math",
							},
						},
					},
				},
			},
			expectedCode: YELLOW,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stateCode := evaluateSchedule(tt.schedule)

			if stateCode != tt.expectedCode {
				t.Errorf("got %v, want %v", stateCode, tt.expectedCode)
			}
		})
	}
}

func evaluateSchedule(schedule MongoSchedule) StateCode {

	if len(schedule.Main) > 1 && len(schedule.Groups) > 1 {
		return RED
	}

	if schedule.Metrics != nil {
		for _, load := range schedule.Metrics.TeacherLoads {
			if load > 10 {
				return ORANGE
			}
		}
	}

	if schedule.Metrics != nil {
		for _, load := range schedule.Metrics.TeacherLoads {
			if load > 5 {
				return YELLOW
			}
		}
	}

	return OK
}
