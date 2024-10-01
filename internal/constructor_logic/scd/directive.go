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
	Type DirType      `json:"type"`
	ID   string       `json:"id"` // correlation_id
	Data interface{}  `json:"data"`
	Resp chan DirResp `json:"-"`
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
