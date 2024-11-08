package methods

type MongoSchedule struct {
	Name                      string              `bson:"name" json:"name"`
	Groups                    []string            `json:"groups"`
	Teachers                  []string            `json:"teachers"`
	Cabinets                  []string            `json:"cabinets"`
	Plans                     []string            `json:"plans"`
	Days                      int                 `json:"days"`
	Pairs                     int                 `json:"pairs"`
	Metrics                   *MongoMetrics       `json:"metrics"`
	Main                      [][][]*MongoLecture `json:"schedule"`
	MaxGroupLecturesFor2Weeks int                 `json:"weeklimit"`
	MaxGroupLecturesForDay    int                 `json:"daylimit"`
}

type MongoWindows struct {
	Groups   map[string][]int `json:"groups"`
	Teachers map[string][]int `json:"teachers"`
}

type MongoMetrics struct {
	Plans        map[string]map[string]int `json:"plans"`
	Wins         *MongoWindows             `json:"windows"`
	TeacherLoads map[string]int            `json:"teacher_loads"`
}

type MongoLecture struct {
	Cabinet string   `json:"cabinet"`
	Teacher string   `json:"teacher"`
	Groups  []string `json:"group"`
	Subject string   `json:"subject"`
}

type StateCode string

type StateInst struct {
	Key   string `json:"key"`   // For example: "Insufficient pairs"
	Value string `json:"value"` // For example: "Group A // Teacher Vasya"
}

const (
	OK     StateCode = "OK"
	RED    StateCode = "RED"
	ORANGE StateCode = "ORANGE"
	YELLOW StateCode = "YELLOW"
)

// Reviewscomms is the output of the analysis
type Reviewscomms map[StateCode][]StateInst
