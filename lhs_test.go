package goset

import "testing"

type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

type MapNode struct {
	hashMap map[Node]struct{}
}

type MapValue struct {
	hashMap map[interface{}]struct{}
}

func BenchmarkMapNode(b *testing.B) {
	sut := MapNode{
		hashMap: make(map[Node]struct{}),
	}
	var tmp *Node
	for n := 0; n < b.N; n++ {
		curr := Node{Value: n, next: tmp, prev: tmp}
		sut.hashMap[curr] = struct{}{}
		tmp = &curr
	}
}
func BenchmarkMapValue(b *testing.B) {
	sut := MapValue{
		hashMap: make(map[interface{}]struct{}),
	}
	for n := 0; n < b.N; n++ {
		sut.hashMap[n] = struct{}{}
	}
}
