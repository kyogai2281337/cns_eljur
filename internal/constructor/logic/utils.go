package constructor

import (
	"golang.org/x/exp/rand"
)

func _ShuffleArr[T any](arr []T) []T {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

// func ShuffleCabArray(arr []*model.Cabinet) []*model.Cabinet {
// 	for i := len(arr) - 1; i > 0; i-- {
// 		j := rand.Intn(i + 1)
// 		arr[i], arr[j] = arr[j], arr[i]
// 	}

// 	return arr
// }

// func ShuffleSpeArray(arr []*model.Specialization) []*model.Specialization {
// 	for i := len(arr) - 1; i > 0; i-- {
// 		j := rand.Intn(i + 1)
// 		arr[i], arr[j] = arr[j], arr[i]
// 	}

// 	return arr
// }

// func ShuffleGroupArray(arr []*model.Group) []*model.Group {
// 	for i := len(arr) - 1; i > 0; i-- {
// 		j := rand.Intn(i + 1)
// 		arr[i], arr[j] = arr[j], arr[i]
// 	}

// 	return arr
// }

// func ShuffleTeachArray(arr []*model.Teacher) []*model.Teacher {
// 	for i := len(arr) - 1; i > 0; i-- {
// 		j := rand.Intn(i + 1)
// 		arr[i], arr[j] = arr[j], arr[i]
// 	}

// 	return arr
// }
