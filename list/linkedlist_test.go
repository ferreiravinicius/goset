package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// type List interface {
// 	Add(item interface{}) bool
// 	AddAll(items ...interface{}) bool
// 	Len() int
// 	Get(index int) interface{}
// 	RemoveAt(index int) bool
// 	Remove(item interface{}) bool
// }

type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

func (it Node) HasNext() bool {
	return it.next != nil
}

func (it Node) Next() *Node {
	return it.next
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (list *LinkedList) AddEnd(item interface{}) bool {
	freshNode := &Node{Value: item}
	if list.head == nil {
		list.head = freshNode
		list.tail = freshNode
	} else {
		freshNode.prev = list.tail
		list.tail.next = freshNode
		list.tail = freshNode
	}
	list.size++
	return true
}

func TestAddOne(t *testing.T) {
	list := new(LinkedList)
	assert.True(t, list.AddEnd(1))
	assert.NotNil(t, list.head)
	assert.NotNil(t, list.tail)
	assert.Equal(t, 1, list.size)
}

func TestAddEndMoreThanOne(t *testing.T) {
	list := new(LinkedList)
	assert.True(t, list.AddEnd(1))
	assert.True(t, list.AddEnd(2))
	assert.True(t, list.AddEnd(3))
	assert.Equal(t, 1, list.head.Value)
	assert.Equal(t, 3, list.tail.Value)
	assert.Equal(t, 3, list.size)
}

func (list LinkedList) Begin() *Node {
	return list.head
}

func (list LinkedList) End() *Node {
	return list.tail
}

func TestEnd(t *testing.T) {
	l := new(LinkedList)
	assert.Nil(t, l.End())
	l.AddEnd(1)
	assert.Equal(t, 1, l.End().Value)
	l.AddEnd(2)
	assert.Equal(t, 2, l.End().Value)
}

func TestCanIterateOverItems(t *testing.T) {
	list := new(LinkedList)
	list.AddEnd(1)
	list.AddEnd(2)
	list.AddEnd(3)

	r := make([]int, 0, 3)
	for n := list.Begin(); n != nil; n = n.Next() {
		r = append(r, n.Value.(int))
	}
	assert.ElementsMatch(t, r, []int{1, 2, 3})
}

func (list LinkedList) Len() int {
	return list.size
}

func TestLinkedListLen(t *testing.T) {
	l := new(LinkedList)
	assert.Equal(t, 0, l.Len())
	l.AddEnd(1)
	assert.Equal(t, 1, l.Len())
}

func (list *LinkedList) AddBegin(item interface{}) *Node {
	freshNode := &Node{Value: item}
	if list.head == nil {
		list.head = freshNode
		list.tail = freshNode
	} else {
		freshNode.next = list.head
		list.head.prev = freshNode
		list.head = freshNode
	}
	return freshNode
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
}
