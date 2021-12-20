package goset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOneItem(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	assert.Len(t, set.hashMap, 1)
	assert.NotNil(t, set.head)
	assert.NotNil(t, set.tail)
	assert.Equal(t, 1, set.head.item)
	assert.Equal(t, 1, set.tail.item)
}

func TestAddTwoItems(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	set.Add(2)

	assert.Len(t, set.hashMap, 2)
	assert.Equal(t, 1, set.tail.item)
	assert.Equal(t, 2, set.head.item)
}

func TestAddMultipleItems(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)

	assert.Len(t, set.hashMap, 4)
	assert.Equal(t, 1, set.tail.item)
	assert.Equal(t, 4, set.head.item)

	currNode := set.tail
	for i := 1; i <= len(set.hashMap); i++ {
		assert.Equal(t, i, currNode.item)
		currNode = currNode.next
	}
}

func TestAddDuplicateItems(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	set.Add(1)
	set.Add(2)
	assert.Len(t, set.hashMap, 2)
	assert.Equal(t, 1, set.tail.item)
	assert.Equal(t, 2, set.head.item)
}
