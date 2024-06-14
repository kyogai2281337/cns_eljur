package set

import (
	"errors"
	"fmt"
)

var (
	// errFail = errors.New("failed to do")
	errVal = errors.New("there is no items with this name")
)

// Push adds an item to the Set.
//
// The item parameter is the value to be added to the Set.
// It should be of any type.
//
// The function returns an error if there is any issue adding the item to the Set.
// It returns nil if the item is successfully added to the Set.
func New() *Set {
	return &Set{Set: make(map[interface{}]struct{})}
}
func (s *Set) Push(item interface{}) error {
	if s.Set == nil {
		s.Set = make(map[interface{}]struct{})
	}
	s.Set[item] = struct{}{}
	return nil
}

// Remove removes an item from the Set.
//
// The item parameter is the value to be removed from the Set.
// It should be of any type.
//
// The function returns the removed item and an error if there is any issue removing the item from the Set.
// It returns nil and nil if the item is successfully removed from the Set.
func (s *Set) Remove(item any) (interface{}, error) {
	_, ok := s.Set[item]
	if !ok {
		return nil, errVal
	}
	delete(s.Set, item)
	return item, nil
}

// Union takes two sets as input and returns a new set that contains all the elements from both sets.
//
// Parameters:
// - s1: a pointer to a Set object representing the first set.
// - s2: a pointer to a Set object representing the second set.
//
// Returns:
// - a pointer to a Set object representing the union of the two sets.
// - an error if there was an issue creating the new set.
func Union(s1, s2 *Set) (*Set, error) {
	resp := &Set{make(map[interface{}]struct{})}
	for k := range s1.Set {
		resp.Push(k)
	}
	for k := range s2.Set {
		resp.Push(k)
	}
	return resp, nil
}

// Out prints all the elements of the Set to the console, separated by semicolons.
//
// No parameters.
// No return values.
func (s *Set) Out() {
	for k := range s.Set {
		fmt.Print(k, "; ")
	}
	fmt.Print("\n")
}

// Intersection takes two sets as input and returns a new set that contains all the elements from both sets.
//
// Parameters:
// - s1: a pointer to a Set object representing the first set.
// - s2: a pointer to a Set object representing the second set.
//
// Returns:
// - a pointer to a Set object representing the intersection of the two sets.
// - an error if there was an issue creating the new set.
func Intersection(s1, s2 *Set) (*Set, error) {
	resp := &Set{}
	for k := range s1.Set {
		_, ok := s2.Set[k]
		if ok {
			resp.Push(k)
		}
	}
	return resp, nil
}

func (s *Set) getRandItem() any {
	for i := range s.Set {
		return i
	}
	return nil
}
