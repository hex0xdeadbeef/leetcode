package doublelinkedlist

const (
	empty    = 0
	onlyHead = 1
)

type Node struct {
	Val int

	prev, next *Node
}

type List struct {
	tail, head *Node

	count uint64
}

func New() *List {
	return &List{}
}

func (l *List) Head() *Node {
	if l.head != nil {
		return l.head
	}

	return nil
}

func (l *List) Tail() *Node {
	if l.tail != nil {
		return l.tail
	}

	return nil
}

func (l *List) Count() uint64 {
	return l.count
}

func (l *List) InsertTail(val int) {
	n := &Node{Val: val}

	switch l.count {
	case empty:
		l.head, l.tail = n, n

	case onlyHead:
		l.tail, l.head.prev, n.next = n, n, l.head

	default:
		n.next, l.tail.prev, l.tail = l.tail, n, n
	}

	l.count++
}

func (l *List) InsertHead(val int) {
	n := &Node{Val: val}

	switch l.count {
	case empty:
		l.head = n

	case onlyHead:
		l.tail, l.head.next, n.prev, l.head = l.head, n, l.head, n

	default:
		l.head.next, n.prev, l.head = n, l.head, n
	}

	l.count++
}

func (l *List) PopTail() *Node {
	switch l.count {
	case empty:
		return nil

	case onlyHead:
		n := l.head
		l.head, l.tail = nil, nil
		l.count--
		return n

	default:
		n := l.tail
		l.tail.next.prev, l.tail.next, l.tail = nil, nil, l.tail.next
		l.count--
		return n
	}

}

func (l *List) PopHead() *Node {
	switch l.count {
	case empty:
		return nil

	case onlyHead:
		n := l.head
		l.head, l.tail = nil, nil
		l.count--
		return n

	default:
		n := l.head
		l.head.prev.next, l.head.prev, l.head = nil, nil, l.head.prev
		l.count--
		return n
	}
}

func (l *List) IsEmpty() bool {
	return l.tail == nil && l.head == nil
}
