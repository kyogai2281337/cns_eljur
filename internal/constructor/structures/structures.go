package structures

import (
	"encoding/json"
	"fmt"

	constructor_logic_entrypoint "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/scd"
)

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
	ID     string      `json:"id"`
	Values []UpdateDir `json:"data"` // Array of directive/s, see internal\constructor_logic\scd\directive.go
}

type UpdateInsertRequest struct {
	Type constructor_logic_entrypoint.DirType `json:"type"`
	Data struct {
		Day     int `json:"day"`
		Pair    int `json:"pair"`
		Lecture struct {
			Groups  []string `json:"groups"`
			Teacher string   `json:"teacher"`
			Cabinet string   `json:"cabinet"`
			Subject string   `json:"subject"`
		} `json:"lecture"`
	} `json:"data"`
}

func (s *UpdateInsertRequest) TypeName() string { return "insert" }

type UpdateDeleteRequest struct {
	Type constructor_logic_entrypoint.DirType `json:"type"`
	Data struct {
		Day  int    `json:"day"`
		Pair int    `json:"pair"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"data"`
}

func (s *UpdateDeleteRequest) TypeName() string { return "delete" }

type RenameRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SaveXLSXRequest struct {
	ID string `json:"id"`
}

// DirInsert = 1
// DirDelete = 2
type Directive struct {
	Type constructor_logic_entrypoint.DirType `json:"type"`
	//ID         string                                    `json:"id"` // correlation_id
	ScheduleID string      `json:"id"`
	Data       interface{} `json:"data"`
	// Resp       chan constructor_logic_entrypoint.DirResp `json:"-"`
}

// Update Interfaces
type UpdateDir interface {
	TypeName() string
}

func (u *UpdateRequest) UnmarshalJSON(data []byte) error {
	// Создаем временную структуру для работы с полями
	var rawValues struct {
		ID     string            `json:"id"`
		Values []json.RawMessage `json:"data"` // Обратите внимание, ключ совпадает с `json:"data"`
	}

	// Парсим общие данные
	if err := json.Unmarshal(data, &rawValues); err != nil {
		return err
	}
	u.ID = rawValues.ID

	// Обрабатываем каждый элемент массива `Values`
	for _, rawValue := range rawValues.Values {
		// Сначала определяем тип директивы
		var rawDirective struct {
			Type uint8 `json:"type"` // Ожидаем числовой тип для директивы
		}
		if err := json.Unmarshal(rawValue, &rawDirective); err != nil {
			return err
		}

		switch constructor_logic_entrypoint.DirType(rawDirective.Type) {
		case constructor_logic_entrypoint.DirInsert:
			var insertReq UpdateInsertRequest
			if err := json.Unmarshal(rawValue, &insertReq); err != nil {
				return err
			}
			insertReq.Type = constructor_logic_entrypoint.DirInsert
			u.Values = append(u.Values, &insertReq)
		case constructor_logic_entrypoint.DirDelete:
			var deleteReq UpdateDeleteRequest
			if err := json.Unmarshal(rawValue, &deleteReq); err != nil {
				return err
			}
			deleteReq.Type = constructor_logic_entrypoint.DirDelete
			u.Values = append(u.Values, &deleteReq)
		default:
			return fmt.Errorf("unknown directive type: %v", rawDirective.Type)
		}
	}
	return nil
}
