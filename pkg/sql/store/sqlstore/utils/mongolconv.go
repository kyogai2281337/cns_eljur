package utils

import (
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func ConvertToSL(data bson.M) (map[int64][]int64, error) {
	sl := make(map[int64][]int64)
	for k, v := range data {
		// Преобразование ключа в int64, если это возможно
		key, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			// Пропускаем ключи, которые не являются числами
			return nil, fmt.Errorf("Skipping non-numeric key: %s", k)
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
