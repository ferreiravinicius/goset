package gset_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/ferreiravinicius/gset"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	s := gset.NewSet()
	assert.True(t, s.Add("t"))
	assert.Len(t, s, 1)
	assert.Contains(t, s, "t")
	assert.False(t, s.Add("t"))
}

func TestString(t *testing.T) {
	assert.Equal(t, "Set{}", fmt.Sprint(gset.NewSet()))
	assert.Equal(t, "Set{1}", fmt.Sprint(gset.NewSet(1)))
	assert.Equal(t, "Set{1 2}", fmt.Sprint(gset.NewSet(1, 2)))
}

func TestSlice(t *testing.T) {
	s := gset.NewSet(1, 1, 2, 2, 3, 3)
	assert.Len(t, s.Slice(), 3)
	assert.Contains(t, s.Slice(), 1)
	assert.Len(t, gset.NewSet(), 0)
}

const sz = 20_000

func randomize() []int {
	rand.Seed(time.Now().UnixNano())
	list := make([]int, sz)
	for n := 0; n < sz; n++ {
		list[n] = rand.Intn(sz / 2)
	}
	return list
}

var (
	inputs = randomize()
)

func BenchmarkSimple(b *testing.B) {
	m := make(map[int]struct{})
	for n := 0; n < b.N; n++ {
		v := inputs[n%sz]
		m[v] = struct{}{}
	}
}

func BenchmarkSimpleWithVerify(b *testing.B) {
	m := make(map[int]struct{})
	// b.Logf("num %v", b.N)
	for n := 0; n < b.N; n++ {
		if _, ok := m[n]; !ok {
			v := inputs[n%sz]
			m[v] = struct{}{}
		}
	}
}
func BenchmarkSlice(b *testing.B) {
	m := make(map[int]struct{})
	s := make([]int, 100)
	// b.Logf("num %v", b.N)
	for n := 0; n < b.N; n++ {
		if _, ok := m[n]; !ok {
			v := inputs[n%sz]
			m[v] = struct{}{}
			s = append(s, n)
		}
	}
	_ = s
}
