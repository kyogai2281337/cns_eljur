package model

import (
	"encoding/json"
	"fmt"
)

type CabType int

func (c CabType) String() string {
	switch c {
	case Normal:
		return "Normal"
	case Flowable:
		return "Flowable"
	case Laboratory:
		return "Laboratory"
	case Computered:
		return "Computered"
	case Sport:
		return "Sport"
	}
	return "Unknown"
}

const (
	Normal CabType = iota
	Flowable
	Laboratory
	Computered
	Sport
)

func (c CabType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *CabType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case "Normal":
		*c = Normal
	case "Flowable":
		*c = Flowable
	case "Laboratory":
		*c = Laboratory
	case "Computered":
		*c = Computered
	case "Sport":
		*c = Sport
	default:
		return fmt.Errorf("unknown CabType: %s", s)
	}

	return nil
}
