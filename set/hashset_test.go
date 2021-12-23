package set

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSet(t *testing.T) {
	assert.Empty(t, HashSet().hashMap)
	assert.Contains(t, HashSet(1).hashMap, 1)
	assert.Len(t, HashSet(1, 2, 3).hashMap, 3)
}

func TestHashSetAdd(t *testing.T) {
	set := HashSet()
	assert.True(t, set.Add(1))
	assert.Equal(t, 1, set.Len())
	assert.True(t, set.Contains(1))
	assert.False(t, set.Add(1))
}

func TestHashSetAddCantDuplicate(t *testing.T) {
	set := HashSet()
	set.Add(999)
	inserted := set.Add(999)
	assert.False(t, inserted)
	assert.Equal(t, 1, set.Len())
}

func TestHashSetString(t *testing.T) {
	set := HashSet()
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
	set := HashSet()
	assert.Equal(t, 0, set.Len())
	set.Add(1)
	set.Add(2)
	set.Add(3)
	collected := set.Collect()
	assert.Len(t, collected, 3)
	assert.Contains(t, collected, 1)
}

func TestHashSetContains(t *testing.T) {
	set := HashSet()
	set.Add(1)
	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}

func TestHashSetRemove(t *testing.T) {
	set := HashSet()
	set.Add(1)
	assert.False(t, set.Remove(999))
	assert.True(t, set.Remove(1))
	assert.Equal(t, 0, set.Len())
}

func TestHashSetContainsAll(t *testing.T) {
	set := HashSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	assert.True(t, set.ContainsAll(1, 2, 3))
	assert.True(t, set.ContainsAll())
	assert.False(t, set.ContainsAll(1, 2, 3, 4))
	assert.True(t, set.ContainsAll(1, 2))
}

func TestHashSetLen(t *testing.T) {
	set := HashSet()
	set.Add(1)
	assert.Equal(t, 1, set.Len())
	set.Add(1)
	assert.Equal(t, 1, set.Len())
	set.Add(2)
	assert.Equal(t, 2, set.Len())
}
