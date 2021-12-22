package list

type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

func (it Node) Next() *Node {
	return it.next
}

func (it Node) Prev() *Node {
	return it.prev
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (it LinkedList) Begin() *Node {
	return it.head
}

func (it LinkedList) End() *Node {
	return it.tail
}

func (it *LinkedList) AddEnd(item interface{}) *Node {
	freshNode := &Node{Value: item}
	if it.head == nil {
		it.head = freshNode
		it.tail = freshNode
	} else {
		freshNode.prev = it.tail
		it.tail.next = freshNode
		it.tail = freshNode
	}
	it.size++
	return freshNode
}

func (it *LinkedList) RemoveEnd() *Node {
	return it.Remove(it.tail)
}

func (list *LinkedList) Remove(node *Node) *Node {

	if node == nil {
		return nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}

	node.prev = nil // avoid mess
	node.next = nil // avoid mess
	list.size--
	return node
}
