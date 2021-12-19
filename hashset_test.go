package goset_test

import (
	"testing"

	"github.com/ferreiravinicius/goset"
	"github.com/stretchr/testify/assert"
)

func TestHashSet(t *testing.T) {
	assert.Empty(t, goset.HashSet())
	assert.Contains(t, goset.HashSet(1), 1)
	assert.Len(t, goset.HashSet(1, 2, 3), 3)
}

func TestAdd(t *testing.T) {
	testAdd(t, goset.HashSet())
}

func TestAddCantDuplicate(t *testing.T) {
	testAddCantDuplicate(t, goset.HashSet())
}
func TestString(t *testing.T) {
	testString(t, goset.HashSet())
}

func TestCollect(t *testing.T) {
	testCollect(t, goset.HashSet())
}

func TestContains(t *testing.T) {
	testContains(t, goset.HashSet())
}

func TestRemove(t *testing.T) {
	testRemove(t, goset.HashSet())
}

func TestContainsAll(t *testing.T) {
	testContainsAll(t, goset.HashSet())
}
