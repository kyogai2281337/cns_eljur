package constructor_logic_entrypoint

import "encoding/json"

type DirType uint8

const (
	_ DirType = iota
	DirInsert
	DirDelete
	DirTX
)

type Directive struct {
	Type       DirType      `json:"type"`
	ID         string       `json:"id"` // correlation_id
	ScheduleID string       `json:"schedule_id"`
	Data       interface{}  `json:"data"`
	Resp       chan DirResp `json:"-"`
}

type DirResp struct {
	Data interface{} `json:"data"`
	Err  error       `json:"error"`
}

func (dir *Directive) Marshal() ([]byte, error) { return json.Marshal(dir) }
func (rsp *DirResp) Marshal() ([]byte, error)   { return json.Marshal(rsp) }

func UnmarshalDirective(data []byte) (Directive, error) {
	var dir Directive
	err := json.Unmarshal(data, &dir)
	return dir, err
}

// for internal interfaces, as a conv
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

type UpdateTXRequest struct {
	Data []Directive `json:"data"`
}
