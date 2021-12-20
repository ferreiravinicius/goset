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
	testSetAdd(t, HashSet())
}

func TestHashSetAddCantDuplicate(t *testing.T) {
	testSetAddCantDuplicate(t, HashSet())
}
func TestHashSetString(t *testing.T) {
	testSetString(t, HashSet())
}

func TestHashSetCollect(t *testing.T) {
	testSetCollect(t, HashSet())
}

func TestHashSetContains(t *testing.T) {
	testSetContains(t, HashSet())
}

func TestHashSetRemove(t *testing.T) {
	testSetRemove(t, HashSet())
}

func TestHashSetContainsAll(t *testing.T) {
	testSetContainsAll(t, HashSet())
}

func TestHashSetLen(t *testing.T) {
	testSetLen(t, HashSet())
}
