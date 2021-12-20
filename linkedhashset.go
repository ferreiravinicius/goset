package goset

type node struct {
	item interface{}
	next *node
	prev *node
}

type linkedhashset struct {
	hashMap map[interface{}]struct{}
	head    *node
	tail    *node
}

func LinkedHashSet() linkedhashset {
	return linkedhashset{
		hashMap: make(map[interface{}]struct{}),
	}
}

func (st *linkedhashset) Add(item interface{}) bool {
	if _, exists := st.hashMap[item]; exists {
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

	st.hashMap[item] = struct{}{}
	return true
}
