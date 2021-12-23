package hashset

import (
	"fmt"
	"strings"
)

type HashSet struct {
	// inside struct to prevent external modification
	hashMap map[interface{}]struct{}
}

func New() *HashSet {
	return &HashSet{hashMap: make(map[interface{}]struct{})}
}

func FromSlice(slice []interface{}) *HashSet {
	return From(slice...)
}

func From(items ...interface{}) *HashSet {
	s := New()
	for _, item := range items {
		s.Add(item)
	}
	return s
}

func (s HashSet) Add(item interface{}) bool {
	if _, exists := s.hashMap[item]; exists {
		return false
	}
	s.hashMap[item] = struct{}{}
	return true
}

func (s HashSet) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "Set{")
	first := true
	for item := range s.hashMap {
		template := " %v"
		if first {
			template = "%v"
		}
		fmt.Fprintf(&sb, template, item)
		first = false
	}
	fmt.Fprint(&sb, "}")
	return sb.String()
}

func (s HashSet) ToSlice() []interface{} {
	result := make([]interface{}, len(s.hashMap))
	index := 0
	for item := range s.hashMap {
		result[index] = item
		index++
	}
	return result
}

func (s HashSet) Contains(item interface{}) bool {
	_, exists := s.hashMap[item]
	return exists
}

func (s HashSet) Remove(item interface{}) bool {
	beforeSize := len(s.hashMap)
	delete(s.hashMap, item)
	afterSize := len(s.hashMap)
	return beforeSize > afterSize
}

func (s HashSet) ContainsAll(items ...interface{}) bool {
	for _, item := range items {
		if _, exists := s.hashMap[item]; !exists {
			return false
		}
	}
	return true
}

func (s HashSet) Len() int {
	return len(s.hashMap)
}
