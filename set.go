package gset

import (
	"fmt"
	"strings"
)

type set map[interface{}]struct{}

func NewSet(items ...interface{}) Set {
	s := make(set)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func (s set) Add(item interface{}) bool {
	if _, exists := s[item]; exists {
		return false
	}
	s[item] = struct{}{}
	return true
}

func (s set) String() string {
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

func (s set) Slice() []interface{} {
	result := make([]interface{}, len(s))
	index := 0
	for item := range s {
		result[index] = item
		index++
	}
	return result
}
