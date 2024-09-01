package constructor

import (
	"fmt"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

type GroupHeap []*model.Group

type Lecture struct {
	Cabinet *model.Cabinet `json:"cabinet"`
	Teacher *model.Teacher `json:"teacher"`

	Groups  []*model.Group `json:"group"`
	Subject *model.Subject `json:"subject"`
}

func (l *Lecture) String() string {
	response := ""

	// Проверка на nil перед использованием значений
	if l.Cabinet != nil && l.Cabinet.Name != "" {
		response += fmt.Sprintf("Cabinet: %s", l.Cabinet.Name)
	} else {
		response += "Cabinet: N/A"
	}

	if l.Teacher != nil && l.Teacher.Name != "" {
		response += fmt.Sprintf(", Teacher: %s", l.Teacher.Name)
	} else {
		response += ", Teacher: N/A"
	}

	if l.Subject != nil && l.Subject.Name != "" {
		response += fmt.Sprintf(", Subject: %s", l.Subject.Name)
	} else {
		response += ", Subject: N/A"
	}

	// Добавляем группы, если они есть
	if l.Groups != nil {
		for _, g := range l.Groups {
			if g != nil && g.Name != "" {
				response += fmt.Sprintf("\n\t\t\tGroup: %s", g.Name)
			} else {
				response += "\n\t\t\tGroup: N/A"
			}
		}
	}

	return response
}

func (h GroupHeap) Len() int           { return len(h) }
func (h GroupHeap) Less(i, j int) bool { return h[i].Priority() < h[j].Priority() }
func (h GroupHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GroupHeap) Push(x interface{}) {
	*h = append(*h, x.(*model.Group))
}

func (h *GroupHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
func (h *GroupHeap) Peek() *model.Group {
	return (*h)[len(*h)-1]
}

func (h *GroupHeap) Find(g *model.Group) int {
	for i, v := range *h {
		if v == g {
			return i
		}
	}
	return -1
}

// Возвращение приоритета по количеству оставшихся пар
