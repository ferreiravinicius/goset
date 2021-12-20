package goset

import (
	"fmt"
	"strings"
)

type hashSet map[interface{}]struct{}

// Implemention of set backed by a hash table (go builtin map).
// Offers constant time performance for the basic operations.
// This implementation is not syncronized.
func HashSet(items ...interface{}) hashSet {
	s := make(hashSet)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func (s hashSet) Add(item interface{}) bool {
	if _, exists := s[item]; exists {
		return false
	}
	s[item] = struct{}{}
	return true
}

func (s hashSet) String() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "Set{")
	first := true
	for item := range s {
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
	result := make([]interface{}, len(s))
	index := 0
	for item := range s {
		result[index] = item
		index++
	}
	return result
}

func (s hashSet) Contains(item interface{}) bool {
	_, exists := s[item]
	return exists
}

func (s hashSet) Remove(item interface{}) bool {
	beforeSize := len(s)
	delete(s, item)
	afterSize := len(s)
	return beforeSize > afterSize
}

func (s hashSet) ContainsAll(items ...interface{}) bool {
	for _, item := range items {
		if _, exists := s[item]; !exists {
			return false
		}
	}
	return true
}

func (s hashSet) Len() int {
	return len(s)
}
