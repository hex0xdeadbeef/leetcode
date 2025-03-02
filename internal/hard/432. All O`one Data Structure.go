package hard

/*
a := hard.Constructor()

	// Increment operations
	a.Inc("a")
	a.Inc("b")
	a.Inc("c")
	a.Inc("d")
	a.Inc("e")
	a.Inc("f")
	a.Inc("g")
	a.Inc("h")
	a.Inc("i")
	a.Inc("j")
	a.Inc("k")
	a.Inc("l")

	// Decrement operations
	a.Dec("a")
	a.Dec("b")
	a.Dec("c")
	a.Dec("d")
	a.Dec("e")
	a.Dec("f")

	// Increment operations
	a.Inc("g")
	a.Inc("h")
	a.Inc("i")
	a.Inc("j")

	// Get max and min keys
	fmt.Println(a.GetMaxKey()) // Expected: "i"
	fmt.Println(a.GetMinKey()) // Expected: "k"

	// Additional increment operations
	a.Inc("k")
	a.Inc("l")

	// Get max and min keys again
	fmt.Println(a.GetMaxKey()) // Expected: "g"
	fmt.Println(a.GetMinKey()) // Expected: "g"

	// Final operations
	a.Inc("a")
	a.Dec("j")

	// Final max/min key queries
	fmt.Println(a.GetMaxKey()) // Expected: "g"
	fmt.Println(a.GetMinKey()) // Expected: "a"
*/

type (
	AllOne struct {
		body       map[string]*Node
		tail, head *Node
	}

	Node struct {
		val string
		cnt int

		prev, next *Node
	}
)

func Constructor() AllOne {
	return AllOne{
		body: make(map[string]*Node, 1<<10),
	}
}

func (a *AllOne) Inc(key string) {
	if len(a.body) == 0 {
		n := &Node{
			val: key,
			cnt: 1,
		}

		a.head, a.body[key] = n, n
		return
	}

	_, ok := a.body[key]
	if ok {
		a.incPresent(key)
		return
	}
	a.insertTail(key)
}

func (a *AllOne) incPresent(key string) {
	n := a.body[key]
	n.cnt++

	if a.head == n {
		return
	}

	if n.cnt <= n.next.cnt {
		return
	}

	if n.next == a.head && n.cnt > a.head.cnt {
		if n == a.tail {
			a.tail = a.head
		}

		start := n.prev
		if start != nil {
			start.next = a.head
		}

		n.next, n.prev, a.head.next, a.head.prev, a.head = nil, a.head, n, n.prev, n
		return
	}

	if n == a.tail {
		a.tail = n.next
	}

	start, finish := n.prev, n.next.next
	if start != nil {
		start.next = n.next
	}
	if finish != nil {
		finish.prev = n
	}
	n.next, n.next.next, n.prev, n.next.prev = finish, n, n.next, start
}

func (a *AllOne) insertTail(key string) {
	n := &Node{
		val: key,
		cnt: 1,
	}

	if len(a.body) == 1 {
		n.next, a.head.prev, a.tail, a.body[key] = a.head, n, n, n
		return
	}

	n.next, a.tail.prev, a.tail, a.body[key] = a.tail, n, n, n
}

func (a *AllOne) Dec(key string) {
	if len(a.body) == 0 {
		return
	}

	n, ok := a.body[key]
	if !ok {
		return
	}

	n.cnt--
	if n.prev != nil && n.cnt < n.prev.cnt {
		a.moveNodeLeft(key)
	}

	if n.cnt == 0 {
		a.removeNode(key)
	}
}

func (a *AllOne) moveNodeLeft(key string) {
	n := a.body[key]

	if n.prev == nil {
		return
	}

	if n.prev == a.tail {
		a.tail = n
	}

	if n == a.head {
		a.head = n.prev

		start := n.prev.prev
		if start != nil {
			start.next = n
		}

		n.prev.next, n.next, n.prev, n.prev.prev = nil, n.prev, start, n
		return
	}

	start, finish := n.prev.prev, n.next
	if start != nil {
		start.next = n.prev
	}
	if finish != nil {
		finish.prev = n.prev
	}
	n.next, n.prev, n.prev.next, n.prev.prev = n.prev, start, finish, n

}

func (a *AllOne) removeNode(key string) {
	n := a.body[key]
	delete(a.body, n.val)

	if n.prev != nil && n.next != nil {
		n.prev.next, n.next.prev = n.next.prev, n.prev.next
		return
	}

	if n == a.head {
		n.prev.next, a.head = nil, n.prev
		return
	}

	a.tail, n.next.prev = n.next, nil
}

func (a *AllOne) GetMaxKey() string {
	if len(a.body) == 0 {
		return ""
	}

	return a.head.val

}

func (a *AllOne) GetMinKey() string {
	if len(a.body) == 0 {
		return ""
	}

	if a.tail == nil {
		return a.head.val
	}

	return a.tail.val
}
