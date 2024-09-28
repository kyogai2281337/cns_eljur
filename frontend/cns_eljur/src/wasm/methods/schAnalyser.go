package methods

/*
Cases for rewieving:

	- 2 instances on 1 time (excludes flowable CABS, or subs) RED
	- overload of some subs from ppl`s participations (internal\constructor\logic\constructor.go, String, MakeRewiew, same structs, usage of fields) ORANGE
	- avg, daily overload YELLOW
*/

// Entrypoint
type MongoSchedule struct {
	Name string `bson:"name" json:"name"`
	//ID                        primitive.ObjectID  `bson:"_id" json:"-"`
	Groups                    []string            `bson:"groups" json:"groups"`
	Teachers                  []string            `bson:"teachers" json:"teachers"`
	Cabinets                  []string            `bson:"cabinets" json:"cabinets"`
	Plans                     []string            `bson:"plans" json:"plans"`
	Days                      int                 `bson:"days" json:"days"`
	Pairs                     int                 `bson:"pairs" json:"pairs"`
	Metrics                   *MongoMetrics       `bson:"metrics" json:"metrics"`
	Main                      [][][]*MongoLecture `bson:"schedule" json:"schedule"`
	MaxGroupLecturesFor2Weeks int                 `bson:"weeklimit" json:"weeklimit"`
	MaxGroupLecturesForDay    int                 `bson:"daylimit" json:"daylimit"`
}

type MongoWindows struct {
	Groups   map[string][]int `bson:"groups" json:"groups"`
	Teachers map[string][]int `bson:"teachers" json:"teachers"`
}
type MongoMetrics struct {
	Plans        map[string]map[string]int `bson:"plans" json:"plans"`
	Wins         *MongoWindows             `bson:"windows" json:"windows"`
	TeacherLoads map[string]int            `bson:"teacher_loads" json:"teacher_loads"`
}

type MongoLecture struct {
	Cabinet string   `bson:"cabinet" json:"cabinet"`
	Teacher string   `bson:"teacher" json:"teacher"`
	Groups  []string `bson:"group" json:"group"`
	Subject string   `bson:"subject" json:"subject"`
}

type StateCode string

type StateInst struct {
	Key   string `json:"key"`   // 4ex "Insufficient of pairs"
	Value string `json:"value"` // 4ex "Group penis // Teach Vasya"
}

const (
	OK     StateCode = "OK"
	RED    StateCode = "RED"
	ORANGE StateCode = "ORANGE"
	YELLOW StateCode = "YELLOW"
)

// stdout
type Rewiewscomms map[StateCode][]StateInst
