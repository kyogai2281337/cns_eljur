package structures

import constructor_logic_entrypoint "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/scd"

type CreateLimits struct {
	MaxWeeks int `json:"max_weeks"`
	MaxDays  int `json:"max_days"`
	Days     int `json:"days"`
	Pairs    int `json:"pairs"`
}

type CreateRequest struct {
	Name     string        `json:"name"`
	Limits   *CreateLimits `json:"limits"`
	Groups   []int64       `json:"groups"`
	Plans    []int64       `json:"plans"`
	Cabinets []int64       `json:"cabinets"`
	Teachers []int64       `json:"teachers"`
}

type GetByIDRequest struct {
	ID string `json:"id"`
}

type UpdateRequest struct {
	ID     string `json:"id"`
	Values []any  `json:"values"` // Array of directive/s, see internal\constructor_logic\scd\directive.go
}

type UpdateInsertRequest struct {
	Day     int `json:"day"`
	Pair    int `json:"pair"`
	Lecture struct {
		Groups  []string `json:"group"`
		Teacher string   `json:"teacher"`
		Cabinet string   `json:"cabinet"`
		Subject string   `json:"subject"`
	} `json:"lecture"`
}

type UpdateDeleteRequest struct {
	Day  int    `json:"day"`
	Pair int    `json:"pair"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type RenameRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SaveXLSXRequest struct {
	ID string `json:"id"`
}

//	DirInsert = 1
//	DirDelete = 2
type Directive struct {
	Type constructor_logic_entrypoint.DirType `json:"type"`
	//ID         string                                    `json:"id"` // correlation_id
	ScheduleID string      `json:"id"`
	Data       interface{} `json:"data"`
	// Resp       chan constructor_logic_entrypoint.DirResp `json:"-"`
}
