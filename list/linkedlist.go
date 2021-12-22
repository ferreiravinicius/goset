package list

// Element holds the value for each item inside the list.
type Element struct {
	Value interface{}
	prev  *Element
	next  *Element
}

// Get the next element or nil if it's the end of the list.
func (el Element) Next() *Element {
	return el.next
}

// Get the previous element or nil if it's the begin of the list.
func (el Element) Prev() *Element {
	return el.prev
}

// Doubly linked list.
// Keeps order of insertion.
// Not synchronized.
type LinkedList struct {
	head *Element
	tail *Element
	size int
}

// Get the first element in this list or nil if empty - O(1).
func (list LinkedList) Begin() *Element {
	return list.head
}

// Get the last element in this list or nil if empty - O(1).
func (list LinkedList) End() *Element {
	return list.tail
}

// Returns the size of this list - O(1).
func (list LinkedList) Len() int {
	return list.size
}

// Adds the provided value at the end of this list - O(1).
// Returns the created element.
func (list *LinkedList) AddEnd(value interface{}) *Element {
	newElement := &Element{Value: value}
	if list.head == nil {
		list.head = newElement
		list.tail = newElement
	} else {
		newElement.prev = list.tail
		list.tail.next = newElement
		list.tail = newElement
	}
	list.size++
	return newElement
}

// Adds the provided value at the begin of this list - O(1).
// Returns the created element.
func (list *LinkedList) AddBegin(value interface{}) *Element {
	newElement := &Element{Value: value}
	if list.head == nil {
		list.head = newElement
		list.tail = newElement
	} else {
		newElement.next = list.head
		list.head.prev = newElement
		list.head = newElement
	}
	list.size++
	return newElement
}

// Removes the last element from this list - O(1).
// Returns the removed element or nil if list is empty.
func (list *LinkedList) RemoveEnd() *Element {
	return list.Remove(list.tail)
}

// Removes the first element from this list - O(1).
// Returns the removed element or nil if list is empty.
func (list *LinkedList) RemoveBegin() *Element {
	return list.Remove(list.head)
}

// Removes provided element from this list - O(1).
// Returns the removed element.
func (list *LinkedList) Remove(element *Element) *Element {
	if element == nil {
		return nil
	}

	if element.prev != nil {
		element.prev.next = element.next
	} else {
		list.head = element.next
	}

	if element.next != nil {
		element.next.prev = element.prev
	} else {
		list.tail = element.prev
	}

	element.prev = nil // avoid mess
	element.next = nil // avoid mess
	list.size--
	return element
}

// Returns first element containing provided value
// or nil if none is found - O(n)
func (list LinkedList) FirstMatch(value interface{}) *Element {
	for n := list.head; n != nil; n = n.next {
		if n.Value == value {
			return n
		}
	}
	return nil
}

// Returns true if this list contains provided value - O(n).
func (list LinkedList) Contains(value interface{}) bool {
	return list.FirstMatch(value) != nil
}

// Adds multiple values at the end of this list - O(1).
func (list *LinkedList) AddAll(values ...interface{}) {
	for _, value := range values {
		list.AddEnd(value)
	}
}

// Adds the provided value after other element in this list - O(1).
// Returns the created element or nil if other element is also nil.
func (list *LinkedList) AddAfter(otherElement *Element, value interface{}) *Element {
	if otherElement == nil {
		return nil
	}
	newElement := &Element{
		Value: value,
		prev:  otherElement,
		next:  otherElement.next,
	}
	otherElement.next = newElement
	if newElement.next == nil {
		list.tail = newElement
	}
	list.size++
	return newElement
}
