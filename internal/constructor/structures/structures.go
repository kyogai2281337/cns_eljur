package structures

type CreateLimits struct {
	MaxWeeks int `json:"max_weeks"`
	MaxDays  int `json:"max_days"`
	Days     int `json:"days"`
	Pairs    int `json:"pairs"`
}

type CreateRequest struct {
	Limits   *CreateLimits `json:"limits"`
	Groups   []int64       `json:"groups"`
	Plans    []int64       `json:"plans"`
	Cabinets []int64       `json:"cabinets"`
	Teachers []int64       `json:"teachers"`
}
