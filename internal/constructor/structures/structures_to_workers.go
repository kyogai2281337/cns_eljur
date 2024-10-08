package structures

import (
	"encoding/json"
	"fmt"

	constructor_logic_entrypoint "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/scd"
)

/*
	! Важно !
	? Данный файл был создан для того, чтобы не разбегались глаза при вчитывании в 500 однотипных обьявлений структур
	? Конкретно здесь расположен маппинг для Update и Rename
*/

type UpdateRequest struct {
	ID     string      `json:"id"`
	Values []UpdateDir `json:"data"` // Array of directive/s, see internal\constructor_logic\scd\directive.go
}

type UpdateInsertRequest struct {
	Type constructor_logic_entrypoint.DirType `json:"type"`
	Data InsertReqData                        `json:"data"`
}

type InsertReqData struct {
	Day     int `json:"day"`
	Pair    int `json:"pair"`
	Lecture struct {
		Groups  []string `json:"groups"`
		Teacher string   `json:"teacher"`
		Cabinet string   `json:"cabinet"`
		Subject string   `json:"subject"`
	} `json:"lecture"`
}

func (s *UpdateInsertRequest) TypeName() string { return "insert" }

type UpdateDeleteRequest struct {
	Type constructor_logic_entrypoint.DirType `json:"type"`
	Data DeleteReqData                        `json:"data"`
}

type DeleteReqData struct {
	Day  int    `json:"day"`
	Pair int    `json:"pair"`
	Type string `json:"type"`
	Name string `json:"name"`
}

func (s *UpdateDeleteRequest) TypeName() string { return "delete" }

type UpdateRenameRequest struct {
	Type constructor_logic_entrypoint.DirType `json:"type"`
	Data RenameReqData                        `json:"data"`
}

func (s *UpdateRenameRequest) TypeName() string { return "rename" }

type RenameReqData struct {
	Name string `json:"name"`
}

// ----------------------------------------------

// * DirInsert = 1
// * DirDelete = 2
// * DirTX = 3
// * DirRename = 4
type Directive struct {
	Type       constructor_logic_entrypoint.DirType `json:"type"`
	ScheduleID string                               `json:"id"`
	Data       interface{}                          `json:"data"`
}

// ? Update Interfaces
// * UpdateDir
// * Uses for internal interfaces
// * Provides to map directives without any complicative handwork
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
		case constructor_logic_entrypoint.DirRename:
			var renameReq UpdateRenameRequest
			if err := json.Unmarshal(rawValue, &renameReq); err != nil {
				return err
			}
			renameReq.Type = constructor_logic_entrypoint.DirRename
			u.Values = append(u.Values, &renameReq)
		default:
			return fmt.Errorf("unknown directive type: %v", rawDirective.Type)
		}
	}
	return nil
}
