package goset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSet(t *testing.T) {
	assert.Empty(t, HashSet())
	assert.Contains(t, HashSet(1), 1)
	assert.Len(t, HashSet(1, 2, 3), 3)
}

func TestHashSetAdd(t *testing.T) {
	testAdd(t, HashSet())
}

func TestHashSetAddCantDuplicate(t *testing.T) {
	testAddCantDuplicate(t, HashSet())
}
func TestHashSetString(t *testing.T) {
	testString(t, HashSet())
}

func TestHashSetCollect(t *testing.T) {
	testCollect(t, HashSet())
}

func TestHashSetContains(t *testing.T) {
	testContains(t, HashSet())
}

func TestHashSetRemove(t *testing.T) {
	testRemove(t, HashSet())
}

func TestHashSetContainsAll(t *testing.T) {
	testContainsAll(t, HashSet())
}

func TestHashSetLen(t *testing.T) {
	testLen(t, HashSet())
}
