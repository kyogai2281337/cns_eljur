package utils

import (
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func ConvertToSL(data bson.M) (map[int64][]int64, error) {
	sl := make(map[int64][]int64)
	v, ok := data["links"]
	if !ok {
		return nil, fmt.Errorf("key 'links' not found in data")
	}
	links, ok := v.(bson.M)
	if !ok {
		return nil, fmt.Errorf("expected bson.M for key 'links' but got %T", v)
	}
	for k, v := range links {
		// Преобразование ключа в int64, если это возможно
		key, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			// Пропускаем ключи, которые не являются числами
			fmt.Printf("Skipping non-numeric key: %s\n", k)
			continue
		}

		// Преобразование значения в []int64
		var values []int64
		arr, ok := v.(bson.A)
		if !ok {
			return nil, fmt.Errorf("expected bson.A for key %s but got %T", k, v)
		}
		for _, item := range arr {
			// Здесь необходимо учитывать, что значения могут быть в различных форматах
			// Например, они могут быть float64 (так как MongoDB использует float64 для всех чисел)
			switch val := item.(type) {
			case int64:
				values = append(values, val)
			case int32:
				values = append(values, int64(val))
			case float64:
				values = append(values, int64(val))
			default:
				return nil, fmt.Errorf("unexpected type for value: %T", item)
			}
		}
		sl[key] = values
	}
	return sl, nil
}

func equalSlices(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualMaps(map1, map2 map[int64][]int64) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, val1 := range map1 {
		val2, ok := map2[key]
		if !ok {
			return false
		}
		if !equalSlices(val1, val2) {
			return false
		}
	}

	return true
}

func EqualEasyMaps(a, b map[int64]int) bool {

	if len(a) != len(b) {
		return false
	}
	for k, v := range a {

		val, ok := b[k]
		if !ok {
			return false
		}
		if val != v {
			return false
		}
	}
	return true

}

func ConvertToPlan(data bson.M) (map[int64]int, error) {
	plan := make(map[int64]int)
	v, ok := data["plans"]
	if !ok {
		return nil, fmt.Errorf("key 'plans' not found in data")
	}
	plans, ok := v.(bson.M)
	if !ok {
		return nil, fmt.Errorf("expected bson.M for key 'plans' but got %T", v)
	}
	for k, v := range plans {
		// Преобразование ключа в int64, если это возможно
		key, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			// Пропускаем ключи, которые не являются числами
			fmt.Printf("Skipping non-numeric key: %s\n", k)
			continue
		}

		// Преобразование значения в int

		switch value := v.(type) {
		case int64:
			plan[key] = int(value)
		case int32:
			plan[key] = int(value)
		case float64:
			plan[key] = int(value)
		default:
			return nil, fmt.Errorf("unexpected type for value: %T", v)
		}
	}

	return plan, nil
}
