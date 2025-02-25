package middle

// https://leetcode.com/problems/lru-cache/description/

type (
	LRUCache struct {
		nodes      map[int]*NodeLRU
		tail, head *NodeLRU
		capacity   int
	}

	NodeLRU struct {
		k, v       int
		prev, next *NodeLRU
	}
)

func NewLRU(capacity int) LRUCache {
	return LRUCache{
		nodes:    make(map[int]*NodeLRU, capacity),
		capacity: capacity,
	}
}

func (lc *LRUCache) Put(key int, value int) {
	node := &NodeLRU{
		k: key,
		v: value,
	}

	if len(lc.nodes) == 0 {
		lc.head, lc.nodes[key] = node, node
		return
	}

	if len(lc.nodes) != lc.capacity && len(lc.nodes) == 1 {
		lc.tail, lc.head, lc.head.next, node.prev, lc.nodes[key] = lc.head, node, node, lc.head, node
		return
	}

	_, ok := lc.nodes[key]
	if !ok {
		lc.putNotPresent(key, value)
		return
	}

	lc.putPresent(key, value)
}

func (lc *LRUCache) Get(key int) int {
	node, ok := lc.nodes[key]
	if !ok {
		return -1
	}

	lc.turnToHead(node)
	return node.v
}

func (lc *LRUCache) putPresent(key, value int) {
	lc.nodes[key].v = value
	lc.turnToHead(lc.nodes[key])
}

func (lc *LRUCache) putNotPresent(key, value int) {
	node := &NodeLRU{
		k: key,
		v: value,
	}

	if len(lc.nodes) == lc.capacity {
		if lc.capacity == 1 {
			delete(lc.nodes, lc.head.k)
			lc.head, lc.nodes[key] = node, node
			return
		}

		delete(lc.nodes, lc.tail.k)
		lc.tail, lc.tail.next.prev = lc.tail.next, nil
		lc.head, lc.head.next, node.prev, lc.nodes[key] = node, node, lc.head, node
		return
	}

	lc.nodes[key] = node
	if len(lc.nodes) == 1 {
		lc.tail, lc.head, lc.head.next, node.prev = lc.head, node, node, lc.head
		return
	}
	lc.head, lc.head.next, node.prev = node, node, lc.head
}

func (lc *LRUCache) turnToHead(node *NodeLRU) {
	if len(lc.nodes) <= 1 {
		return
	}

	switch node {
	case lc.head:
		return

	case lc.tail:
		lc.tail, lc.tail.next, lc.tail.next.prev = lc.tail.next, nil, nil

	default:
		node.prev.next, node.next.prev = node.next, node.prev
	}

	node.prev, lc.head.next, lc.head = lc.head, node, node
}
