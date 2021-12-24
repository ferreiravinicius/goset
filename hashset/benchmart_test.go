package hashset

import "testing"

const size = 100_000

func generateHashSet() *HashSet {
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	return set
}

var sut = generateHashSet()

func BenchmarkIterateForEach(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sut.ForEach(func(item interface{}) {
			num := item.(int)
			_ = num
		})
	}
}
func BenchmarkIterateNormal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for key := range sut.hashMap {
			num := key.(int)
			_ = num
		}
	}
}
