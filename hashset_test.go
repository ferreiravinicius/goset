package gset_test

import (
	"testing"

	"github.com/ferreiravinicius/gset"
	"github.com/stretchr/testify/assert"
)

func TestHashSet(t *testing.T) {
	assert.Empty(t, gset.HashSet())
	assert.Contains(t, gset.HashSet(1), 1)
	assert.Len(t, gset.HashSet(1, 2, 3), 3)
}

func TestAdd(t *testing.T) {
	testAdd(t, gset.HashSet())
}

func TestString(t *testing.T) {
	testString(t, gset.HashSet())
}

func TestCollect(t *testing.T) {
	testCollect(t, gset.HashSet())
}
