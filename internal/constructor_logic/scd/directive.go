package constructor_logic_entrypoint

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

type DirType uint8

const (
	_ DirType = iota
	DirInsert
	DirDelete
	DirTX
	DirRename
)

type Directive struct {
	Type       DirType      `json:"type"`
	ID         string       `json:"id"` // correlation_id
	ScheduleID string       `json:"schedule_id"`
	Data       interface{}  `json:"data"`
	Resp       chan DirResp `json:"-"`
}

type DirResp struct {
	Data string `json:"data"`
	Err  error  `json:"error"`
}

func NewErrorResp(err error) *DirResp {
	return &DirResp{
		Err: err,
	}
}

func (dir *Directive) Marshal() ([]byte, error) { return json.Marshal(dir) }

// ! Fixed specsyms there, using regex:
// * Were: "instruction 5: cabinet already exists at 0: Cabinet: 201, Teacher: Petr Ivanov, Subject: С++\n\t\t\tGroup: 201IT"
//* Now: "instruction 5: cabinet already exists at 0: Cabinet: 201, Teacher: Petr Ivanov, Subject: С++ Group: 201IT"

func (rsp *DirResp) Marshal() ([]byte, error) {
	if rsp.Err != nil {
		re := regexp.MustCompile(`[\n\t]+`)
		errStr := rsp.Err.Error()
		cleanedError := re.ReplaceAllString(errStr, " ") // заменяем символы новой строки и табуляции на пробел
		rsp.Err = errors.New(cleanedError)

		return json.Marshal(map[string]interface{}{
			"data":  rsp.Data,
			"error": rsp.Err.Error(),
		})
	} else {
		return json.Marshal(map[string]interface{}{
			"data":  rsp.Data,
			"error": "",
		})
	}
}

// func UnmarshalDirective(data []byte) (Directive, error) {
// 	var dir Directive
// 	err := json.Unmarshal(data, &dir)
// 	return dir, err
// }

func UnmarshalDirective(data []byte) (*Directive, error) {
	var directiveMap map[string]interface{}

	// Десериализация входящих данных в общий формат
	if err := json.Unmarshal(data, &directiveMap); err != nil {
		return nil, err
	}

	var directive Directive

	// Получаем тип директивы
	dirType, ok := directiveMap["type"].(float64)
	if !ok {
		return nil, fmt.Errorf("missing or invalid directive type")
	}
	t := DirType(dirType)

	// Преобразование типа директивы в соответствующую структуру
	switch t {
	case DirInsert:
		// Преобразуем директиву вставки
		var insertReq UpdateInsertRequest
		insertData, err := json.Marshal(directiveMap)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal insert request: %v", err)
		}
		if err := json.Unmarshal(insertData, &insertReq); err != nil {
			return nil, fmt.Errorf("failed to unmarshal insert request: %v", err)
		}
		directive = Directive{
			Type: DirInsert,
			Data: insertReq,
		}
	case DirRename:
		var renameReq UpdateRenameRequest
		renameData, err := json.Marshal(directiveMap)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal rename request: %v", err)
		}
		if err := json.Unmarshal(renameData, &renameReq); err != nil {
			return nil, fmt.Errorf("failed to unmarshal rename request: %v", err)
		}
		directive = Directive{
			Type: DirRename,
			Data: renameReq,
		}
	case DirDelete:
		// Преобразуем директиву удаления
		var deleteReq UpdateDeleteRequest
		deleteData, err := json.Marshal(directiveMap)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal delete request: %v", err)
		}
		if err := json.Unmarshal(deleteData, &deleteReq); err != nil {
			return nil, fmt.Errorf("failed to unmarshal delete request: %v", err)
		}
		directive = Directive{
			Type: DirDelete,
			Data: deleteReq,
		}

	case DirTX:
		// Преобразуем директиву транзакции
		var txReq UpdateTXRequest

		// Достаем массив вложенных директив
		rawData, ok := directiveMap["data"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("failed to parse transaction data")
		}
		corrId, ok := directiveMap["id"].(string)
		if !ok {
			return nil, fmt.Errorf("failed to parse correlation_id")
		}
		scdId, ok := directiveMap["schedule_id"].(string)
		if !ok {
			return nil, fmt.Errorf("failed to parse schedule_id")
		}

		// Рекурсивно обрабатываем каждую вложенную директиву
		for _, rawDirective := range rawData {
			// Сериализуем вложенную директиву обратно в JSON
			rawDirectiveData, err := json.Marshal(rawDirective)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal nested directive: %v", err)
			}

			// Рекурсивно вызываем UnmarshalDirective для каждой вложенной директивы
			nestedDirective, err := UnmarshalDirective(rawDirectiveData)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal nested directive: %v", err)
			}

			// Добавляем результат в список директив транзакции
			txReq.Data = append(txReq.Data, *nestedDirective)
		}

		directive = Directive{
			Type:       DirTX,
			ID:         corrId,
			ScheduleID: scdId,
			Data:       txReq,
		}

	default:
		return nil, fmt.Errorf("unexpected directive type: %v", dirType)
	}

	return &directive, nil
}

func UnmarshalDirResp(data []byte) (*DirResp, error) {
	var respMap map[string]interface{}
	if err := json.Unmarshal(data, &respMap); err != nil {
		return nil, err
	}
	fmt.Println(respMap)
	id, ok := respMap["data"].(string)
	if !ok {
		return nil, errors.New("empty data field")
	}
	resp := new(DirResp)
	resp.Data = id
	errorString, ok := respMap["error"].(string)
	if ok {
		resp.Err = errors.New(errorString)
	}

	return resp, nil
}

// for internal interfaces, as a conv
type UpdateInsertRequest struct {
	Type DirType `json:"type"`
	Data struct {
		Day     uint `json:"day"`
		Pair    uint `json:"pair"`
		Lecture struct {
			Groups  []string `json:"groups"`
			Teacher string   `json:"teacher"`
			Cabinet string   `json:"cabinet"`
			Subject string   `json:"subject"`
		} `json:"lecture"`
	} `json:"data"`
}

type UpdateDeleteRequest struct {
	Type DirType `json:"type"`
	Data struct {
		Day  uint   `json:"day"`
		Pair uint   `json:"pair"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"data"`
}

type UpdateTXRequest struct {
	ID         string      `json:"id"`
	ScheduleID string      `json:"schedule_id"`
	Data       []Directive `json:"data"`
}

type UpdateRenameRequest struct {
	Type DirType `json:"type"`
	Data struct {
		Name string `json:"name"`
	} `json:"data"`
}
