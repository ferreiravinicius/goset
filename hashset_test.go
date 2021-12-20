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

func TestAdd(t *testing.T) {
	testAdd(t, HashSet())
}

func TestAddCantDuplicate(t *testing.T) {
	testAddCantDuplicate(t, HashSet())
}
func TestString(t *testing.T) {
	testString(t, HashSet())
}

func TestCollect(t *testing.T) {
	testCollect(t, HashSet())
}

func TestContains(t *testing.T) {
	testContains(t, HashSet())
}

func TestRemove(t *testing.T) {
	testRemove(t, HashSet())
}

func TestContainsAll(t *testing.T) {
	testContainsAll(t, HashSet())
}
