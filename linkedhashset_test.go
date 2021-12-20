package goset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedHashSetRemove(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	set.Add(2)
	set.Remove(2)
	assert.Equal(t, 1, set.Len())
	// TODO: implement linked list stuff for this
	// assert.Equal(t, 1, set.head.item)
	// assert.Equal(t, 1, set.tail.item)

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
	assert.Equal(t, 2, set.tail.item)
	assert.Equal(t, 1, set.head.item)
}
func TestLinkedHashSetAddStart(t *testing.T) {
	set := LinkedHashSet()
	assert.True(t, set.AddStart(1))
	assert.Equal(t, 1, set.tail.item)
	assert.Equal(t, 1, set.head.item)

	assert.True(t, set.AddStart(2))
	assert.Equal(t, 1, set.tail.item)
	assert.Equal(t, 2, set.head.item)

	assert.True(t, set.AddStart(3))
	assert.Equal(t, 1, set.tail.item)
	assert.Equal(t, 3, set.head.item)

	curr := set.tail
	for num := 1; curr != nil; num++ {
		assert.Equal(t, num, curr.item)
		curr = curr.prev
	}

}
func TestLinkedHashSetAddStartCantDuplicate(t *testing.T) {
	set := LinkedHashSet()
	assert.True(t, set.AddStart(1))
	assert.False(t, set.AddStart(1))
	assert.Equal(t, 1, set.Len())
}

func TestLinkedHashSetAddEndCantDuplicate(t *testing.T) {
	set := LinkedHashSet()
	assert.True(t, set.AddEnd(1))
	assert.False(t, set.AddEnd(1))
	assert.Equal(t, 1, set.Len())
}

func TestLinkedHashSetAddMultipleItems(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)

	assert.Len(t, set.hashSet, 4)
	assert.Equal(t, 4, set.tail.item)
	assert.Equal(t, 1, set.head.item)

	curr := set.head
	for num := 1; curr != nil; num++ {
		assert.Equal(t, num, curr.item)
		curr = curr.next
	}
}

func TestLinkedHashSetAddCantDuplicate(t *testing.T) {
	set := LinkedHashSet()
	set.Add(1)
	set.Add(1)
	set.Add(2)
	assert.Len(t, set.hashSet, 2)
	assert.Equal(t, 2, set.tail.item)
	assert.Equal(t, 1, set.head.item)
}

func TestLinkedHashSetAddEnd(t *testing.T) {
	set := LinkedHashSet()
	assert.True(t, set.AddEnd(1))
	assert.Equal(t, 1, set.tail.item)
	assert.Equal(t, 1, set.head.item)

	assert.True(t, set.AddEnd(2))
	assert.Equal(t, 2, set.tail.item)
	assert.Equal(t, 1, set.head.item)

	assert.True(t, set.AddEnd(3))
	assert.Equal(t, 3, set.tail.item)
	assert.Equal(t, 1, set.head.item)

	curr := set.head
	for num := 1; curr != nil; num++ {
		assert.Equal(t, num, curr.item)
		curr = curr.next
	}
}

func TestLinkedHashSet_Set(t *testing.T) {
	testSet(t, func() Set {
		return LinkedHashSet()
	})
}
