package constructor

import "github.com/kyogai2281337/cns_eljur/pkg/sql/model"

/*
	Хочется сказать, что этот модуль в принципе введен для того,
	чтобы декомпозировать логику из конструктора, так как 600 строчек кода,
	которые там находятся начали напоминать результат алгоритма ed25519.

	Такс, здесть в принципе опись всей логики, чтобы не забыть). Итак:

	1) Инициализация пустой структуры внутри функции, не метода;

	2) Заполнение этой пустой структуры в той же функции Flow-ами;

	3) По механике сказать не могу, но в голове крутится просто для
	   предмета так же добавить из вводных Capacity и от них уже плясать, так как
	   не очень будет понятно, с чего и как вообще заполнять кабинеты с разной вместимостью.
	   А если подетальнее, то FlowHolder будет хранить в себе массивчик Flow'ов,
	   каждый из которых содержит связь: Teacher->Subject->[]Group и при каждой итерации конструктора
	   он будет глядеть, свободны ли те-то группы при текущей итерации препа,
	   и надо ли нам вообще эти предметы еще кидать туда или нет, но как говорится, данный
	   Storage не берет на себя ответственность, иначе это не код а параша).

	4) После заполнения необходимо, как выше сказано, интегрировать при генерации
	   данный массивчикб причем по best-practice чтобы не мучать себя все потоки при генерации
	   сортировать в конструкторе сначана, или генерить сначала, а потом распределять по плейсу
	   рандомно.

*/

// Тип связи : T->S->[]G
type Flow struct {
	Teacher *model.Teacher
	Groups  []*model.Group
	Subject *model.Subject
}

// FlowHolder - содержит все Flow'ы
type FlowHolder struct {
	Flows    []*Flow                     // Пайка в памяти для того, чтобы не мучаться перебросом групп с места на место
	_canonic map[int64]map[int64][]int64 // хочется от удаления потворить, поэтому сделал эталон, который буду копировать
	buffer   map[int64]map[int64][]int64 // teacher_idx->map[sub.id]->arr_group_id, чтобы было весело
}

func NewFlowHolder() *FlowHolder {
	return &FlowHolder{
		Flows:    make([]*Flow, 0),
		buffer:   make(map[int64]map[int64][]int64, 0),
		_canonic: make(map[int64]map[int64][]int64, 0),
	}
}

func (f *FlowHolder) addFlow(flow *Flow) {
	f.Flows = append(f.Flows, flow)
}

// Для правильного использования рекомендую тыкать в начале генератора, чтобы нормально можно было проиндексировать
func (f *FlowHolder) InitBuf(teachers []*model.Teacher) {
	for _, teacher := range teachers {
		for g, ss := range teacher.Links {
			for _, s := range ss {

				if s.RecommendCabType == model.Flowable || s.RecommendCabType == model.Sport {
					// Нужно откуда-то рожать, каким образом и по какому capacity мне распределять группы и какие кабы забирать
					// Задать вопрос, как они распределяют группы по кабинетам, как им это надо предоставить...
				}
			}

			// Костыль
			mock(g)
		}
		sl := teacher.ShorterLinks()
		f._canonic[teacher.ID] = sl
		f.buffer[teacher.ID] = sl

	}
}

func (f *FlowHolder) Refresh() {
	f.buffer = f._canonic
}

// PopBuf - возвращает Flow, если он есть в буфере, или nil, если его нет
//
// Возвращает nil, если не нашел, а иначе Flow, который содержит группу g и
// предмет s
func (f *FlowHolder) PopBuf(g *model.Group, s *model.Subject) *Flow {
	for _, flow := range f.Flows {
		if flow.Subject == s {
			for _, gr := range flow.Groups {
				if gr == g {
					return flow
				}
			}
		}
	}
	return nil
}

func mock(items ...any) {
	_ = items[:0]
}
