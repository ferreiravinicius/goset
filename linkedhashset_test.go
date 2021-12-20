package goset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedHashSetRemove(t *testing.T) {
	//TODO:
}

func TestAddOneItem(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	assert.Len(t, set.hashSet, 1)
	assert.NotNil(t, set.head)
	assert.NotNil(t, set.tail)
	assert.Equal(t, 1, set.head.item)
	assert.Equal(t, 1, set.tail.item)
}

func TestLinkedHashSetAddTwoItems(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	set.Add(2)

	assert.Len(t, set.hashSet, 2)
	assert.Equal(t, 1, set.tail.item)
	assert.Equal(t, 2, set.head.item)
}

func TestLinkedHashSetAddMultipleItems(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)

	assert.Len(t, set.hashSet, 4)
	assert.Equal(t, 1, set.tail.item)
	assert.Equal(t, 4, set.head.item)

	currNode := set.tail
	for i := 1; i <= len(set.hashSet); i++ {
		assert.Equal(t, i, currNode.item)
		currNode = currNode.next
	}
}

func TestLinkedHashSetAddCantDuplicate(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	set.Add(1)
	set.Add(2)
	assert.Len(t, set.hashSet, 2)
	assert.Equal(t, 1, set.tail.item)
	assert.Equal(t, 2, set.head.item)
}

func TestLinkedHashSet_Set(t *testing.T) {
	testSet(t, func() Set {
		return LinkedHashSet()
	})
}
