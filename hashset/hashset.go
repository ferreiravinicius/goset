package hashset

import (
	"fmt"
	"strings"
)

// Set is a collection that contains no duplicate elements.
// Implemention of set backed by a HashTable (Go builtin Map).
// Order of insertion is NOT guaranteed.
// This implementation is not synchronized.
//
// Example:
//
// s := hashset.New()
type HashSet struct {
	// inside struct to prevent external modification
	hashMap map[interface{}]struct{}
}

// Creates a new HashSet instance
func New() *HashSet {
	return &HashSet{hashMap: make(map[interface{}]struct{})}
}

// Creates a new HashSet containg all provided items.
// It can panics if you don't provide the variadic items in a correct way.
//
// Example:
//
// s := hashset.From(1, 2, 3) // correct
//
// s := hashset.From([]int{1, 2, 3}) // wrong
func From(items ...interface{}) *HashSet {
	s := New()
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// Adds a new item to this set if not already exists - O(1).
// Returns true if the provided item is added or false otherwise.
//
// Order of insertion is NOT guaranteed.
func (s HashSet) Add(item interface{}) bool {
	if _, exists := s.hashMap[item]; exists {
		return false
	}
	s.hashMap[item] = struct{}{}
	return true
}

// Returns a representation of this set as a string in the for `Set{1, 2, 3}` - O(n).
// Order of insertion is NOT guaranteed.
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

// Creates a slice containing all items - O(n).
func (s HashSet) ToSlice() []interface{} {
	result := make([]interface{}, len(s.hashMap))
	index := 0
	for item := range s.hashMap {
		result[index] = item
		index++
	}
	return result
}

// Returns true if provided item already exists in this set - O(1).
func (s HashSet) Contains(item interface{}) bool {
	_, exists := s.hashMap[item]
	return exists
}

// Removes the provided item from this set if it exists - O(1).
// Returns true if removed or false otherwise.
func (s HashSet) Remove(item interface{}) bool {
	beforeSize := len(s.hashMap)
	delete(s.hashMap, item)
	afterSize := len(s.hashMap)
	return beforeSize > afterSize
}

// Returns true if all provided items exists in this set.
func (s HashSet) ContainsAll(items ...interface{}) bool {
	for _, item := range items {
		if _, exists := s.hashMap[item]; !exists {
			return false
		}
	}
	return true
}

// Returns the size/lenght of this set.
func (s HashSet) Len() int {
	return len(s.hashMap)
}

// Function to perform an operation on each item whe used with ForEach
type Consumer func(item interface{})

// Performs an action for each item inside this set.
// Does not guarantee order.
func (s HashSet) ForEach(consumer Consumer) {
	for item := range s.hashMap {
		consumer(item)
	}
}
