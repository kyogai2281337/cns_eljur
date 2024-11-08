package methods

// MongoSchedule представляет расписание из MongoDB
type MongoSchedule struct {
	Name     string   `json:"name"`
	Groups   []string `json:"groups"`
	Teachers []string `json:"teachers"`
	Cabinets []string `json:"cabinets"`
	Plans    []string `json:"plans"`
	Days     int      `json:"days"`
	Pairs    int      `json:"pairs"`
	//Metrics                   *MongoMetrics       `json:"metrics"`
	Main                      [][][]*MongoLecture `json:"schedule"`
	MaxGroupLecturesFor2Weeks int                 `json:"weeklimit"`
	MaxGroupLecturesForDay    int                 `json:"daylimit"`
}

// MongoLecture представляет лекцию
type MongoLecture struct {
	Cabinet string   `json:"cabinet"`
	Teacher string   `json:"teacher"`
	Groups  []string `json:"group"`
	Subject string   `json:"subject"`
}

// Дополнительные структуры, такие как MongoMetrics, могут быть определены по необходимости.
