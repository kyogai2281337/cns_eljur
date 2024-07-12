package structures

type FindScheduleRequest struct {
	ObjType string `json:"obj_type"`
	ObjID   string `json:"obj_id"`
	Day     int    `json:"day"`
	Pair    int    `json:"pair"`
}

type Placement [2]int

type FindScheduleResponse struct {
	CabName     string    `json:"cab_name"`
	GroupName   string    `json:"group_name"`
	TeachName   string    `json:"teach_name"`
	SubjectName string    `json:"subject_name"`
	Place       Placement `json:"place"`
}
