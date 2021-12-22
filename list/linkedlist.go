package list

type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

func (node Node) Next() *Node {
	return node.next
}

func (node Node) Prev() *Node {
	return node.prev
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (list LinkedList) Begin() *Node {
	return list.head
}

func (list LinkedList) End() *Node {
	return list.tail
}

func (list *LinkedList) AddEnd(item interface{}) *Node {
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
	return freshNode
}

func (list *LinkedList) RemoveEnd() *Node {
	return list.Remove(list.tail)
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
