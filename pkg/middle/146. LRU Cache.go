package middle

// https://leetcode.com/problems/lru-cache/description/

type LRUCache struct {
	c int
	l int

	elems      map[int]*NodeLRU
	head, tail *NodeLRU
}

type NodeLRU struct {
	key int
	val int

	prev *NodeLRU
	next *NodeLRU
}

func NewLRU(capacity int) LRUCache {
	obj := LRUCache{
		c: capacity,
		l: 0,

		elems: make(map[int]*NodeLRU, 1<<11),
		head:  nil,
		tail:  nil,
	}

	return obj
}

func (lc *LRUCache) Get(key int) int {
	if lc.l == 0 {
		return -1
	}

	node, ok := lc.elems[key]
	if !ok {
		return -1
	}

	return swapAndReturnhHotNode(lc, node).val
}

func (lc *LRUCache) Put(key int, value int) {
	var (
		node *NodeLRU
	)

	// Empty LRU Cache +
	if lc.l == 0 {
		node = &NodeLRU{key: key, val: value, prev: nil, next: nil}
		lc.tail, lc.head = node, node

		lc.elems[key] = node

		lc.l++
		return
	}

	// Element is present +
	if node, ok := lc.elems[key]; ok {
		node.val = value
		swapAndReturnhHotNode(lc, node)
		return
	}

	// Element isn't present
	node = &NodeLRU{key: key, val: value, prev: lc.head, next: nil}
	lc.head.next = node
	lc.head = lc.head.next
	lc.elems[key] = node

	if lc.l < lc.c {
		lc.l++
		return
	}

	if lc.l == lc.c {
		delete(lc.elems, lc.tail.key)

		lc.tail = lc.tail.next
		lc.tail.prev = nil
		return
	}

}

func swapAndReturnhHotNode(lc *LRUCache, node *NodeLRU) *NodeLRU {
	switch {
	case node.prev != nil && node.next == nil:
		// Head
		return node

	case node.prev != nil && node.next != nil:
		// Intermediate elem
		node.prev.next = node.next

	case node.prev == nil && node.next != nil:
		// Tail
		lc.tail = lc.tail.next
		lc.tail.prev = nil

	}

	// Intermediate elem && Tail postprocessing
	node.next = nil
	node.prev = lc.head

	lc.head.next = node
	lc.head = node

	return node
}
