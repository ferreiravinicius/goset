package hashset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSetNew(t *testing.T) {
	assert.Empty(t, New().hashMap)
}

func TestHashSetFrom(t *testing.T) {
	assert.Contains(t, From(1).hashMap, 1)
	assert.Len(t, From(1, 2, 3).hashMap, 3)
	s := []interface{}{1, 2, 3}
	assert.Len(t, From(s...).hashMap, 3)
}

func TestHashSetAdd(t *testing.T) {
	set := New()
	assert.True(t, set.Add(1))
	assert.Equal(t, 1, set.Len())
	assert.True(t, set.Contains(1))
	assert.False(t, set.Add(1))
}

func TestHashSetAddCantDuplicate(t *testing.T) {
	set := New()
	set.Add(999)
	inserted := set.Add(999)
	assert.False(t, inserted)
	assert.Equal(t, 1, set.Len())
}

func TestHashSetString(t *testing.T) {
	set := New()
	assert.Equal(t, "Set{}", fmt.Sprint(set.String()))
	set.Add(1)
	assert.Equal(t, "Set{1}", fmt.Sprint(set.String()))
	set.Add(2)

	// builtin map doesn't guarantee order
	// we can't predict the output with multiple items
	r := set.String()
	assert.Contains(t, r, "1")
	assert.Contains(t, r, "2")
}

func TestHashSetCollect(t *testing.T) {
	set := New()
	assert.Equal(t, 0, set.Len())
	set.Add(1)
	set.Add(2)
	set.Add(3)
	collected := set.ToSlice()
	assert.Len(t, collected, 3)
	assert.Contains(t, collected, 1)
}

func TestHashSetContains(t *testing.T) {
	set := New()
	set.Add(1)
	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}

func TestHashSetRemove(t *testing.T) {
	set := New()
	set.Add(1)
	assert.False(t, set.Remove(999))
	assert.True(t, set.Remove(1))
	assert.Equal(t, 0, set.Len())
}

func TestHashSetContainsAll(t *testing.T) {
	set := New()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	assert.True(t, set.ContainsAll(1, 2, 3))
	assert.True(t, set.ContainsAll())
	assert.False(t, set.ContainsAll(1, 2, 3, 4))
	assert.True(t, set.ContainsAll(1, 2))
}

func TestHashSetLen(t *testing.T) {
	set := New()
	set.Add(1)
	assert.Equal(t, 1, set.Len())
	set.Add(1)
	assert.Equal(t, 1, set.Len())
	set.Add(2)
	assert.Equal(t, 2, set.Len())
}

func TestForEach(t *testing.T) {
	set := From(1, 2, 3)
	r := make([]int, 0, 3)

	set.ForEach(func(item interface{}) {
		r = append(r, item.(int))
	})

	assert.ElementsMatch(t, set.ToSlice(), r)
}
