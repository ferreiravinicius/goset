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

func TestHashSet_Set(t *testing.T) {
	testSet(t, func() Set {
		return HashSet()
	})
}
