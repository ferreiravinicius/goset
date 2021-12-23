package set

import (
	"fmt"
	"strings"
)

type hashSet struct {
	// inside struct to prevend external modification
	hashMap map[interface{}]struct{}
}

// Implemention of set backed by a HashTable (Go builtin Map).
// Offers constant time performance for the basic operations.
// This implementation is not synchronized.
func HashSet(items ...interface{}) *hashSet {
	s := &hashSet{hashMap: make(map[interface{}]struct{})}
	for _, item := range items {
		s.hashMap[item] = struct{}{}
	}
	return s
}

func (s hashSet) Add(item interface{}) bool {
	if _, exists := s.hashMap[item]; exists {
		return false
	}
	s.hashMap[item] = struct{}{}
	return true
}

func (s hashSet) String() string {
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

func (s hashSet) Collect() []interface{} {
	result := make([]interface{}, len(s.hashMap))
	index := 0
	for item := range s.hashMap {
		result[index] = item
		index++
	}
	return result
}

func (s hashSet) Contains(item interface{}) bool {
	_, exists := s.hashMap[item]
	return exists
}

func (s hashSet) Remove(item interface{}) bool {
	beforeSize := len(s.hashMap)
	delete(s.hashMap, item)
	afterSize := len(s.hashMap)
	return beforeSize > afterSize
}

func (s hashSet) ContainsAll(items ...interface{}) bool {
	for _, item := range items {
		if _, exists := s.hashMap[item]; !exists {
			return false
		}
	}
	return true
}

func (s hashSet) Len() int {
	return len(s.hashMap)
}
