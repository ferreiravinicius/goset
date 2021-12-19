package gset

import (
	"fmt"
	"strings"
)

type hashset map[interface{}]struct{}

func HashSet(items ...interface{}) Set {
	s := make(hashset)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func (s hashset) Add(item interface{}) bool {
	if _, exists := s[item]; exists {
		return false
	}
	s[item] = struct{}{}
	return true
}

func (s hashset) String() string {
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

func (s hashset) Collect() []interface{} {
	result := make([]interface{}, len(s))
	index := 0
	for item := range s {
		result[index] = item
		index++
	}
	return result
}
