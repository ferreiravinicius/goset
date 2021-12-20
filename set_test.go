package goset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testAdd(t *testing.T, set Set) {
	assert.True(t, set.Add(1))
	assert.Equal(t, 1, set.Len())
	assert.True(t, set.Contains(1))
	assert.False(t, set.Add(1))
}

func testAddCantDuplicate(t *testing.T, set Set) {
	set.Add(999)
	inserted := set.Add(999)
	assert.False(t, inserted)
	assert.Equal(t, 1, set.Len())
}

func testString(t *testing.T, set Set) {
	stringer := set.(fmt.Stringer)
	assert.Equal(t, "Set{}", fmt.Sprint(stringer.String()))
	set.Add(1)
	assert.Equal(t, "Set{1}", fmt.Sprint(stringer.String()))
	set.Add(2)

	// builtin map doesn't guarantee order
	// we can't predict the output with multiple items
	r := stringer.String()
	assert.Contains(t, r, "1")
	assert.Contains(t, r, "2")
}

func testCollect(t *testing.T, set Set) {
	assert.Equal(t, 0, set.Len())
	set.Add(1)
	set.Add(2)
	set.Add(3)
	collected := set.Collect()
	assert.Len(t, collected, 3)
	assert.Contains(t, collected, 1)
}

func testContains(t *testing.T, set Set) {
	set.Add(1)
	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}

func testRemove(t *testing.T, set Set) {
	set.Add(1)
	assert.False(t, set.Remove(999))
	assert.True(t, set.Remove(1))
	assert.Equal(t, 0, set.Len())
}

func testContainsAll(t *testing.T, set Set) {
	set.Add(1)
	set.Add(2)
	set.Add(3)
	assert.True(t, set.ContainsAll(1, 2, 3))
	assert.True(t, set.ContainsAll())
	assert.False(t, set.ContainsAll(1, 2, 3, 4))
	assert.True(t, set.ContainsAll(1, 2))
}

func testLen(t *testing.T, set Set) {
	set.Add(1)
	assert.Equal(t, 1, set.Len())
	set.Add(1)
	assert.Equal(t, 1, set.Len())
	set.Add(2)
	assert.Equal(t, 2, set.Len())
}
