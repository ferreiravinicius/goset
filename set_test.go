package goset_test

import (
	"fmt"
	"testing"

	"github.com/ferreiravinicius/goset"
	"github.com/stretchr/testify/assert"
)

func testAdd(t *testing.T, set goset.Set) {
	assert.True(t, set.Add(1))
	assert.Len(t, set, 1)
	assert.Contains(t, set, 1)
	assert.False(t, set.Add(1))
}

func testAddCantDuplicate(t *testing.T, set goset.Set) {
	set.Add(999)
	inserted := set.Add(999)
	assert.False(t, inserted)
	assert.Len(t, set, 1)
}

func testString(t *testing.T, set goset.Set) {
	stringer := set.(fmt.Stringer)
	assert.Equal(t, "Set{}", fmt.Sprint(stringer.String()))
	set.Add(1)
	assert.Equal(t, "Set{1}", fmt.Sprint(stringer.String()))
	set.Add(2)
	assert.Equal(t, "Set{1 2}", fmt.Sprint(stringer.String()))
}

func testCollect(t *testing.T, set goset.Set) {
	assert.Len(t, set, 0)
	set.Add(1)
	set.Add(2)
	set.Add(3)
	collected := set.Collect()
	assert.Len(t, collected, 3)
	assert.Contains(t, collected, 1)
}

func testContains(t *testing.T, set goset.Set) {
	set.Add(1)
	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}



