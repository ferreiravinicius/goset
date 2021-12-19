package gset_test

import (
	"fmt"
	"testing"

	"github.com/ferreiravinicius/gset"
	"github.com/stretchr/testify/assert"
)

func testAdd(t *testing.T, s gset.Set) {
	assert.True(t, s.Add("t"))
	assert.Len(t, s, 1)
	assert.Contains(t, s, "t")
	assert.False(t, s.Add("t"))
}

func testString(t *testing.T, set gset.Set) {
	if stringer, ok := set.(fmt.Stringer); ok {
		assert.Equal(t, "Set{}", fmt.Sprint(stringer.String()))
		set.Add(1)
		assert.Equal(t, "Set{1}", fmt.Sprint(stringer.String()))
		set.Add(2)
		assert.Equal(t, "Set{1 2}", fmt.Sprint(stringer.String()))
	}
}

func testCollect(t *testing.T, set gset.Set) {
	assert.Len(t, set, 0)
	set.Add(1)
	set.Add(2)
	set.Add(3)
	collected := set.Collect()
	assert.Len(t, collected, 3)
	assert.Contains(t, collected, 1)
}
