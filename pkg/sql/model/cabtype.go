package model

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
