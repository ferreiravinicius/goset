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

func LinkedHashSet() *linkedHashSet {
	return &linkedHashSet{
		hashSet: make(hashSet),
	}
}

func (st *linkedHashSet) Add(item interface{}) bool {

	if wasAdded := st.hashSet.Add(item); !wasAdded {
		return false
	}

	currNode := &node{item: item}

	if st.head == nil {
		st.tail = currNode
		st.head = currNode
	} else {
		currNode.prev = st.head
		st.head.next = currNode
		st.head = currNode
	}

	return true
}
