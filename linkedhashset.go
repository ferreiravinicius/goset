package goset

type node struct {
	item interface{}
	next *node
	prev *node
}

type linkedHashSet struct {
	hashSet
	head *node
	tail *node
}

// Implemention of set backed by a HashSet and LinkedList.
// The LinkedList is used to gurantee order of insertion.
// Adds the item to the end (tail) of the LinkedList when using via Set.Add(...).
// This implementation is not synchronized.
func LinkedHashSet() *linkedHashSet {
	return &linkedHashSet{
		hashSet: make(hashSet),
	}
}

func (st *linkedHashSet) AddEnd(item interface{}) bool {

	if wasAdded := st.hashSet.Add(item); !wasAdded {
		return false
	}

	currNode := &node{item: item}

	if st.tail == nil {
		st.tail = currNode
		st.head = currNode
	} else {
		currNode.prev = st.tail
		st.tail.next = currNode
		st.tail = currNode
	}

	return true
}

func (st *linkedHashSet) AddStart(item interface{}) bool {

	if wasAdded := st.hashSet.Add(item); !wasAdded {
		return false
	}

	currNode := &node{item: item}

	if st.head == nil {
		st.tail = currNode
		st.head = currNode
	} else {
		currNode.next = st.head
		st.head.prev = currNode
		st.head = currNode
	}

	return true
}

func (st *linkedHashSet) Add(item interface{}) bool {
	return st.AddEnd(item)
}

func (st *linkedHashSet) Remove(item interface{}) bool {
	if wasRemoved := st.hashSet.Remove(item); !wasRemoved {
		return false
	}

	return true

}
