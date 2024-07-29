package constructor

import (
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"golang.org/x/exp/rand"
)

func ShuffleCabArray(arr []*model.Cabinet) []*model.Cabinet {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func ShuffleSpeArray(arr []*model.Specialization) []*model.Specialization {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func ShuffleGroupArray(arr []*model.Group) []*model.Group {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func ShuffleTeachArray(arr []*model.Teacher) []*model.Teacher {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
