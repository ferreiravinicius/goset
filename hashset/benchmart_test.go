package hashset

import "testing"

const size = 100_000

func generateHashSet() *HashSet {
	set := New(size)
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	return set
}

var sut = generateHashSet()

func IgnoreBenchmarkIterateForEach(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sut.ForEach(func(item interface{}) {
			num := item.(int)
			_ = num
		})
	}
}

func IgnoreBenchmarkIterateNormal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for key := range sut.hashMap {
			num := key.(int)
			_ = num
		}
	}
}

func deleteUsingExists(s *HashSet, v interface{}) bool {
	if _, exists := s.hashMap[v]; exists {
		delete(s.hashMap, v)
		return true
	}
	return false
}

func deleteUsingLen(s *HashSet, v interface{}) bool {
	b := len(s.hashMap)
	delete(s.hashMap, v)
	return b > len(s.hashMap)
}

func BenchmarkDeleteExists(b *testing.B) {
	s := generateHashSet()
	tmp := false
	for n := 0; n < b.N; n++ {
		which := n % size
		tmp = deleteUsingExists(s, which)
	}
	_ = tmp
}

func BenchmarkDeleteLen(b *testing.B) {
	s := generateHashSet()
	tmp := false
	for n := 0; n < b.N; n++ {
		which := n % size
		tmp = deleteUsingLen(s, which)
	}
	_ = tmp
}
