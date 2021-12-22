package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListAddEndReturnElement(t *testing.T) {
	l := new(LinkedList)
	el := l.AddEnd(1)
	assert.NotNil(t, el)
	assert.Equal(t, 1, el.Value)
}

func TestLinkedListAddEndOne(t *testing.T) {
	l := new(LinkedList)
	l.AddEnd(1)
	assert.NotNil(t, l.head)
	assert.NotNil(t, l.tail)
	assert.Equal(t, 1, l.size)
}

func TestLinkedListAddEndMoreThanOne(t *testing.T) {
	list := new(LinkedList)
	assert.NotNil(t, list.AddEnd(1))
	assert.NotNil(t, list.AddEnd(2))
	assert.NotNil(t, list.AddEnd(3))
	assert.Equal(t, 1, list.head.Value)
	assert.Equal(t, 3, list.tail.Value)
	assert.Equal(t, 3, list.size)
	assert.Equal(t, 2, list.head.next.Value)
}

func TestLinkedListEnd(t *testing.T) {
	l := new(LinkedList)
	assert.Nil(t, l.End())
	l.AddEnd(1)
	assert.Equal(t, 1, l.End().Value)
	l.AddEnd(2)
	assert.Equal(t, 2, l.End().Value)
}

func TestLinkedListIterateOverItems(t *testing.T) {
	list := new(LinkedList)
	list.AddEnd(1)
	list.AddEnd(2)
	list.AddEnd(3)

	r := make([]int, 0, 3)
	for el := list.Begin(); el != nil; el = el.Next() {
		r = append(r, el.Value.(int))
	}
	assert.ElementsMatch(t, r, []int{1, 2, 3})
}
func TestLinkedListIterateReverse(t *testing.T) {
	list := new(LinkedList)
	list.AddEnd(1)
	list.AddEnd(2)
	list.AddEnd(3)

	r := make([]int, 0, 3)
	for el := list.End(); el != nil; el = el.Prev() {
		r = append(r, el.Value.(int))
	}
	assert.ElementsMatch(t, r, []int{1, 2, 3})
}

func TestLinkedListLen(t *testing.T) {
	l := new(LinkedList)
	assert.Equal(t, 0, l.Len())
	l.AddEnd(1)
	assert.Equal(t, 1, l.Len())
	l.AddBegin(2)
	assert.Equal(t, 2, l.Len())
}

func TestLinkedListAddBegin(t *testing.T) {
	l := new(LinkedList)
	assert.NotNil(t, l.AddBegin(1))
	assert.Equal(t, l.head, l.AddBegin(2))
	assert.Equal(t, l.head, l.AddBegin(3))
	assert.Equal(t, 3, l.head.Value)
	assert.Equal(t, 1, l.tail.Value)
	assert.Equal(t, 2, l.tail.prev.Value)
	assert.Equal(t, 2, l.head.next.Value)
	assert.Equal(t, 3, l.Len())
}

func TestLinkedListRemoveEnd(t *testing.T) {
	l := new(LinkedList)
	assert.Nil(t, l.RemoveEnd())
	assert.Equal(t, 0, l.size)
	l.AddEnd(1)
	removed := l.RemoveEnd()
	assert.NotNil(t, removed)
	assert.Equal(t, 1, removed.Value)
	assert.Nil(t, removed.next)
	assert.Nil(t, removed.prev)

	l.AddEnd(1)
	l.AddEnd(2)
	l.RemoveEnd()
	assert.Equal(t, 1, l.size)

	l.AddBegin(9)
	assert.Equal(t, 1, l.RemoveEnd().Value)
}

func TestLinkedListRemove(t *testing.T) {
	l := new(LinkedList)
	assert.Nil(t, l.Remove(nil))

	el := l.AddBegin(1)
	assert.NotNil(t, l.Remove(el))
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)

	l.AddBegin(3)
	two := l.AddBegin(2)
	l.AddBegin(1)
	assert.NotNil(t, l.Remove(two))
	assert.Equal(t, 1, l.head.Value)
	assert.Equal(t, 3, l.tail.Value)
	assert.Equal(t, 1, l.tail.prev.Value)
	assert.Equal(t, 3, l.head.next.Value)
}

func TestLinkedListFirstMatch(t *testing.T) {
	l := new(LinkedList)
	assert.Nil(t, l.FirstMatch(1))
	one := l.AddBegin(1)
	assert.Equal(t, one, l.FirstMatch(1))
	two := l.AddEnd(2)
	l.AddEnd(3)
	assert.Equal(t, two, l.FirstMatch(2))
}

func TestLinkedListContains(t *testing.T) {
	l := new(LinkedList)
	assert.False(t, l.Contains(1))
	l.AddEnd(2)
	assert.False(t, l.Contains(1))
	l.AddEnd(1)
	assert.True(t, l.Contains(1))
}

func TestLinkedListRemoveBegin(t *testing.T) {
	l := new(LinkedList)
	assert.Nil(t, l.RemoveBegin())

	one := l.AddEnd(1)
	l.AddEnd(2)
	l.AddEnd(3)
	assert.Equal(t, 1, l.head.Value)
	assert.Equal(t, one, l.RemoveBegin())
	assert.Equal(t, 2, l.head.Value)
}

func TestLinkedListAddAll(t *testing.T) {
	l := new(LinkedList)
	l.AddAll(1, 2, 3)
	assert.Equal(t, 3, l.size)
	assert.Equal(t, 1, l.head.Value)
	assert.Equal(t, 3, l.tail.Value)
}
